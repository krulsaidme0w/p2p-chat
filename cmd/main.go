package main

import (
	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/rivo/tview"

	"p2p-messenger/internal/network"
	"p2p-messenger/internal/proto"
	"p2p-messenger/internal/ui"
)

func main() {
	p := proto.NewProto(faker.Name())

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
	app := ui.NewApp(p)
	if err := tview.NewApplication().SetRoot(app.View, true).SetFocus(app.Sidebar.View).Run(); err != nil {
		return err
	}

	return nil
}
