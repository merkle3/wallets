package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type WalletStore struct {
	db *mongo.Client
}

func NewStore() (*WalletStore, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/wallets"))

	if err != nil {
		return nil, err
	}

	// try to connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	return &WalletStore{
		db: client,
	}, nil
}

func (s *WalletStore) GetDatabase() *mongo.Database {
	return s.db.Database("wallet_db")
}
