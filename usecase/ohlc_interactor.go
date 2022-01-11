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
