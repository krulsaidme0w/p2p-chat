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
	runNetworkManager()

	if err := runUI(); err != nil {
		log.Fatal(err)
	}
}

func runNetworkManager() {
	p := proto.NewProto(faker.Name())

	networkManager := network.NewManager(p)
	networkManager.Start()
}

func runUI() error {
	app := ui.NewApp()
	if err := tview.NewApplication().SetRoot(app.View, true).SetFocus(app.Sidebar.View).Run(); err != nil {
		return err
	}
	return nil
}
