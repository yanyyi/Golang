protoc -I=proto --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-gr
pc_out=proto --go-grpc_opt=paths=source_relative --grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative  proto/hello.proto
