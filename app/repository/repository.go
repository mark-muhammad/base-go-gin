package repository

import "base-gin/storage"

var (
	accountRepo *AccountRepository
	personRepo  *PersonRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = newAccountRepository(db)
	personRepo = newPersonRepository(db)
}

func GetAccountRepo() *AccountRepository {
	return accountRepo
}

func GetPersonRepo() *PersonRepository {
	return personRepo
}
