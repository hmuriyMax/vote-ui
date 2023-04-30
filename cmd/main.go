package main

import (
	"context"
	"log"
	uiServer "vote-ui/internal/ui-server"
)

func main() {
	ctx := context.Background()
	app := uiServer.NewServer(8080)
	err := app.Start(ctx)
	if err != nil {
		log.Fatalf("ui-server.Start: %v", err)
	}
}
