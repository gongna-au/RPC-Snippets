# gRPC实现RPC调用的简单示例

### 生成pb.go文件

```shell
protoc -I . --go_out=plugins=grpc:. calculator.proto
```

### 运行

运行server.go来启动服务器，并运行client.go来调用RPC方法并打印结果。

这个例子展示了如何使用gRPC定义、实现和调用一个简单的RPC服务。gRPC使用了Protocol Buffers作为接口定义语言，并且可以生成客户端和服务器的代码，以便在不同的平台和语言之间进行通信。