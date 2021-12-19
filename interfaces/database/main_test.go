package database_test

import (
	"blockchain-trading/interfaces/database"
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "user=root password=secret host=localhost dbname=ohlc sslmode=disable"
)

var testRepo *database.CurrencyRepository
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer testDB.Close()

	testRepo = &database.CurrencyRepository{Db: testDB}

	os.Exit(m.Run())
}
