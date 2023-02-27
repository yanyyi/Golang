### 大规模可扩展的Web应用通常需要具备以下特质:
- 可扩展
- 模块化
- 可维护
- 高性能


### 一个程序只需要满足以下两个条件,我们就可以把它看作是一个Web应用:
- 这个程序必须向发送命令请求的客户端返回HTML，而客户端则会向用户展示渲染后的HTML
- 这个程序在向客户端传送数据时必须使用HTTP协议


### HTTP简介
HTTP是万维网的应用层通信协议,Web页面中的所有数据都是通过这个看似简单的文本协议进行传输的。HTTP非常朴素,但却异常强大——这个协议自20世纪90年代定义以来,至今只进行过3次迭代修改,其中HTTP 1.1是目前使用最为广泛的一个版本,而最新的一个版本则是HTTP2.0,又称HTTP/2。  
- HTTP是一种无状态的、由文本构成的请求-响应协议(request-response),这种协议使用的是客户端-服务器(client-server)计算模型。


### HTTP请求  
(1) 请求行(request-line)  
(2) 零个或任意多个请求首部(header)  
(3) 一个空行  
(4) 可选的报文主体(body)
1. 请求方法  
- GET HEAD POST OUT DELETE TRACE OPTIONS CONNECT PATCH
2. 安全的请求方法  
如果一个HTTP方法只要求服务器提供信息而不会对服务器的状态做任何修改,那么这个方法就是安全的
- GET HEAD OPTIONS TRACE是  安全的方法
- POST PUT DELETE  是不安全的方法
3. 幂等的请求方法  
如果一个HTTP方法在使用相同的数据进行第二次调用的时候,不会对服务器的状态造成任何改变,那么这个方法就是幂等的(idempotent)。所有安全的方法都是幂等的,PUT和DELETE也是幂等的。POST既非安全也非幂等。
4. 浏览器对请求方法的支持  
HTML不支持除GET和POST之外的其他HTTP方法。话虽如此,流行的浏览器通常不会只支持HTML一种数据格式——用户可以使用XMLHttpRequest(XHR)来获得对PUT方法和DELETE方法的支持。XHR是一系列浏览器API,这些API通常由JavaScript包裹。
5. 请求首部  
HTTP请求的首部记录了与请求本身以及客户端有关的信息。大多数HTTP请求首部都是可选的,宿主(Host)首部字段是HTTP1.1唯一强制要求的首部。  
常见的HTTP请求首部:  
- Accept   Accept-Charset   Authorization   Cookie   Content-Length   Content-Type   Host   Referrer   UserAgent


### HTTP响应  
(1) 一个状态行  
(2) 零个或任意多个响应首部(header)  
(3) 一个空行  
(4) 可选的报文主体(body)  
响应状态码:
- 1XX: 情报状态码
- 2XX: 成功状态码
- 3XX: 重定向状态码
- 4XX: 客户端状态错误码
- 5XX: 服务器状态错误码
- 常见的HTTP响应首部:
- Allow   Content-Length   Content-Type   Date   Location   Server   Set-Cookie   WWW-Authenticate

### URI
URI包含了URN(统一资源标识符)和URL(统一资源定位符)  
URI的一般格式为:  
- <方案名称>:<分层部分>[ ？<查询参数>] [ # <片段>]
- 方案名称:大多数情况下只会使用HTTP方案  
一个使用HTTP方案的的URI示例:  
http://sausheong:password@www.example.com/docs/file?name=sausheong&location=singapore#summary  
- http:方案名称
- sausheong:password:记录用户名和密码
- www.example.com/docs/file: 分层部分的其他部分


### HTTP/2  
与使用纯文本方式表示的HTTP 1.x不同, HTTP/2是一种二进制协议

### Web应用的各个组成部分  
Web应用就是一个执行以下任务的程序:  
(1) 通过HTTP协议,以HTTP请求报文的形式获取客户端输入  
(2) 对HTTP请求报文进行处理,并执行必要的操作  
(3) 生成HTML，并以HTTP响应报文的形式将其返回给客户端  
为完成这些任务,Web应用被分成了处理器(handler)和模板引擎(template engine)这两个部分

