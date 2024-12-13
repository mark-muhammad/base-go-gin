package service

import (
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
	"base-gin/config"
	"base-gin/exception"
	"base-gin/util"
)

type AccountService struct {
	cfg  *config.Config
	repo *repository.AccountRepository
}

func (s *AccountService) init(cfg *config.Config) {
	s.cfg = cfg
	s.repo = *repository.GetRepository[*repository.AccountRepository]()
}

func (s *AccountService) Login(p dto.AccountLoginReq) (dto.AccountLoginResp, error) {
	var resp dto.AccountLoginResp

	item, err := s.repo.GetByUsername(p.Username)
	if err != nil {
		return resp, err
	}

	if paswdOk := item.VerifyPassword(p.Password); !paswdOk {
		return resp, exception.ErrUserLoginFailed
	}

	aToken, err := util.CreateAuthAccessToken(*s.cfg, item.Username)
	if err != nil {
		return resp, err
	}

	rToken, err := util.CreateAuthRefreshToken(*s.cfg, item.Username)
	if err != nil {
		return resp, err
	}

	resp.AccessToken = aToken
	resp.RefreshToken = rToken

	return resp, nil
}
