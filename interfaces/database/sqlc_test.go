package database_test

import (
	"blockchain-trading/entity"
	"blockchain-trading/interfaces/database"
	"blockchain-trading/usecase"
	"blockchain-trading/util"
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

const (
	dbDriver = "postgres"
	dbSource = "user=root password=secret host=localhost dbname=ohlc sslmode=disable"
)

var testQueries *database.Queries
var testDB *sql.DB

func createRandomCurrency(t *testing.T) entity.Currency {
	arg := usecase.ResisterCurrencyParams{
		Coin: util.RandomCoin(),
		Name: util.RandomName(),
	}

	currency, err := testQueries.ResisterCurrency(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, currency)

	require.Equal(t, arg.Coin, currency.Coin)
	require.Equal(t, arg.Name, currency.Name)

	require.NotZero(t, currency.ID)

	return currency
}

func TestResisterCurrency(t *testing.T) {
	createRandomCurrency(t)
}

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = &database.Queries{Db: testDB}

	os.Exit(m.Run())
}
