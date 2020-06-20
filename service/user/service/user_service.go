package service

import (
	"errors"
	"github.com/bytefly/go-micro-example/api/auth"
	"github.com/bytefly/go-micro-example/service/constant/code"
	"github.com/bytefly/go-micro-example/service/user/dto"
	"github.com/bytefly/go-micro-example/service/user/proto"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) Login(req *user.UserInfo) (*dto.LoginResponse, int32, error) {
	if req == nil || req.Id == "" || req.Username == "" || req.Password == "" {
		return nil, code.InvalidParam, errors.New("param invalid")
	}

	token, e := auth.Encode(req)
	if e != nil {
		return nil, code.JwtEncodeError, e
	}

	return &dto.LoginResponse{
		Token: token,
	}, code.OK, nil
}
