package presenter

import (
	"blockchain-trading/entity"
	"blockchain-trading/usecase"
	"context"
)

type OHLCPresenter struct {
	Interactor *usecase.OHLCInteractor
}

func (op *OHLCPresenter) RegisterOHLCs(ohlcs []entity.OHLC) error {
	for _, ohlc := range ohlcs {
		arg := usecase.RegisterOHLCParams{
			Symbol:                   ohlc.Symbol,
			Interval:                 ohlc.Interval,
			Opentime:                 ohlc.OpenTime,
			Open:                     ohlc.Open,
			High:                     ohlc.High,
			Low:                      ohlc.Low,
			Close:                    ohlc.Close,
			Volume:                   ohlc.Volume,
			Closetime:                ohlc.CloseTime,
			QuoteAssetVolume:         ohlc.QuoteAssetVolume,
			NumberOfTrades:           ohlc.NumberOfTrades,
			TakerBuyBaseAssetVolume:  ohlc.TakerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: ohlc.TakerBuyQuoteAssetVolume,
		}
		_, err := op.Interactor.AddOHLC(context.Background(), arg)
		if err != nil {
			return err
		}
	}
	return nil
}
