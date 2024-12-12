package service

import (
	"base-gin/config"
)

type Service interface {
	init(*config.Config)
}

var services []Service = []Service{
	&AccountService{},
	&PersonService{},
}

func SetupServices(cfg *config.Config) {
	for _, v := range services {
		v.init(cfg)
	}
}

func GetService[T Service]() *T {
	for _, v := range services {
		if r, ok := v.(T); ok {
			return &r
		}
	}

	return nil
}
