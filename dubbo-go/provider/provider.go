package main

import (
	"context"
	"os"

	"github.com/RPC-Snippets/dubbo-go/common"
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/apache/dubbo-go/config"
)

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
	os.Setenv("CONF_PROVIDER_FILE_PATH", "./conf/provider.yml")
	config.Load()
}
