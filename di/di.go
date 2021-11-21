package di

import (
	"blockchain-trading/infrastructure"
	"blockchain-trading/interfaces/exchange"
	"blockchain-trading/interfaces/presenter"
	"blockchain-trading/usecase"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func New(target infrastructure.Target) (*dig.Container, error) {
	c := dig.New()

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

	if err := c.Provide(func(ac exchange.APIClient) *exchange.ExchangeRepository {
		return &exchange.ExchangeRepository{APIClient: ac}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := c.Provide(func(t infrastructure.Target) exchange.APIClient {
		return &infrastructure.APIClient{
			HTTPClient: &http.Client{},
			Target:     t,
		}
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	c.Provide(func() infrastructure.Target {
		return target
	})

	return c, nil
}
