package main

import (
	"log"

	"p2p-messenger/internal/ui"
)

func main() {
	app := ui.NewApp()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
