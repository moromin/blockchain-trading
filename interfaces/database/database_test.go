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

var testRepo *database.DatabaseRepository
var testDB *sql.DB

func createRandomCurrency(t *testing.T) entity.Currency {
	arg := usecase.ResisterCurrencyParams{
		Coin: util.RandomCoin(),
		Name: util.RandomName(),
	}

	currency, err := testRepo.ResisterCurrency(context.Background(), arg)
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

func TestGetCurrency(t *testing.T) {
	currency1 := createRandomCurrency(t)
	currency2, err := testRepo.GetCurrency(context.Background(), currency1.Coin)

	require.NoError(t, err)
	require.NotEmpty(t, currency2)

	require.Equal(t, currency1.ID, currency2.ID)
	require.Equal(t, currency1.Coin, currency2.Coin)
	require.Equal(t, currency1.Name, currency2.Name)
}

func TestListCurrencies(t *testing.T) {
	arg := usecase.ListCurrenciesParams{
		Limit:  10,
		Offset: 5,
	}

	currencies, err := testRepo.ListCurrencies(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, currencies, 10)

	for _, currency := range currencies {
		require.NotEmpty(t, currency)
	}
}

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testRepo = &database.DatabaseRepository{Db: testDB}

	os.Exit(m.Run())
}
