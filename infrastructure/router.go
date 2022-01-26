package infrastructure

import (
	"blockchain-trading/interfaces/controllers"

	"github.com/labstack/echo"
	"go.uber.org/dig"
)

var (
	OHLCContainer *dig.Container
)

func getOHLCsBySymbol(c echo.Context) error {
	var resErr error
	if err := OHLCContainer.Invoke(func(oc *controllers.OHLCController) error {
		resErr = oc.GetOHLCsBySymbol(c)
		if resErr != nil {
			return resErr
		}
		return nil
	}); err != nil {
		return err
	}
	if resErr != nil {
		return resErr
	}
	return nil
}

func OHLCRouting(e *echo.Echo, container *dig.Container) {
	OHLCContainer = container
	e.GET("ohlc/:symbol", getOHLCsBySymbol)
}
