package service

import (
	"base-gin/app/repository"
	"base-gin/config"
)

var (
	accountService *AccountService
	personService  *PersonService
)

func SetupServices(cfg *config.Config) {
	accountService = newAccountService(cfg, repository.GetAccountRepo())
	personService = newPersonService(repository.GetPersonRepo())
}

func GetAccountService() *AccountService {
	if accountService == nil {
		panic("account service is not initialised")
	}

	return accountService
}

func GetPersonService() *PersonService {
	return personService
}
