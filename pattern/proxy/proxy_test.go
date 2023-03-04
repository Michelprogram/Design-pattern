package proxy

import (
	"context"
	"testing"
	"time"
)

func TestProxy(t *testing.T) {
	ctx := context.Background()
	proxy := NewProxyCryptoFinance()

	time := time.Now().Add(time.Second * 2)
	timeout, _ := context.WithDeadline(ctx, time)

	proxy.ListCrypto()

	if timeout.Err() != context.DeadlineExceeded {
		t.Errorf("Expected %v, got %v", context.DeadlineExceeded, timeout.Err())
	}
}

func TestProxyCached(t *testing.T) {
	ctx := context.Background()
	proxy := NewProxyCryptoFinance()

	proxy.ListCrypto()

	time := time.Now().Add(time.Second * 2)
	timeout, _ := context.WithDeadline(ctx, time)

	proxy.ListCrypto()

	if timeout.Err() == context.DeadlineExceeded {
		t.Errorf("Expected %v, got %v", context.DeadlineExceeded, timeout.Err())
	}
}

func TestProxyGetInfo(t *testing.T) {
	proxy := NewProxyCryptoFinance()

	proxy.GetInfoCrypto(1)

	_, exist := proxy.CryptoCached[1]

	if !exist {
		t.Errorf("Expected %v, got %v", true, exist)
	}

}
