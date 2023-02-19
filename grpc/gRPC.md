gRPC是什么？
1. gRPC来自Google,它是一个高性能、开源、和通用的RPC框架,支持多种语言
2. 基于IDL(接口定义语言(Interface Definition Language))文件定义服务,通过proto3工具生成指定语言的数据结构、服务端接口以及客户端Stub
3. 通信协议基于标准的HTTP/2设计,支持双向流、消息头压缩、单TCP的多路复用、服务端推送等特性
4. 序列化支持PB(Protocol Buffer)和JSON,PB是一种与语言无关的高性能序列化框架,基于HTTP/2+PB,保障了RPC调用的高性能

gRPC请求流程
1. 客户端(gRPC Stub)调用方法,发起RPC调用
2. 对请求信息使用Protobuf进行对象序列化压缩
3. 服务端(gRPC Server)接收到请求后,解码请求体,进行业务逻辑处理并返回
4. 对响应结果使用Protobuf进行对象序列化压缩
5. 客户端接收到服务端响应,解码请求体。回调被调用的A方法,唤醒正在等待响应(阻塞)的客户端调用并返回响应结果

gRPC使用流程
1. 定义标准的proto文件
2. 通过Proto工具生成标准的代码,代码包括.pb.go和grpc.pb.go
3. 服务端使用生成的代码提供服务
4. 客户端使用生成的代码提供服务