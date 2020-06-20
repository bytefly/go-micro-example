package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytefly/go-micro-example/hystrix"
	. "github.com/bytefly/go-micro-example/service/config"
	"github.com/bytefly/go-micro-example/service/constant/micro_c"
	"github.com/bytefly/go-micro-example/service/greeter/dto"
	greeterApi "github.com/bytefly/go-micro-example/service/greeter/proto"
	"github.com/bytefly/go-micro-example/service/greeter/service"
	"github.com/bytefly/go-micro-example/service/user/proto"
	"github.com/bytefly/go-micro-example/service/util"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/api/proto"
	"log"
)

type Greeter struct {
	userClient user.UserService
}

func (this *Greeter) Hello(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	log.Println("Received Greeter.Hello API request")
	var helloRequest *dto.HelloRequest
	json.Unmarshal([]byte(req.Body), &helloRequest)
	response, code, err := service.NewGreeterService().Greeter(ctx, this.userClient, helloRequest)
	return util.Resp(code, err, rsp, response)
}

func main() {
	hystrix.Configure([]string{"go.micro.api.user.User.GetUserInfo"})
	greeterService := micro.NewService(
		micro.Name(micro_c.MicroNameGreeter),
		micro.WrapClient(hystrix.NewClientWrapper()),
		micro.Flags(
			&cli.StringFlag{
				Name:  "prof",
				Usage: "Running environment, eg: test, prod",
			},
		),
	)
	greeterService.Init()

	greeterService.Init(
		micro.Action(func(c *cli.Context) error {
			profile := c.String("prof")
			if len(profile) > 0 {
				// http://config-server:8081/greeter-prod.yml
				LocalConfig = GetConfig(micro_c.MicroConfigService, "greeter", profile)
				fmt.Printf("config loaded from config-server is: %s\n", LocalConfig)
			}
			return nil
		}))

	greeterApi.RegisterGreeterHandler(greeterService.Server(), &Greeter{
		userClient: user.NewUserService(micro_c.MicroNameUser, greeterService.Client())})

	if err := greeterService.Run(); err != nil {
		log.Fatal(err)
	}
}
