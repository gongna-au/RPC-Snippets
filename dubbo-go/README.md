# Dubbo-go 实现RPC调用的简单示例

这是一个使用Dubbo-go实现的简单RPC的例子，包括一个服务提供者和一个服务消费者。

### 服务接口
服务接口，位于common/user.go下：

```go
type User struct {
	ID   string
	Name string
	Age  int32
}

type UserProvider struct {
	GetUser func(ctx context.Context, req []interface{}, rsp *User) error
}

func (User) JavaClassName() string {
	return "com.example.User"
}

```

### 服务提供者
服务提供者，位于provider/provider.go下：

```go


type UserProvider struct {
	// ...
}

func (u *UserProvider) Reference() string {
	return "UserProvider"
}

func (u *UserProvider) Destroy() {
	// 这里可以放一些清理逻辑，比如关闭数据库连接等。
	// 如果你没有需要清理的资源，那么这个方法可以留空。
}

func (u *UserProvider) GetUser(ctx context.Context, req []interface{}, rsp *common.User) error {
	rsp.ID = req[0].(string)
	rsp.Name = "Tom"
	rsp.Age = 23
	return nil
}

func init() {
	config.SetProviderService(&UserProvider{})
	hessian.RegisterPOJO(&common.User{})
}

func main() {
	config.Load()
}
```

Reference 方法返回的是这个服务的唯一标识符。当客户端请求这个服务时，就会用这个标识符来找到服务。

Destroy 方法是当这个服务被销毁时会被调用的。你可以在这里放一些清理逻辑，比如关闭数据库连接等。如果你没有需要清理的资源，那么这个方法可以留空。


### 服务消费者
服务消费者，位于consumer/consumer.go下：

```go
var userProvider = new(common.UserProvider)

func main() {
	hessian.RegisterPOJO(&common.User{})
	config.Load()

	user := &common.User{}
	err := userProvider.GetUser(context.Background(), []interface{}{"1"}, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
```


如果你想使用 Hessian 序列化，你需要实现 JavaClassName 方法来让 Dubbo-go 知道这个类型的 Java 类名，因为 Dubbo-go 需要它来正确地序列化和反序列化数据。

JavaClassName 方法应返回你在 Java 服务端对应的类的全限定名。比如，如果你在 Java 服务端有一个 com.example.User 类，那么在你的 Go 服务端的 User 类应该这样实现 JavaClassName 方法：

conf目录下创建服务提供者和消费者的配置文件。

服务提供者的配置文件provider_config.yml如下：

```yaml
dubbo:
  registries:
    "hangzhouzk":
      protocol: "zookeeper"
      timeout: "3s"
      address: "127.0.0.1:2181"
  provider:
    services:
      "UserProvider":
        path: "common.UserProvider"
        registry: "hangzhouzk"
        protocol: "dubbo"
        interface: "common.UserProvider"
```
服务消费者的配置文件consumer_config.yml如下：

```yaml
dubbo:
  registries:
    "hangzhouzk":
      protocol: "zookeeper"
      timeout: "3s"
      address: "127.0.0.1:2181"
  consumer:
    references:
      "UserProvider":
        registry: "hangzhouzk"
        protocol: "dubbo"
        interface: "common.UserProvider"
```
```text
dubbo: 这是 Dubbo-go 配置的根级别。
registries: 这里定义了注册中心的相关配置。
hangzhouzk: 一个名为 "hangzhouzk" 的注册中心的配置。
protocol: 注册中心使用的协议，这里是 "zookeeper"。
timeout: 与注册中心的通信超时时间，这里设置为 3 秒。
address: 注册中心的地址，这里是运行在本机的 ZooKeeper 实例。
```

```text
consumer: 这里定义了消费者（服务调用方）的全局配置。
references: 定义了消费者调用的远程服务引用。
UserProvider: 一个名为 "UserProvider" 的服务引用的配置。
registry: 此服务引用使用的注册中心的名称，这里引用了前面定义的 "hangzhouzk"。
protocol: 服务调用使用的协议，这里是 "dubbo"。
interface: 要调用的远程服务的接口名称，这里是 "common.UserProvider"。
```

本地运行了ZooKeeper注册中心，并且其地址为127.0.0.1:2181。

你可以使用 Docker 来在本地运行一个 ZooKeeper 实例。以下是运行 ZooKeeper 的一种常见方法：

