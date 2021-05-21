# Go Caching Example on Web API

Implementation of Go Caching from `"github.com/patrickmn/go-cache"` to reduce
load time for frequently accessed web items. Caching is placed on the handler/service layer.


## Build

`go build src/main.go`

## Entity

CryptoToken = ID, Name, Token Available, Price

## DB

Build your DB schema with following requirements :

1. DB engine = Postgre SQL
2. Schema name = learncaching
3. Table name = cryptocurr
4. Cryptocurr columns : id (bigint PK), crypto_name (varchar), token_availability (bigint), price (bigint)
5. Set these config in `config` folder