# 我的笔记

本项目参考极客兔兔7天用Go从零实现Web框架Gee教程，自定义实现了Dragon框架  
极客兔兔网址:https://geektutu.com/post/gee.html

### 设计一个框架  
Web框架:  
Python: django、flask  
Golang: Gin、Beego、Iris

net/http提供了基础的Web功能，即监听端口，映射静态路由，解析HTTP报文。  
然而一些Web开发中简单的需求并不支持，需要手动实现  
· 动态路由:例如`hello/:name`,`hello/*`这类的规则  
· 鉴权:没有分组/统一鉴权的能力，需要再每个路由映射的handler中实现  
· 模板:没有统一简化的HTML机制  
· ......  

微框架提供的特性:  
· 路由(Routing):将请求映射到函数，支持动态路由。例如:`/hello/:name`  
· 模板(Templates):使用内置模板引擎提供模板渲染机制。  
· 工具集(Utilites):提供对cookies,headers等处理机制  
· 插件(Plugin):插件机制，可以选择安装到全局，也可以只针对某几个路由生效  
· ......


## 第一天
`
简单介绍net/http库以及http.Handler接口
` 

`
搭建dragon框架的雏形
`
### 标准库启动Web服务
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", testHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

```
main函数的最后一行，是用来启动Web服务的。  
第一个参数是地址,而第二个参数则代表处理所有的HTTP请求的实例,nil代表使用标准库中的实例处理。

### 实现http.Handler接口  
```go
package http

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

func ListenAndServe(address string, h Handler) error
```
第二个参数的类型通过查看`net/http`源码可以发现，Handler是一个接口，需要实现方法ServerHTTP.也就是说，只要传入任何实现了ServeHTTP接口的实例，所有的HTTP请求，就都交给了该实例处理。
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
```
我们定义了一个空的结构体Engine，实现了方法ServeHTTP。这个方法有2个参数，第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应。

在 main 函数中，我们给 ListenAndServe 方法的第二个参数传入了刚才创建的engine实例。至此，我们走出了实现Web框架的第一步，即，将所有的HTTP请求转向了我们自己的处理逻辑。还记得吗，在实现Engine之前，我们调用 http.HandleFunc 实现了路由和Handler的映射，也就是只能针对具体的路由写处理逻辑。比如/hello。但是在实现Engine之后，我们拦截了所有的HTTP请求，拥有了统一的控制入口。在这里我们可以自由定义路由映射的规则，也可以统一添加一些处理逻辑，例如日志、异常处理等。


### Dragon框架的雏形
```go
dragon/
   |--dragon.go 
   |--go.mod
main.go
go.mod
```
dragon/dragon.go
```go
package dragon

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND：%s\n", r.URL)
	}
}

```

main.go
```go
package main

import (
	"Golang/gee-web/day0-myWeb/dragon"
	"fmt"
	"net/http"
)

func main() {
	r := dragon.New()
	r.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "output:%q\n", r.URL.Path)
	})

	r.GET("/hello", helloHandler)
	r.Run(":8080")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

```
## 第二天  
`
将路由(Router)独立出来，方便之后增强
`

`
设计上下文(Context),封装Request和Response,提供对JSON、HTML等返回类型的支持
`



### 设计Context  
1. 对于Web服务来说，无非是根据请求*http.Request，构造响应http.ResponseWriter。但这两个对象提供的接口粒度太细，
比如我们要构造一个完整的响应，需要考虑消息头(Header)和消息体(Body),而Header包含了状态码(StatusCode),消息类型(ContentType)等
几乎每次请求都需要设置的信息。因此，如果不进行有效的封装，那么框架的用户将需要写大量重复、繁杂的代码。而且容易出错。
针对常用场景，能够高效地构造出HTTP响应是一个好的框架必须考虑的点。  

