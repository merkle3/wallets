package wallets

import (
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"strconv"
)

type RawWallet struct {
	account *accounts.Account
}

// get the n-th wallet from the seed
func (s *HDWallet) GetWallet(n int64) (*RawWallet, error) {
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + strconv.FormatInt(n, 10))

	wallet, err := hdwallet.NewFromSeed(s.seed)

	if err != nil {
		return nil, err
	}

	account, err := wallet.Derive(path, false)

	if err != nil {
		return nil, err
	}

	return &RawWallet{
		account: &account,
	}, nil
}

func (a *RawWallet) GetAddress() string {
	return a.account.Address.String()
}
