package main

import (
	"context"

	user "github.com/f1748x/xquser/protobuf/protobuf"
	"github.com/micro/go-micro/v2"
)

type UserService struct {
}

func (u *UserService) GetAccountRequest(ctx context.Context, req *user.GetAccountRequest, res *user.Account) error {

	res.Id = req.Id
	res.Username = "小明"
	res.Address = "杭州市100号"
	res.Phone = "130000000"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)
	service.Init()
	//user.RegisterUserServiceHandler(service.Server(),new(UserService))
	// 启动服务
	if err := service.Run(); err != nil {
		panic(err)
	}
}
