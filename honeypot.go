package main

import (
	"fmt"

	"github.com/chennqqi/go-HoneyPot/config"
	"github.com/chennqqi/go-HoneyPot/tcp"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("read config error:", err)
		return
	}

	tcpServer, err := tcp.NewServer(&cfg)
	if err != nil {
		fmt.Println("NewServer error:", err)
		return
	}
	tcpServer.Run()
}
