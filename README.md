<p align="center">
<h1 align="center">ebest-go</h1>
<p align="center"><a href="https://ebestsec.co.kr/">ebest investment</a> Unofficial client library for Go </p>


## Installation

```bash
go get github.com/gobenpark/ebest-go
```


## Usage

```go
cli := ebest_go.NewClient(
	WithAuth(key,secret),
	WithAutomaticTokenCache(true),
	)

// Do something
cli.AccessToken(context.TODO())

bt, err := cli.Execute(context.TODO(), ebest_go.AccountBalanceOption{
PriceType:       "",
ContractType:    "",
SinglePriceType: "",
Charge:          "",
Ticker:          "",
})
if err != nil {
	return
}
```

## TODO

- 연속조회