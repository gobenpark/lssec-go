package ebest_go

import (
	"context"
	"ebest-go/test"
	"fmt"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func Test_ClientRealtime(t *testing.T) {
	key, secret := test.Secret(t)
	cli := NewClient(
		WithAuth(key, secret), WithAutomaticTokenCache(true))

	ko, err := cli.Kospi(context.Background())
	require.NoError(t, err)

	contents := lo.Map(ko, func(item Code, index int) SubscriptionContent {
		return SubscriptionContent{
			Type:   AddPriceTRType,
			TRCD:   KOSPITR,
			Ticker: item.Code,
		}
	})

	ch, err := cli.Subscribe(context.TODO(), contents...)
	require.NoError(t, err)
	go func() {

		for i := range ch {
			fmt.Println("kp: " + string(i))
		}
	}()

	kq, err := cli.Kosdaq(context.Background())
	require.NoError(t, err)

	contentskq := lo.Map(kq, func(item Code, index int) SubscriptionContent {
		return SubscriptionContent{
			Type:   AddPriceTRType,
			TRCD:   KOSDAQTR,
			Ticker: item.Code,
		}
	})
	cch, err := cli.Subscribe(context.TODO(), contentskq...)
	require.NoError(t, err)

	go func() {
		for i := range cch {
			fmt.Println("kd: " + string(i))
		}
	}()
	time.Sleep(10 * time.Hour)
}