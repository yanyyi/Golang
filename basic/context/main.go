package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 创建 HTTP 服务器
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 使用 WithTimeout 创建一个具有超时机制的 context
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// 执行一些耗时的工作，如果在 ctx 上下文超时或取消，则退出
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello, world!")
	case <-ctx.Done():
		fmt.Fprint(w, "Cancelled or timed out.")
	}

}

/*
在上述示例中，我们首先使用 http.HandleFunc() 函数将处理器函数 handler 注册到默认的 HTTP 服务器上。然后在 handler 函数中，我们使用 context.WithTimeout 创建了一个具有超时机制的 context，并将其传递给一些耗时的工作代码。

在 select 代码块中，我们使用 time.After() 语句模拟了一个持续 10 秒钟的耗时操作。如果在此过程中，context 接收到了取消信号，那么 case <- ctx.Done() 分支将被执行，并立即返回一个响应。否则，工作将等待 10 秒钟，并在完成后返回一个响应。

需要注意的是，在函数执行结束后，我们必须使用 defer cancel() 关闭 context，即使在 select 代码块执行之前就出现了错误或异常情况。这样可以避免意外的资源泄漏和意外的超时行为。

这只是一个简单的示例，在实际应用中 context 可以有更多的用途和变体，通过使用 context 可以实现更好的代码控制和错误处理。
*/
