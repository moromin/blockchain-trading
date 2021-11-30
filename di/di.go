package di

import (
	"blockchain-trading/infrastructure"
	"blockchain-trading/interfaces/api"
	"blockchain-trading/interfaces/database"
	"blockchain-trading/interfaces/exchange"
	"blockchain-trading/interfaces/presenter"
	"blockchain-trading/usecase"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func New(target infrastructure.Target) (*dig.Container, error) {
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

func NewDB(handler infrastructure.SqlHandler) (*dig.Container, error) {
	c := dig.New()

	// sql handler
	if err := c.Provide(func(dbi *usecase.DatabaseInteractor) *presenter.DatabasePresenter {
		return &presenter.DatabasePresenter{Interactor: dbi}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(dbRepo *database.DatabaseRepository) *usecase.DatabaseInteractor {
		return &usecase.DatabaseInteractor{DatabaseRepository: dbRepo}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(sh database.SqlHandler) *database.DatabaseRepository {
		return &database.DatabaseRepository{SqlHandler: sh}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func() database.SqlHandler {
		return &handler
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}
