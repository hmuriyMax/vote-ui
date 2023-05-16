package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	auth "vote-ui/internal/services/auth-service"
	vote "vote-ui/internal/services/vote-service"
	uiServer "vote-ui/internal/ui-server"
)

func main() {
	ctx := context.Background()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	voteServer := vote.NewVoteService(conn)
	authServer := auth.NewAuthService(conn)

	err = voteServer.InitPublicKey(ctx)
	if err != nil {
		log.Fatalf("fail to init public key: %v", err)
	}

	app := uiServer.NewServer(8080, voteServer, authServer)

	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("ui-server.Start: %v", err)
	}
}
