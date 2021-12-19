package database_test

import (
	"blockchain-trading/entity"
	"blockchain-trading/usecase"
	"blockchain-trading/util"
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createRandomCurrency(t *testing.T) entity.Currency {
	arg := usecase.ResisterCurrencyParams{
		Coin: util.RandomCoin(),
		Name: util.RandomName(),
	}

	currency, err := testRepo.RegisterCurrency(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, currency)

	require.Equal(t, arg.Coin, currency.Coin)
	require.Equal(t, arg.Name, currency.Name)

	require.NotZero(t, currency.ID)

	return currency
}

func TestResisterCurrency(t *testing.T) {
	t.Parallel()

	createRandomCurrency(t)
}

func TestListCurrencies(t *testing.T) {
	t.Parallel()

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
