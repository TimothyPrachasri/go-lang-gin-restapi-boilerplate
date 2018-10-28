package service

import (
	"log"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

type Stellar struct{}

func (c Stellar) LogBalances(addresses [2]string) {
	for _, address := range addresses {
		account, err := horizon.DefaultTestNetClient.LoadAccount(address)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Balances for address:", address)
		for _, balance := range account.Balances {
			log.Println(balance)
		}
	}
}

func (c Stellar) SendLumens(amount string, sourcePrivateKey string, destinationPublicKey string) (bool, error) {
	if _, err := horizon.DefaultTestNetClient.LoadAccount(destinationPublicKey); err != nil {
		return false, err
	}

	tx, err := build.Transaction(
		build.TestNetwork,
		build.SourceAccount{sourcePrivateKey},
		build.AutoSequence{horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{destinationPublicKey},
			build.NativeAmount{amount},
		),
	)

	if err != nil {
		return false, err
	}

	txe, err := tx.Sign(sourcePrivateKey)
	if err != nil {
		return false, err
	}

	txeB64, err := txe.Base64()
	if err != nil {
		return false, err
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		return false, err
	}

	log.Println("Successfully sent", amount, "lumens to", destinationPublicKey, ". Hash:", resp.Hash)
	return true, nil
}
