package lssec_go

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gobenpark/lssec-go/test"
	"github.com/stretchr/testify/require"
)

func ClientHelper(t *testing.T) *Client {
	t.Helper()
	key, secret := test.Secret(t)
	cli := NewClient(
		WithAuth(key, secret), WithAutomaticTokenCache(true), WithDebug(true))
	return cli
}

func Test_ClientRealtime(t *testing.T) {
	cli := ClientHelper(t)

	contentskq := SubscriptionContent{
		Type:   AddPriceTRType,
		TRCD:   KOSDAQContract,
		Ticker: "063170",
	}
	cch, err := cli.Subscribe(context.TODO(), contentskq)
	require.NoError(t, err)

	go func() {
		for i := range cch {
			fmt.Println("kd: " + string(i))
		}
	}()
	time.Sleep(10 * time.Hour)
}
