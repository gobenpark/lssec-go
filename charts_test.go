package ebest_go

import (
	"context"
	"ebest-go/test"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func Test_Charts(t *testing.T) {
	key, secret := test.Secret(t)
	cli := NewClient(
		WithAuth(key, secret), WithAutomaticTokenCache(true))

	t.Run("TotalChartOption", func(t *testing.T) {
		_, err := cli.Charts(context.TODO(), "", TotalChartOption{
			Code:       KOSDAQ_CODE,
			ChartType:  "1",
			TickCount:  0,
			QueryCount: 500,
			Tdgb:       "0",
			StartDate:  "20230404",
			EndDate:    time.Now().Format("20060102"),
		})
		require.NoError(t, err)
	})

}

func TestCreateFile(t *testing.T) {
	f, err := os.Create("samplefile")
	require.NoError(t, err)

	//f.Truncate(0)
	//
	f.Write([]byte("hello world2"))
	//
	f.Close()
}

func TestJWT(t *testing.T) {
	bt, err := os.ReadFile("token")
	require.NoError(t, err)

	tk, err := jwt.Parse(string(bt), nil)
	d, err := tk.Claims.GetExpirationTime()
	require.NoError(t, err)
	fmt.Println(d.After(time.Now()))

}
