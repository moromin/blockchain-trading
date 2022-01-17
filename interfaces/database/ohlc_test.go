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

func TestRegisterOHLC(t *testing.T) {
	type args struct {
		ctx context.Context
		arg usecase.RegisterOHLCParams
	}

	testArgs := usecase.RegisterOHLCParams{
		Symbol:                   util.RandomString(),
		Interval:                 util.RandomString(),
		Opentime:                 util.RandomDate(),
		Open:                     util.RandomFloat(),
		High:                     util.RandomFloat(),
		Low:                      util.RandomFloat(),
		Close:                    util.RandomFloat(),
		Volume:                   util.RandomFloat(),
		Closetime:                util.RandomDate(),
		QuoteAssetVolume:         util.RandomFloat(),
		NumberOfTrades:           util.RandomInt(),
		TakerBuyBaseAssetVolume:  util.RandomFloat(),
		TakerBuyQuoteAssetVolume: util.RandomFloat(),
	}

	tests := []struct {
		name        string
		mockClosure func(sqlmock.Sqlmock)
		args        args
		want        entity.OHLC
		assertion   assert.ErrorAssertionFunc
	}{
		{
			name: "Success",
			mockClosure: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "symbol", "interval", "opentime", "open", "high", "low", "close", "volume", "closetime", "quote_asset_volume", "number_of_trades", "taker_buy_base_asset_volume", "taker_buy_quote_asset_volume"}).
					AddRow(1, testArgs.Symbol, testArgs.Interval, testArgs.Opentime, testArgs.Open, testArgs.High, testArgs.Low, testArgs.Close, testArgs.Volume, testArgs.Closetime, testArgs.QuoteAssetVolume, testArgs.NumberOfTrades, testArgs.TakerBuyBaseAssetVolume, testArgs.TakerBuyQuoteAssetVolume)
				mock.ExpectQuery(regexp.QuoteMeta(database.RegisterOHLC)).
					WillReturnRows(rows)
			},
			args: args{
				ctx: context.Background(),
				arg: testArgs,
			},
			want: entity.OHLC{
				ID:                       1,
				Symbol:                   testArgs.Symbol,
				Interval:                 testArgs.Interval,
				OpenTime:                 testArgs.Opentime,
				Open:                     testArgs.Open,
				High:                     testArgs.High,
				Low:                      testArgs.Low,
				Close:                    testArgs.Close,
				Volume:                   testArgs.Volume,
				CloseTime:                testArgs.Closetime,
				QuoteAssetVolume:         testArgs.QuoteAssetVolume,
				NumberOfTrades:           testArgs.NumberOfTrades,
				TakerBuyBaseAssetVolume:  testArgs.TakerBuyBaseAssetVolume,
				TakerBuyQuoteAssetVolume: testArgs.TakerBuyQuoteAssetVolume,
			},
			assertion: assert.NoError,
		},
		{
			name: "Failure on insert",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(database.RegisterOHLC)).
					WillReturnError(fmt.Errorf("Register OHLC data error"))
			},
			args: args{
				ctx: context.Background(),
				arg: usecase.RegisterOHLCParams{},
			},
			want:      entity.OHLC{},
			assertion: assert.Error,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				log.Fatal("failed to init db mock:", err)
			}

			tt.mockClosure(mock)

			or := &database.OHLCRepository{
				Db: db,
			}
			got, err := or.RegisterOHLC(tt.args.ctx, tt.args.arg)
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
