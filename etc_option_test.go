package ebest_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ETC(t *testing.T) {
	cli := ClientHelper(t)
	t.Run("NewTickerOption", func(t *testing.T) {
		_, err := cli.Execute(context.TODO(), NewTickerOption{
			Market:     "0",
			StartMonth: "202301",
			EndMonth:   "202308",
			Index:      0,
		})
		require.NoError(t, err)
	})

}
