package main

import (
	_ "go-blog/src/model"
	"go-blog/src/server"
)

func main() {
	r := server.NewRouter()
	r.Run(":9099")
}

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)
