# go-kit分层

最内层的服务(service)域是所有内容都基于特定服务定义的地方,也是实现所有业务逻辑的地方

中间端点(endpoint)域是将服务的每个方法抽象为通用的`endpoint.Endpoint`以及实现安全性和抗脆弱性逻辑的位置

最外层的传输(transport)域是端点绑定到HTTP或gRPC等具体传输的地方