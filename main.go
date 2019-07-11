package main

import (
	"github.com/f-katkit/simple-gin/db"
	"github.com/f-katkit/simple-gin/server"
)

func main() {
	db.Init()
	server.Init()
}
