package ebest_go

import (
	"testing"

	"github.com/gobenpark/ebest-go/test"
)

func ClientHelper(t *testing.T) *Client {
	t.Helper()
	key, secret := test.Secret(t)
	cli := NewClient(
		WithAuth(key, secret), WithAutomaticTokenCache(true), WithDebug(true))
	return cli
}

//
//func Test_ClientRealtime(t *testing.T) {
//	cli := ClientHelper(t)
//	ko, err := cli.Kospi(context
//
//	contentskq := lo.Map(kq, func(item Code, index int) SubscriptionContent {
//		return SubscriptionContent{
//			Type:   AddPriceTRType,
//			TRCD:   KOSDAQContract,
//			Ticker: item.Code,
//		}
//	})
//	cch, err := cli.Subscribe(context.TODO(), contentskq...)
//	require.NoError(t, err)
//
//	go func() {
//		for i := range cch {
//			fmt.Println("kd: " + string(i))
//		}
//	}()
//	time.Sleep(10 * time.Hour)
//}
