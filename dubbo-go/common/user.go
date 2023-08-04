package common

import "context"

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
