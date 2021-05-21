package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

var DB *sql.DB

func InitViperConfig() {
	//get config file path
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//set path to look for config file (in: GinAPI)
	viper.AddConfigPath(path + "/config")
	//set the file name of config file
	viper.SetConfigName("Config")

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", viper.GetString("database.name"))
	fmt.Println("Port is\t\t", viper.GetString("server.port"))
}

func InitDBConnection() {
	InitViperConfig()
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("server.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
	)
	engine := viper.GetString("database.engine")

	client, err := sql.Open(engine, psqlconn)
	if err != nil {
		log.Fatalf("Error database configurations, %v", err)
	}
	err = client.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database (ping): ", err)
	}
	DB = client
}

func GetDBConnection() (*sql.DB, error) {
	err := DB.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return DB, nil

}

func CloseDBConnection() error {
	return DB.Close()
}
