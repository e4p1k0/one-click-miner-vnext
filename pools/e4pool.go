package pools

import (
	"fmt"
	"time"

	"github.com/vertcoin-project/one-click-miner-vnext/util"
)

var _ Pool = &E4pool{}

type E4pool struct {
	Address           string
	LastFetchedPayout time.Time
	LastPayout        uint64
}

func NewE4pool(addr string) *E4pool {
	return &E4pool{Address: addr}
}

func (p *E4pool) GetPendingPayout() uint64 {
	jsonPayload := map[string]interface{}{}
	err := util.GetJson(fmt.Sprintf("https://mcore.e4pool.com/api/pools/vtc/miners/%s", p.Address), &jsonPayload)
	if err != nil {
		return 0
	}
	vtc, ok := jsonPayload["pendingBalance"].(float64)
	if !ok {
		return 0
	}
	vtc *= 100000000
	return uint64(vtc)
}

func (p *E4pool) GetStratumUrl() string {
	return "stratum+tcp://vtc.e4pool.com:9112"
}

func (p *E4pool) GetUsername() string {
	return p.Address
}

func (p *E4pool) GetPassword() string {
	return "x"
}

func (p *E4pool) GetID() int {
	return 8
}

func (p *E4pool) GetName() string {
	return "e4pool.com"
}

func (p *E4pool) GetFee() float64 {
	return 1.00
}

func (p *E4pool) OpenBrowserPayoutInfo(addr string) {
	util.OpenBrowser(fmt.Sprintf("https://e4pool.com/vtc/account/%s", addr))
}
