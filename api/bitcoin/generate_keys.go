package bitcoin

import (
	"encoding/hex"
	"log"
	"net/http"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ipoluianov/rusty/utils"
)

func GenerateKeys(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		PrivateKey string `json:"private_key"`
		PublicKey  string `json:"public_key"`
		Address    string `json:"address"`
	}

	// Генерация приватного ключа
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Получение публичного ключа из приватного
	publicKey := privateKey.PubKey()

	// Генерация Bitcoin-адреса
	address, err := btcutil.NewAddressPubKey(publicKey.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}

	var res Result
	res.PrivateKey = hex.EncodeToString(privateKey.Serialize())
	res.PublicKey = hex.EncodeToString(publicKey.SerializeUncompressed())
	res.Address = address.String()

	utils.SendJson(w, res, nil)
}
