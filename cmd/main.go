package main

import (
	"flag"
	"log"

	"p2p-messenger/internal/network"
	"p2p-messenger/internal/proto"
	"p2p-messenger/internal/ui"
)

const (
	Port = "25042"
)

func main() {
	name := flag.String("name", "krulsaidme0w", "peers name")
	flag.Parse()

	p := proto.NewProto(*name, Port)

	runNetworkManager(p)

	if err := runUI(p); err != nil {
		log.Fatal(err)
	}
}

func runNetworkManager(p *proto.Proto) {
	networkManager := network.NewManager(p)
	networkManager.Start()
}

func runUI(p *proto.Proto) error {
	return ui.NewApp(p).Run()
}
