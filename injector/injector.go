package injector

import (
	"blockchain-trading/domain/repository"
	"blockchain-trading/infra"
	"blockchain-trading/usecase"
)

func InjectAPIClient() infra.APIClient {
	apiClient := infra.NewAPIClient()
	return apiClient
}

func InjectExchangeRepository() repository.ExchangeRepository {
	apiClient := InjectAPIClient()
	return infra.NewExchangeRepository(apiClient)
}

func InjectExchangeUsecase() usecase.ExchangeUsecase {
	exRepo := InjectExchangeRepository()
	return usecase.NewExchangeUsecase(exRepo)
}
