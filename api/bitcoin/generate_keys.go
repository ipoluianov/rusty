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
		PrivateKeyCom   string `json:"private_key_com"`
		PrivateKeyUncom string `json:"private_key_uncom"`
		PublicKey       string `json:"public_key"`
		Address         string `json:"address"`
	}

	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyWIFCom, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	privateKeyWIFUnCom, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, false)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	var res Result
	res.PrivateKeyCom = privateKeyWIFCom.String()
	res.PrivateKeyUncom = privateKeyWIFUnCom.String()
	res.PublicKey = ""
	res.Address = ""

	utils.SendJson(w, res, nil)
}
