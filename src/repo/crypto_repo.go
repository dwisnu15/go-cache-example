package repo

import (
	Conn "go_cache_example/db"
	"go_cache_example/src/model"
	"log"
)

type CryptoRepo struct {
}

//func GetCryptoRepo() *CryptoRepo {
//	return &CryptoRepo{}
//}

func (cr *CryptoRepo) GetByID(id int32) (*model.CryptoToken, error) {
	db, err := Conn.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	var data model.CryptoToken
	err = db.QueryRow("SELECT id, crypto_name, token_available, price FROM learncaching.cryptocurr WHERE id = $1", id).Scan(
		&data.ID, &data.Name, &data.TokenAvailable, &data.Price)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &data, nil

}

func (cr *CryptoRepo) GetAll() (*[]model.CryptoToken, error) {
	db, err := Conn.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	var results []model.CryptoToken
	rows, err := db.Query("SELECT id, crypto_name, token_available, price FROM learncaching.cryptocurr LIMIT 10")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.CryptoToken
		err := rows.Scan(&data.ID, &data.Name, &data.TokenAvailable, &data.Price)
		if err != nil {
			log.Println(err)
		}
		results = append(results, data)
	}

	return &results, nil

}
