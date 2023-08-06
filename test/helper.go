package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func Secret(t *testing.T) (key, secret string) {
	t.Helper()

	bt, err := os.ReadFile("secret.json")
	require.NoError(t, err)
	return gjson.GetBytes(bt, "appkey").String(), gjson.GetBytes(bt, "appsecret").String()
}
