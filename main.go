package main

import (
	"go-blog/server"
	"go-blog/store"
)

func main() {
	store.MysqlInit()
	store.InitRedis()
	r := server.NewRouter()
	r.Run(":9099")
}
