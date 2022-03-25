package handlers

import (
	"context"
	"github.com/merkle-chain/wallets/proto"
	"github.com/merkle-chain/wallets/store"
	"github.com/merkle-chain/wallets/wallets"
	log "github.com/sirupsen/logrus"
)

type WalletService struct {
	hdwallet *wallets.HDWallet
	store    *store.WalletStore

	proto.UnimplementedWalletServiceServer
}

func NewWalletService(wallet *wallets.HDWallet, str *store.WalletStore) *WalletService {
	return &WalletService{
		hdwallet: wallet,
		store:    str,
	}
}

func (w *WalletService) GetWallets(ctx context.Context, in *proto.GetWalletsRequest) (*proto.GetWalletsResponse, error) {
	wallets, err := w.store.GetWallets(ctx)

	if err != nil {
		return nil, err
	}

	log.Info("Found ", len(wallets), " wallets")

	walletsResponse := make([]*proto.Wallet, len(wallets))

	for i, wallet := range wallets {
		walletsResponse[i] = &proto.Wallet{
			Name:          wallet.Name,
			Id:            wallet.ID.Hex(),
			DerivationKey: int64(wallet.Derivation),
			Address:       wallet.Address,
			CreatedAt:     wallet.CreatedAt.T,
		}
	}

	log.Info("Returnings ", len(walletsResponse), " wallets")

	return &proto.GetWalletsResponse{
		Wallets: walletsResponse,
	}, nil
}

func (w *WalletService) CreateWallet(ctx context.Context, in *proto.CreateWalletRequest) (*proto.Wallet, error) {
	newWallet, err := w.store.CreateWallet(ctx)

	if err != nil {
		return nil, err
	}

	return &proto.Wallet{
		Id:            newWallet.ID.String(),
		Name:          newWallet.Name,
		DerivationKey: int64(newWallet.Derivation),
		Address:       newWallet.Address,
		CreatedAt:     newWallet.CreatedAt.I,
	}, nil
}
