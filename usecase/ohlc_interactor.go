package usecase

import (
	"blockchain-trading/entity"
	"context"
)

type OHLCInteractor struct {
	Repo OHLCRepository
}

func (oi *OHLCInteractor) AddOHLC(ctx context.Context, arg RegisterOHLCParams) (entity.OHLC, error) {
	ohlc, err := oi.Repo.RegisterOHLC(ctx, arg)
	return ohlc, err
}

func (oi *OHLCInteractor) OHLCsBySymbol(ctx context.Context, symbol string) ([]entity.OHLC, error) {
	ohlcs, err := oi.Repo.FindBySymbol(ctx, symbol)
	return ohlcs, err
}
