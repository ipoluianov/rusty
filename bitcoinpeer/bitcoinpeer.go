package bitcoinpeer

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/ipoluianov/rusty/logger"
)

type BitcoinPeer struct {
	id           string
	started      bool
	stopping     bool
	addresses    []string
	lastUpdateDT time.Time
	mtx          sync.Mutex
}

var BitcoinPeerInstance *BitcoinPeer

func NewBitcoinPeer() *BitcoinPeer {
	var c BitcoinPeer
	BitcoinPeerInstance = &c
	return BitcoinPeerInstance
}

func (c *BitcoinPeer) Start() {
	go c.thDnsSeedMonitoring()
}

func (c *BitcoinPeer) Stop() {
	c.stopping = true
}

func (c *BitcoinPeer) Id() string {
	return fmt.Sprint(c)
}

func (c *BitcoinPeer) Get() (dt time.Time, ips []string) {
	c.mtx.Lock()
	dt = c.lastUpdateDT
	ips = c.addresses
	c.mtx.Unlock()

	//logger.Println("processed:", c.addresses, ips)
	return
}

var dnsSeeds = []string{
	"seed.bitcoin.sipa.be",
	"dnsseed.bluematt.me",
	"dnsseed.bitcoin.dashjr.org",
	"seed.bitcoinstats.com",
	"seed.bitnodes.io",
	"bitseed.xf2.org",
}

func (c *BitcoinPeer) thDnsSeedMonitoring() {
	dtOperationTime := time.Now().Add(-100 * time.Hour)
	c.started = true

	for !c.stopping {
		for {
			if c.stopping || time.Since(dtOperationTime) > time.Duration(5*60000)*time.Millisecond {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		dtOperationTime = time.Now()

		addresses := make([]string, 0)

		for _, seed := range dnsSeeds {
			ips, err := net.LookupIP(seed)
			if err != nil {
				logger.Println("bitcoin thDnsSeedMonitoring error", err)
				continue
			}

			for _, ip := range ips {
				//logger.Println("bitcoin thDnsSeedMonitoring rcv addr:", ip.String())
				addresses = append(addresses, ip.String())
			}
		}

		c.mtx.Lock()
		c.addresses = addresses
		c.lastUpdateDT = time.Now()
		c.mtx.Unlock()

		logger.Println("processed:", c.addresses)
	}
	c.started = false
}
