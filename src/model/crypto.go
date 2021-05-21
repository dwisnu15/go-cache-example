package model

type CryptoToken struct {
	ID             int32  `json:"Crypto_ID"`
	Name           string `json:"CryptoName"`
	TokenAvailable int32  `json:"TokenAvailable"`
	Price          int32  `json:"price"`
}

type CryptoQuery struct {
	ID int32 `json:"id"`
}
