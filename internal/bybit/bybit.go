package bybit

import (
	"os"

	bybitApi "github.com/wuhewuhe/bybit.go.api"
)

func NewTestClient() *bybitApi.Client {
	return bybitApi.NewBybitHttpClient(
		os.Getenv("BYBIT_API_KEY_TEST"),
		os.Getenv("BYBIT_API_KEY_SECRET"),
		bybitApi.WithBaseURL(bybitApi.TESTNET),
	)
}
