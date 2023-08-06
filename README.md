<p align="center">
<h1 align="center">ebest-go</h1>
<p align="center"><a href="https://ebestsec.co.kr/">ebest investment</a> Unofficial client library for Go </p>


## Installation

```bash
go get github.com/gobenpark/ebest-go
```


## Usage

```go
cli := ebest-go.NewClient(
	WithAuth(key,secret),
	WithAutomaticTokenCache(true,time.Hour),
	)

// Do something
cli.AccessToken(context.TODO()) 
```