用返回JSON数据作比较，感受一下封装的差距  
封装前
```go
obj = map[string]interface{}{
    "name": "geektutu",
    "password": "1234",
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
encoder := json.NewEncoder(w)
if err := encoder.Encode(obj); err != nil {
    http.Error(w, err.Error(), 500)
}
```  
封装后
```go
c.JSON(http.StatusOK, gee.H{
    "username": c.PostForm("username"),
    "password": c.PostForm("password"),
})
```
2. 针对使用场景，封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，只是设计 Context 的原因之一。对于框架来说，还需要支撑额外的功能。例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？再比如，框架需要支持中间件，那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。路由的处理函数，以及将要实现的中间件，参数都统一使用 Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。
### 路由(Router)
我们将和路由相关的方法和结构提取了出来，放到了一个新的文件中router.go，方便我们下一次对 router 的功能进行增强，例如提供动态路由的支持。 router 的 handle 方法作了一个细微的调整，即 handler 的参数，变成了 Context。
```go
package dragon

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s \n", c.Path)
	}
}

```

### 框架入口
```go
package dragon

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}

```

## 第三天

`
使用Trie树实现动态路由(dynamic route)解析
`

`
支持两种模式:name和*filepath
`
### Trie树简介
#### 使用map存储路由表具有很大的弊端：键值对的存储的方式，只能用来索引静态路由
#### 要想实现动态路由，有很多种实现方式。例如开源的路由实现gorouter支持在路由规则中嵌入正则表达式，例如/p/[0-9A-Za-z]+,即路径中的参数仅匹配数字和字母；另一个开源实现httprouter就不支持正则表达式。注明的Web开源框架gin在早期的版本，并没有实现自己的路由，而是直接使用了Httprouter,后来放弃了httprouter，自己实现了一个版本
#### 实现动态路由最常见的数据结构，被称为前缀树(Trie)树,每一个节点的所有子节点都拥有相同的前缀，这种结构非常适用于路由匹配
#### HTTP请求的路径恰好是由'/'分隔的多段构成的，因此每一段可以作为前缀树的一个节点。我们通过树结构查询，如果中间某一层的节点都不满足条件，那么就说明没有匹配到的路由，查询结束
#### 我们实现的动态路由具备以下两个功能。
- 参数匹配: 例如`/p/:lang/doc`，可以匹配`/p/c/doc`、`/p/python/doc`和`/p/go/doc`
- 通配`*`: 例如`/static/*filepath`,可以匹配`/static/fav.ico`，也可以匹配`/static/js/jQuery.js`，这种模式常用语静态服务器，能够递归地匹配子路径
### Trie树实现
首先我们需要设计树结点上应该存储哪些信息  
```go
type node struct{
	pattern string //待匹配路由，例如 /p/:lang
	part string //路由中的一部分， 例如 :lang
	children []*node //子节点 例如 [doc, tutorial, intro]
	isWild bool //是否精确匹配, part含有: 或 * 时为true
}
```

与普通的树不同，为了实现动态路由匹配，加上了`isWild`这个参数。即当我们匹配`/p/go/doc`这个路由时，第一层节点,`p`精准
匹配到了`p`,第二层节点，`go`模糊匹配到`:lang`，那么将会把`lang`这个参数赋值为go，继续下一层匹配。我们将匹配的逻辑，包装为一个辅助函数。  
```go
//第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
    for _, child := range n.children {
        if child.part == part || child.isWild {
            return child
        }
    }
    return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
    nodes := make([]*node, 0)
    for _, child := range n.children {
        if child.part == part || child.isWild {
        nodes = append(nodes, child)
        }
    }
    return nodes
}
```

对于路由来说,最重要的当然是注册与匹配了。开发服务时，注册路由规则，映射handler；访问时，匹配路由规则，查找到对应的handler。因此，前缀树需要支持结点
的插入与查询。插入功能很简单，递归查找每一层的节点，如果没有匹配到当前`part`的节点，则新建一个。  
有一点需要注意，`p/:lang/doc`只有在第三层节点，即`doc`节点，`pattern`才会设置为`/p/:lang/doc`。
`p`和`:lang`节点的`pattern`属性皆为空。因此，当匹配结束时，我们可以使用`n.pattern == “”`来判断路由规则是否匹配成功。   
```go
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
```

### Router


### Context与handle的变化


### 单元测试


### 使用Demo

## 第四天
`
实现分组控制
`

