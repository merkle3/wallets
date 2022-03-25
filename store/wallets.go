package store

import (
	"context"
	"errors"
	"github.com/merkle-chain/wallets/models"
	"github.com/merkle-chain/wallets/wallets"
	"go.mongodb.org/mongo-driver/bson"
)

const WALLETS_COLLECTION = "wallets"

func (s *WalletStore) GetWallets(ctx context.Context) ([]models.Wallet, error) {
	walletCollection := s.GetDatabase().Collection(WALLETS_COLLECTION)

	var wallets []models.Wallet

	cursor, err := walletCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &wallets)

	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func (s *WalletStore) CreateWallet(ctx context.Context) (*models.Wallet, error) {
	hd, err := wallets.GetHDWallet()

	if err != nil {
		return nil, err
	}

	count, err := s.GetDatabase().Collection(WALLETS_COLLECTION).CountDocuments(ctx, bson.M{})

	if err != nil {
		return nil, errors.New("could not get wallet count: " + err.Error())
	}

	wallet, err := hd.GetWallet(count + 1)

	if err != nil {
		return nil, errors.New("could not generate new wallet: " + err.Error())
	}

	newWallet := &models.Wallet{
		Derivation: int(count),
		Address:    wallet.GetAddress(),
	}

	r, err := s.GetDatabase().Collection(WALLETS_COLLECTION).InsertOne(ctx, newWallet)

	if err != nil {
		return nil, errors.New("could not insert wallet: " + err.Error())
	}

	result := s.GetDatabase().Collection(WALLETS_COLLECTION).FindOne(ctx, bson.M{
		"_id": r.InsertedID,
	})

	insertedWallet := models.Wallet{}
	result.Decode(&insertedWallet)

	return &insertedWallet, nil
}
