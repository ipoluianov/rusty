package bitcoin

import (
	"net/http"

	"github.com/ipoluianov/rusty/utils"
)

func GenerateKeys(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		PrivateKey string `json:"private_key"`
		PublicKey  string `json:"public_key"`
		Address    string `json:"address"`
	}

	var res Result
	res.PrivateKey = "111"
	res.PublicKey = "111"
	res.Address = "111"

	utils.SendJson(w, res, nil)
}
