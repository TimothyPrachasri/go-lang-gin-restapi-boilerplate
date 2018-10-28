package controller

import (
	"net/http"

	service "app/lib/domain/services/stellar"

	"github.com/gin-gonic/gin"
)

func Transfer(c *gin.Context) {
	sourceSeed := "SDKJ2BUKQ5TCMSLRQBAFSEVJ3LBXFGHEKKPTYNCDWSOJ4CFGFR5SKRME"
	sourceAddr := "GAMMEG4YSIGJO3I45T37A57EN5CFQQH5OVHO4XDUGF7E32VJ76YYX56E"
	destinationAddr := "GCICVEBF5JYDBCTR3TXFGN56WGYBAKKWVHUQYPM72F6ZEQ7BDQZT4NFZ"
	addresses := [2]string{sourceAddr, destinationAddr}
	service.Stellar{}.LogBalances(addresses)
	status, err := service.Stellar{}.SendLumens("100", sourceSeed, destinationAddr)
	service.Stellar{}.LogBalances(addresses)
	if status {
		c.JSON(http.StatusOK, gin.H{
			"status": status,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"err": err,
	})
}

// func logBalances(addresses [2]string) {
// 	for _, address := range addresses {
// 		account, err := horizon.DefaultTestNetClient.LoadAccount(address)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Println("Balances for address:", address)
// 		for _, balance := range account.Balances {
// 			log.Println(balance)
// 		}
// 	}
// }

// func sendLumens(amount string, sourcePrivateKey string, destinationPublicKey string) (bool, error) {
// 	if _, err := horizon.DefaultTestNetClient.LoadAccount(destinationPublicKey); err != nil {
// 		return false, err
// 	}

// 	tx, err := build.Transaction(
// 		build.TestNetwork,
// 		build.SourceAccount{sourcePrivateKey},
// 		build.AutoSequence{horizon.DefaultTestNetClient},
// 		build.Payment(
// 			build.Destination{destinationPublicKey},
// 			build.NativeAmount{amount},
// 		),
// 	)

// 	fmt.Println(tx, "tx")

// 	if err != nil {
// 		return false, err
// 	}

// 	txe, err := tx.Sign(sourcePrivateKey)
// 	if err != nil {
// 		return false, err
// 	}

// 	txeB64, err := txe.Base64()
// 	if err != nil {
// 		return false, err
// 	}

// 	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
// 	if err != nil {
// 		return false, err
// 	}

// 	log.Println("Successfully sent", amount, "lumens to", destinationPublicKey, ". Hash:", resp.Hash)
// 	return true, nil
// }
