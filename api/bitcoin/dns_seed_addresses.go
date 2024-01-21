package bitcoin

import (
	"net/http"

	"github.com/ipoluianov/rusty/bitcoinpeer"
	"github.com/ipoluianov/rusty/utils"
)

func DNSSeedAddresses(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		IPs      []string `json:"ips"`
		UpdateDT string   `json:"update_dt"`
	}

	var res Result
	dt, ips := bitcoinpeer.BitcoinPeerInstance.Get()
	res.UpdateDT = dt.Format("2006-01-02 15:04:05")
	res.IPs = ips

	utils.SendJson(w, res, nil)
}
