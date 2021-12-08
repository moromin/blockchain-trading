package database_test

import (
	"blockchain-trading/infrastructure"
	"blockchain-trading/interfaces/database"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "user=root password=secret host=localhost dbname=ohlc sslmode=disable"
)

var repo *database.DatabaseRepository

// func TestStoreCurrency(t *testing.T) {
// 	currencies := []entity.Currency{
// 		{
// 			Coin: "Hoge",
// 			Name: "hogehoge",
// 		},
// 		{
// 			Coin: "Fuga",
// 			Name: "fugafuga",
// 		},
// 	}

// 	err := repo.StoreCurrency(currencies)
// 	require.NoError(t, err)
// }

// func TestFindAllCurrency(t *testing.T) {
// 	currencies, err := repo.FindAllCurrency()
// 	require.NoError(t, err)
// 	require.NotEmpty(t, currencies)
// }

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	handler := infrastructure.SqlHandler{Conn: conn}
	repo = &database.DatabaseRepository{SqlHandler: &handler}

	os.Exit(m.Run())
}
