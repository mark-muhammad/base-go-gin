package repository

import (
	"base-gin/storage"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Repository interface {
	init(db *gorm.DB)
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

	var t T
	errMsg := fmt.Sprintf("%s is not initialised", reflect.TypeOf(t).String())
	panic(errMsg)
}
