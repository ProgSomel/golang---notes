package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionStirng() string {
	//? user -> postgres
	//? password -> 8808
	//? host -> localhost
	//? port -> 5432
	//? dbname -> ecommerce

	return "user=postgres password=8808 host=localhost port=5432 dbname=ecommerce"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionStirng()
	dbClient, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbClient, nil
}