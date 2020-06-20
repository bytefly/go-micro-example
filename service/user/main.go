package main

import (
	"context"
	"encoding/json"
	"github.com/bytefly/go-micro-example/service/constant/micro_c"
	"github.com/bytefly/go-micro-example/service/user/proto"
	userApi "github.com/bytefly/go-micro-example/service/user/proto"
	"github.com/bytefly/go-micro-example/service/user/service"
	"github.com/bytefly/go-micro-example/service/util"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/metadata"
	"log"
)

type UserService struct {
}

func (us *UserService) GetUserInfo(ctx context.Context, req *userApi.Empty, rsp *userApi.UserInfo) error {
	log.Println("Received User.GetUserInfo RPC request")
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.Unauthorized(micro_c.MicroNameUser, "no auth meta-data found in request")
	}
	rsp.Id = meta["X-Example-Id"]
	rsp.Username = meta["X-Example-Username"]
	rsp.Password = "password from db"
	return nil
}

func (us *UserService) Login(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	var userInfo *user.UserInfo
	json.Unmarshal([]byte(req.Body), &userInfo)
	log.Println("Received User.Login API request with: ", userInfo)
	response, code, err := service.NewUserService().Login(userInfo)
	return util.Resp(code, err, rsp, response)
}

func main() {
	userService := micro.NewService(
		micro.Name(micro_c.MicroNameUser),
	)
	userService.Init()
	userApi.RegisterUserHandler(userService.Server(), &UserService{})
	if err := userService.Run(); err != nil {
		log.Fatal(err)
	}
}
