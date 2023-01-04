package main

import (
	"go-blog/server"
	"go-blog/store"
)

func main() {
	r := server.NewRouter()
	store.RedisInit()
	store.MysqlInit()
	r.Run(":9099")
}
