package repository

import (
	"base-gin/storage"

	"gorm.io/gorm"
)

type Repository interface {
	init(db *gorm.DB)
}

var repos []Repository = []Repository{
	&AccountRepository{},
	&PersonRepository{},
}

func SetupRepositories() {
	db := storage.GetDB()

	for i := 0; i < len(repos); i++ {
		repos[i].init(db)
	}
}

func GetRepository[T Repository]() *T {
	for _, v := range repos {
		if r, ok := v.(T); ok {
			return &r
		}
	}

	return nil
}
