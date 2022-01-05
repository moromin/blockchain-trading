package database_test

import (
	"blockchain-trading/interfaces/database"
	"blockchain-trading/usecase"
	"blockchain-trading/util"
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// func createRandomCurrency(t *testing.T) entity.Currency {
// 	arg := usecase.ResisterCurrencyParams{
// 		Coin: util.RandomCoin(),
// 		Name: util.RandomName(),
// 	}

// 	currency, err := testRepo.RegisterCurrency(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, currency)

// 	require.Equal(t, arg.Coin, currency.Coin)
// 	require.Equal(t, arg.Name, currency.Name)

// 	require.NotZero(t, currency.ID)

// 	return currency
// }

// func TestResisterCurrency(t *testing.T) {
// 	t.Parallel()

// 	createRandomCurrency(t)
// }

// func TestListCurrencies(t *testing.T) {
// 	t.Parallel()

// 	arg := usecase.ListCurrenciesParams{
// 		Limit:  10,
// 		Offset: 5,
// 	}

// 	currencies, err := testRepo.ListCurrencies(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.Len(t, currencies, 10)

// 	for _, currency := range currencies {
// 		require.NotEmpty(t, currency)
// 	}
// }

func TestResisterCurrency(t *testing.T) {
	arg := usecase.ResisterCurrencyParams{
		Coin: util.RandomCoin(),
		Name: util.RandomName(),
	}

	newId := 1
	rows := sqlmock.NewRows([]string{"id", "coin", "name"}).AddRow(newId, arg.Coin, arg.Name)
	testRepo.mock.ExpectQuery(
		regexp.QuoteMeta(database.ResisterCurrency),
	).WillReturnRows(rows)

	res, err := testRepo.repo.RegisterCurrency(context.Background(), arg)

	assert.NoError(t, err)
	assert.Equal(t, newId, res.ID)
	assert.Equal(t, arg.Coin, res.Coin)
	assert.Equal(t, arg.Name, res.Name)

	if err := testRepo.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
