package proxy

import (
	"time"
)

type CryptoFinance interface {
	ListCrypto() []string
	GetInfoCrypto(id int) string
}

type ServiceCryptoFinance struct {
	crypto map[int]string
}

func NewServiceCryptoFinance() *ServiceCryptoFinance {
	crypto := map[int]string{1: "BTC", 2: "ETH", 3: "XRP"}

	return &ServiceCryptoFinance{crypto: crypto}
}

func (s *ServiceCryptoFinance) ListCrypto() []string {
	time.Sleep(time.Second * 2) // Simulate a long process
	values := make([]string, len(s.crypto))

	for _, v := range s.crypto {
		values = append(values, v)
	}

	return values
}

func (s *ServiceCryptoFinance) GetInfoCrypto(id int) string {
	time.Sleep(time.Second * 2) // Simulate a long process
	return s.crypto[id]
}

type ProxyCryptoFinance struct {
	crypto       []string
	CryptoCached map[int]string
	service      *ServiceCryptoFinance
}

func NewProxyCryptoFinance() *ProxyCryptoFinance {
	return &ProxyCryptoFinance{
		service:      NewServiceCryptoFinance(),
		crypto:       []string{},
		CryptoCached: make(map[int]string, 3),
	}
}

func (p *ProxyCryptoFinance) ListCrypto() []string {
	//Simulate a cache
	if len(p.crypto) == 0 {
		p.crypto = p.service.ListCrypto()
	}
	return p.crypto
}

func (p *ProxyCryptoFinance) GetInfoCrypto(id int) string {

	//Simulate a cache

	value, exist := p.CryptoCached[id]

	if !exist {
		p.CryptoCached[id] = p.service.GetInfoCrypto(id)
	}
	return value
}
