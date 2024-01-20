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
		PrivateKeyCom         string `json:"private_key_c"`
		PrivateKeyUncom       string `json:"private_key_u"`
		PublicKeyCompressed   string `json:"public_key_compressed"`
		PublicKeyUncompressed string `json:"public_key_uncompressed"`
		Address               string `json:"address"`
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

	publicKey := privateKey.PubKey()

	serializedPubKeyCompressed := publicKey.SerializeCompressed()     // Сжатый формат
	serializedPubKeyUncompressed := publicKey.SerializeUncompressed() // Несжатый формат

	serializedPubKeyCompressedHex := hex.EncodeToString(serializedPubKeyCompressed)
	serializedPubKeyUncompressedHex := hex.EncodeToString(serializedPubKeyUncompressed)

	var res Result
	res.PrivateKeyCom = privateKeyWIFCom.String()
	res.PrivateKeyUncom = privateKeyWIFUnCom.String()
	res.PublicKeyCompressed = serializedPubKeyCompressedHex
	res.PublicKeyUncompressed = serializedPubKeyUncompressedHex
	res.Address = ""

	utils.SendJson(w, res, nil)
}
