package main

import (
	"go-blog/server"
)

func main() {
	r := server.NewRouter()
	r.Run(":9099")
}
