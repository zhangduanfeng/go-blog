package main

import (
	"go-blog/server"
	"go-blog/store"
)

func main() {
	r := server.NewRouter()
	store.MysqlInit()
	store.RedisInit()
	r.Run(":9099")
}