首先，你需要安装 Docker。如果你还没有安装，可以访问 Docker 的官方网站(https://www.docker.com/)进行下载和安装。

打开终端，并运行以下命令，以在 Docker 中下载并运行 ZooKeeper：

```css
docker run --name my-zookeeper --restart always -d -p 2181:2181 zookeeper
```
这个命令会从 Docker Hub 下载 ZooKeeper 的最新镜像，并在名为 my-zookeeper 的 Docker 容器中运行 ZooKeeper。容器将会在后台(-d)运行，并在需要时自动重新启动(--restart always)。端口 2181 是 ZooKeeper 默认的客户端连接端口，-p 2181:2181 的配置将该端口映射到主机的 2181 端口，使得你可以在主机上通过 localhost:2181 来访问运行在 Docker 容器中的 ZooKeeper。

如果你想验证 ZooKeeper 是否正在运行，可以使用 docker ps 命令查看正在运行的 Docker 容器列表。如果你看到名为 my-zookeeper 的容器正在运行，那么就意味着 ZooKeeper 已经成功启动了。

注意，这只是一个基本的 ZooKeeper 实例，适用于开发和测试环境。在生产环境中，你可能需要使用更复杂的配置来运行 ZooKeeper，例如设置适当的内存限制，或者运行一个 ZooKeeper 集群来提高可用性。

服务提供者提供一个获取用户信息的服务，服务消费者消费这个服务。由于Dubbo-go是Dubbo的Go语言实现，所以它兼容Dubbo的所有特性，包括服务发现、负载均衡、链路追踪等。

Hessian 是一种二进制 Web 服务协议，它被设计为在跨语言环境中高效传输大型对象图。在 Hessian 协议中，序列化和反序列化的过程需要知道数据的类名。这是因为在 Hessian 的数据格式中，数据的类名是和数据一起被序列化和发送的。

Dubbo-go 使用 Hessian 作为其默认的序列化协议。因此，如果你的数据需要在 Dubbo-go 的服务之间传输，那么你需要告诉 Dubbo-go 你的数据的类名，这样 Dubbo-go 才能正确地序列化和反序列化你的数据。

当前版本的实现默认使用 Hessian 作为序列化协议。Hessian 序列化是一种二进制的跨语言序列化协议，原本设计为满足 Java 语言的跨网络对象传输，所以在序列化和反序列化的过程中，它需要知道类名信息，这是其协议的一部分。

因此，即使你的服务只是 Go 语言的服务，但如果你选择使用 Dubbo-go 默认的 Hessian 序列化，也需要为你的对象实现 JavaClassName 方法。当然，JavaClassName 返回的只是一个字符串，它可以是任意字符串，只要确保在你的服务内部保持一致即可。

如果不希望这样做，也可以更换其他的序列化协议。Dubbo-go 支持多种序列化协议，比如 json，protobuf 等。这可能需要对服务的配置做一些调整。

### 角色定位

User 是 RPC 调用的数据对象，UserProvider 是 RPC 服务的定义，而在 provider.go 文件中定义的 UserProvider 结构体是 RPC 服务的实现。

在gRPC中，每个RPC调用都定义了特定的Request和Response类型，这是基于Protocol Buffers的IDL（接口定义语言）特性。这种方式更加形式化和严格，可以清晰地定义和管理API的变更。

然而在Dubbo中，接口和服务通常是以Java接口和实现类的形式定义的，参数和返回类型更加灵活。这种方式更加便于在Java世界内进行快速开发和迭代。

在Dubbo的调用示例中，GetUser方法实际上有两个参数和一个返回值。第一个参数是context对象，用于传递调用上下文信息。第二个参数是一个interface{}数组，用于传递调用参数。在这个例子中，数组中只有一个元素，即用户的ID。返回值是一个error，表示调用是否成功。

至于user对象，它在调用之前被创建并传入GetUser方法，然后在方法内部被填充。这种方式和gRPC中返回一个Response对象的方式本质上是相同的，只是表现形式有所不同。


### 注册中心

注册中心（Registry）在微服务架构中扮演了非常重要的角色。其主要作用是服务发现和服务注册。

当一个服务提供者启动时，它会将自己提供的服务注册到注册中心，告知注册中心自己存在，可提供某种服务。注册中心会保存这些信息。当服务消费者需要调用某种服务时，它会首先从注册中心获取到提供该服务的服务提供者的详细信息，然后才能进行远程调用。

在Dubbo和Dubbo-Go中，服务提供者和消费者都需要知道注册中心的位置，因此在它们的配置文件中都需要定义registry。服务提供者需要知道在哪里注册服务，服务消费者需要知道从哪里发现服务。

而Dubbo-Go支持多注册中心，你可以在配置文件中通过列表的形式定义多个注册中心。例如，你可能有一个ZooKeeper集群作为你的主注册中心，同时又有一个备用的Nacos注册中心。你可以在你的配置文件中定义这两个注册中心，并为每个服务选择使用哪个注册中心。