package main

import (
	"fmt"
	"net"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/merkle-chain/wallets/handlers"
	"github.com/merkle-chain/wallets/proto"
	"github.com/merkle-chain/wallets/store"
	"github.com/merkle-chain/wallets/wallets"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// generate seed
	hdwallet, err := wallets.GetHDWallet()

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Seed loaded")

	// connect to database
	str, err := store.NewStore()

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to database")

	port := 3000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_logrus.UnaryServerInterceptor(log.NewEntry(log.StandardLogger())),
		),
	)

	proto.RegisterWalletServiceServer(server, handlers.NewWalletService(
		hdwallet,
		str,
	))

	log.Info("Server is now accepting requests on port :", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
