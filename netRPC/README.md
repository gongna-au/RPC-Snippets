## net-RPC 实现RPC调用的简单示例

### Server

服务器代码首先定义了一个Calculator类型，它有一个Square方法，该方法可以通过RPC调用。然后我们创建一个新的Calculator对象，并使用rpc.Register函数注册这个对象。服务器监听在1234端口，使用rpc.Accept函数来接收客户端的连接和请求。

### Client

客户端代码首先与服务器建立一个连接，然后调用Calculator的Square方法，并打印出结果。在客户端调用Square方法时，Go的RPC库会将这个方法调用编码为一个请求，然后发送到服务器。服务器收到请求后，找到对应的Calculator对象和Square方法，用请求中的参数调用这个方法，然后将结果编码为一个响应发送回客户端。客户端收到响应后，将响应的结果解码，并返回给调用者。