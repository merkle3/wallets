package wallets

import (
	"errors"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

const SEED_FILE_NAME = "seed"

type HDWallet struct {
	seed []byte
}

// checks if the seed file exists
func isSeedFilePresent() bool {
	if _, err := os.Stat(SEED_FILE_NAME); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// write a need seed master to file
func generateSeedFile() error {
	seed, err := hdwallet.NewSeed()

	if err != nil {
		return err
	}

	file, err := os.Create(SEED_FILE_NAME)

	if err != nil {
		return errors.New("could not create seed file: " + err.Error())
	}

	b, err := file.Write(seed)

	if err != nil {
		return errors.New("could not write seed file: " + err.Error())
	}

	log.Info("Wrote seed file (", b, " bytes)")

	return nil
}

func loadSeedFile() ([]byte, error) {
	file, err := os.Open(SEED_FILE_NAME)

	if err != nil {
		return nil, errors.New("could not open seed file: " + err.Error())
	}

	seed := make([]byte, 32)
	readBytes, err := file.Read(seed)

	if err != nil {
		return nil, errors.New("could not read seed file: " + err.Error())
	}

	if readBytes != 32 {
		return nil, errors.New("expected a 32 bytes (256 bit) seed key, got " + strconv.FormatInt(int64(readBytes), 10) + " bytes")
	}

	return seed, nil
}

// gets the seed
func GetHDWallet() (*HDWallet, error) {
	if !isSeedFilePresent() {
		log.Info("Seed file missing, generating a new one")
		err := generateSeedFile()

		if err != nil {
			return nil, err
		}
	}

	seed, err := loadSeedFile()

	if err != nil {
		return nil, err
	}

	return &HDWallet{
		seed: seed,
	}, nil
}
