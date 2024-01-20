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
		AddressP2PKHHex       string `json:"address_p2pkh_hex"`
		AddressP2PKH          string `json:"address_p2pkh"`
		AddressBech32         string `json:"address_bech32"`
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

	serializedPubKeyCompressed := publicKey.SerializeCompressed()
	serializedPubKeyUncompressed := publicKey.SerializeUncompressed()

	serializedPubKeyCompressedHex := hex.EncodeToString(serializedPubKeyCompressed)
	serializedPubKeyUncompressedHex := hex.EncodeToString(serializedPubKeyUncompressed)

	p2pkhAddress, err := btcutil.NewAddressPubKey(publicKey.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	witnessProg := btcutil.Hash160(serializedPubKeyCompressed)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		utils.SendError(w, err)
		return
	}
	bech32Address := addressWitnessPubKeyHash.EncodeAddress()

	var res Result
	res.PrivateKeyCom = privateKeyWIFCom.String()
	res.PrivateKeyUncom = privateKeyWIFUnCom.String()
	res.PublicKeyCompressed = serializedPubKeyCompressedHex
	res.PublicKeyUncompressed = serializedPubKeyUncompressedHex
	res.AddressP2PKHHex = p2pkhAddress.String()
	res.AddressP2PKH = p2pkhAddress.EncodeAddress()
	res.AddressBech32 = bech32Address

	utils.SendJson(w, res, nil)
}
