package main

import (
	"context"
	"fmt"
	"time"

	"github.com/RPC-Snippets/dubbo-go/common"
	"github.com/apache/dubbo-go/config"
)

var userProvider = new(common.UserProvider)

func main() {
	config.Load()
	time.Sleep(3e9)
	user := &common.User{}
	err := userProvider.GetUser(context.TODO(), []interface{}{"A001"}, user)
	if err != nil {
		panic(err)
	}
	println("response result: %v\n", user)
}

func println(format string, args ...interface{}) {
	fmt.Printf("\033[32;40m"+format+"\033[0m\n", args...)
}
