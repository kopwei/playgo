package main

import (
	"fmt"
	"github.com/kopwei/playgo/playgrpc/server"
)

func main() {
	fmt.Println("Hello Playgo!")
	server.StartServer()
}
