package main

import (
	"github.com/bytefly/go-micro-example/api/auth"
	"github.com/micro/micro/v2/cmd"
	"github.com/micro/micro/v2/plugin"
)

func init() {
	plugin.Register(&auth.Auth{})
}

func main() {
	cmd.Init()
}
