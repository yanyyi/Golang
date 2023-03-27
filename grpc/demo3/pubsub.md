### 发布和订阅模式

在发布和订阅模式中，由调用者主动发起的发布行为类似一个普通函数调用，而被动的订阅者则类似gRPC客户端单向流中的接收者。demo3尝试基于gRPC的流特性构造一个发布和订阅系统。

发布和订阅是一个常见的设计模式，开源社区中已经存在很多该模式的实现。其中Docker项目中提供了一个pubsub的极简实现,代码如下

```go
package main

import (
	"fmt"
	"github.com/docker/docker/pkg/pubsub"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})

	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("docker: https://www.docker.com")

	go func() {
		fmt.Println("golang topic:", <-golang)
	}()

	go func() {
		fmt.Println("docker topic:", <-docker)
	}()

	<-make(chan bool)
}

```

其中pubsub.NewPublisher构造一个发布对象,p.SubscribeTopic()可以通过函数筛选感兴趣的主体进行订阅