### 分组的意义
分组控制(Group Control)是Web框架应提供的基础功能之一。所谓分组，是指路由的分组。如果没有路由分组，
我们需要针对每一个路由进行控制。但是真实的业务场景中，往往某一组路由需要相似的出路。例如  
- 以`/post`开头的路由匿名可访问
- 以`/admin`开头的路由需要鉴权
- 以`/api`开头的路由是RESTful接口，可以对接第三方平台，需要三方平台鉴权  

大部分情况下的路由分组，是以相同的前缀来区分的。因此，我们今天实现的分组控制也是以前缀来区分，并且支持分组的嵌套。
例如`/post`是一个分组，`/post/a`和`/post/b`可以是该分组下的子分组。作用在`/post`的中间件(middleware)，也
都会作用在子分组，子分组还可以应用自己特有的中间件。  

中间件可以给框架无限的扩展能力，应用在分组上，可以使得分组控制的收益更为明显，而不是共享相同的路由前缀这么简单。
例如`/admin`的分组，可以应用鉴权中间件;`/`分组应用日志中间件，`/`是默认的最顶层的分组，也就意味着给所有的路由，
即整个框架增加了记录日志的能力。

分组嵌套  
一个Group对象需要具备的属性：首先是前缀(prefix),比如`/`，或者`/api`;要支持分组嵌套，那么需要知道当前分组
的父亲(parent)是谁。当然了，按照我们一开始的分析，中间件是应用在分组上，那还需要存储应用在该分组上的中间件(middlewares)。
还记得我们之前调用函数`(*Engine).addRoute()`来映射所有的路由规则和Handler。如果Group对象需要直接映射路由规则的话，比如
我们想在使用框架时，这么调用
```go
r := dragon.New()
v1 := r.Group("/v1")
v1.GET("/", func(c *dragon.Context){
	c.HTML(http.StatusOK, "<h1>Hello Dragon</h1>")
})
```
那么Group对象，还需要有访问Ruter的能力，为了方便，我们可以在`Group`中，保存一个指针，指向`Engine`,整个框架的所有资源都是由`Engine`统一协调的
，那么就可以通过`Engine`间接地访问各种接口了。  
最后的Group的定义是这样的
```go
Router struct{
	prefix string
	middlewares []HandlerFunc
	parent *RouterGroup
	engine *Engine
}
```
我们还可以进一步地抽象，将`Engine`作为最顶层的分组，也就是说`Engine`拥有`RouterGroup`所有的能力
```go
Engine struct{
	*RouterGroup
	router *router
	groups []*RouterGroup
}
```
那我们就可以将和路由有关的函数，都交给`RouterGroup`实现了。
```go
// New is the constructor of dragon.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
```

## 第五天

### 中间件是什么
中间件(middlewares)，简单说，就是非业务的技术类组件。Web 框架本身不可能去理解所有的业务，因而不可能实现所有的功能。因此，框架需要有一个插口，允许用户自己定义功能，嵌入到框架中，仿佛这个功能是框架原生支持的一样。因此，对中间件而言，需要考虑2个比较关键的点：

插入点在哪？使用框架的人并不关心底层逻辑的具体实现，如果插入点太底层，中间件逻辑就会非常复杂。如果插入点离用户太近，那和用户直接定义一组函数，每次在 Handler 中手工调用没有多大的优势了。
中间件的输入是什么？中间件的输入，决定了扩展能力。暴露的参数太少，用户发挥空间有限。
那对于一个 Web 框架而言，中间件应该设计成什么样呢？接下来的实现，基本参考了 Gin 框架。

### 中间件设计
Dragon的中间件的定义与路由映射的Handler一致，处理的输入是Context对象。插入点是框架接收到请求初始化Context对象后，
允许用户使用自己定义的中间件做一些额外的处理，例如记录日志等，以及对Context进行二次加工。另外调用(*Context).Next()函数,
中间件可等待用户自己定义的Handler处理结束后，再做一些额外的操作。例如计算本次处理所用时间等。即Dragon的中间件支持用户在
被处理的前后，做一些额外的操作。举个例子。我们希望最终能够支持如下定义的中间件,c.Next()表示等待执行其他的中间件或用户的
Handler.
```go

```

### 代码实现

### 使用Demo



## 第六天




## 第七天