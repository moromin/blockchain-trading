package controllers

import (
	"blockchain-trading/usecase"
	"context"
	"fmt"
	"net/http"
)

type OHLCController struct {
	Interactor *usecase.OHLCInteractor
}

func (oc *OHLCController) GetOHLCsBySymbol(c Context) error {
	symbol := c.Param("symbol")
	ohlcs, err := oc.Interactor.OHLCsBySymbol(context.Background(), symbol)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else if len(ohlcs) == 0 {
		return APICustomError(c, http.StatusBadRequest, fmt.Sprintf("%s is not exist", symbol))
	}
	return c.JSON(http.StatusOK, ohlcs)
}
