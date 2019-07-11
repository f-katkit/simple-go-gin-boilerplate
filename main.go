package main

import (
	"github.com/f-katkit/simple-go-gin-boilerplate/db"
	"github.com/f-katkit/simple-go-gin-boilerplate/server"
)

func main() {
	db.Init()
	server.Init()
}
