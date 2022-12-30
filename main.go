package main

import (
	"go-blog/server"
	"go-blog/store"
)

func main() {
	//g.Initlog()
	store.MysqlInit()
	store.RedisInit()
	r := server.NewRouter()
	r.Run(":9099")
}
