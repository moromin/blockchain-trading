package di

import (
	"blockchain-trading/infrastructure"
	"blockchain-trading/interfaces/api"
	"blockchain-trading/interfaces/controllers"
	"blockchain-trading/interfaces/database"
	"blockchain-trading/interfaces/exchange"
	"blockchain-trading/interfaces/presenter"
	"blockchain-trading/usecase"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func NewAPIClient(target infrastructure.Target) (*dig.Container, error) {
	c := dig.New()

	// exchange
	if err := c.Provide(func(ei *usecase.ExchangeInteractor) *presenter.ExchangePresenter {
		return &presenter.ExchangePresenter{Interactor: ei}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(exRepo *exchange.ExchangeRepository) *usecase.ExchangeInteractor {
		return &usecase.ExchangeInteractor{ExchangeRepository: exRepo}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(ac api.APIClient) *exchange.ExchangeRepository {
		return &exchange.ExchangeRepository{APIClient: ac}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	// apiclient
	if err := c.Provide(func(t infrastructure.Target) api.APIClient {
		return &infrastructure.APIClient{
			HTTPClient: &http.Client{},
			Target:     t,
		}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func() infrastructure.Target {
		return target
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}

func NewCurrency(db *sql.DB) (*dig.Container, error) {
	c := dig.New()

	if err := c.Provide(func(ci *usecase.CurrencyInteractor) *presenter.CurrencyPresenter {
		return &presenter.CurrencyPresenter{Interactor: ci}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(cr *database.CurrencyRepository) *usecase.CurrencyInteractor {
		return &usecase.CurrencyInteractor{Repo: cr}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(db database.DBTX) *database.CurrencyRepository {
		return &database.CurrencyRepository{Db: db}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func() database.DBTX {
		return db
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}

func NewOHLC(db *sql.DB) (*dig.Container, error) {
	c := dig.New()

	if err := c.Provide(func(oi *usecase.OHLCInteractor) *controllers.OHLCController {
		return &controllers.OHLCController{Interactor: oi}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(oi *usecase.OHLCInteractor) *presenter.OHLCPresenter {
		return &presenter.OHLCPresenter{Interactor: oi}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(or *database.OHLCRepository) *usecase.OHLCInteractor {
		return &usecase.OHLCInteractor{Repo: or}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(db database.DBTX) *database.OHLCRepository {
		return &database.OHLCRepository{Db: db}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func() database.DBTX {
		return db
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}
