package main

import (
	"context"
	user "github.com/f1748x/microItem/protobuf/protobuf"
)

type UserService struct {

}

func (u *UserService) GetAccountRequest(ctx context.Context,req *user.GetAccountRequest,res *user.Account)  {

}