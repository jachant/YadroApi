package service

import (
	"yadro/domain"
	"yadro/repository"
	"yadro/usecases"
)

type Information struct {
	repo repository.Information
}

func NewInfoService(repo repository.Information) usecases.Information {
	return &Information{repo: repo}
}

func (i *Information) GetInfo() (*domain.Information, error) {
	result, err := i.repo.GetInfo()
	if err != nil {
		return nil, err
	}
	return result, nil
}
