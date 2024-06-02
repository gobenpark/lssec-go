package lssec_go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRealtime(t *testing.T) {
	cli := ClientHelper(t)
	ch, err := cli.Subscribe(context.TODO(), SubscriptionContent{
		Type:   AddPriceTRType,
		TRCD:   "SHC",
		Ticker: "1",
	}, SubscriptionContent{
		Type:   AddPriceTRType,
		TRCD:   "OK_",
		Ticker: "027740",
	})
	require.NoError(t, err)
	for i := range ch {
		fmt.Println(string(i))
	}
}
