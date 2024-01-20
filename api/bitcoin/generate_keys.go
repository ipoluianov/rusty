package bitcoin

import (
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

	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyWIF, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, false)
	if err != nil {
		log.Fatal(err)
	}

	var res Result
	res.PrivateKey = privateKeyWIF.String()
	res.PublicKey = ""
	res.Address = ""

	utils.SendJson(w, res, nil)
}
