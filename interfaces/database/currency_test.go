package database_test

import (
	"blockchain-trading/entity"
	"blockchain-trading/interfaces/database"
	"blockchain-trading/usecase"
	"blockchain-trading/util"
	"context"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// func TestCurrencyRepository_ListCurrencies(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		arg usecase.ListCurrenciesParams
// 	}
// 	tests := []struct {
// 		name      string
// 		args      args
// 		want      []entity.Currency
// 		assertion assert.ErrorAssertionFunc
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			cr := &database.CurrencyRepository{
// 				Db: tt.fields.Db,
// 			}
// 			got, err := cr.ListCurrencies(tt.args.ctx, tt.args.arg)
// 			tt.assertion(t, err)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }

func TestRegisterCurrency(t *testing.T) {
	type args struct {
		ctx context.Context
		arg usecase.ResisterCurrencyParams
	}

	testArgs := usecase.ResisterCurrencyParams{
		Coin: util.RandomCoin(),
		Name: util.RandomName(),
	}

	tests := []struct {
		name        string
		mockClosure func(sqlmock.Sqlmock)
		args        args
		want        entity.Currency
		assertion   assert.ErrorAssertionFunc
	}{
		{
			name: "Success",
			mockClosure: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "coin", "name"}).
					AddRow(1, testArgs.Coin, testArgs.Name)
				mock.ExpectQuery(
					regexp.QuoteMeta(database.RegisterCurrency),
				).WillReturnRows(rows)
			},
			args: args{
				ctx: context.Background(),
				arg: testArgs,
			},
			want: entity.Currency{
				ID:   1,
				Coin: testArgs.Coin,
				Name: testArgs.Name,
			},
			assertion: assert.NoError,
		},
		{
			name: "Failure on insert",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(database.RegisterCurrency),
				).WillReturnError(fmt.Errorf("some error"))
			},
			args: args{
				ctx: context.Background(),
				arg: usecase.ResisterCurrencyParams{},
			},
			want:      entity.Currency{},
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.New()
			if err != nil {
				log.Fatal("failed to init db mock:", err)
			}

			tt.mockClosure(mock)

			cr := &database.CurrencyRepository{
				Db: db,
			}
			got, err := cr.RegisterCurrency(tt.args.ctx, tt.args.arg)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)

			mock.ExpectClose()
			if err = db.Close(); err != nil {
				t.Errorf("failed to close db: %s", err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
