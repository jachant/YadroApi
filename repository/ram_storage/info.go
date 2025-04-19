package ram_storage

import (
	"yadro/domain"
	"yadro/repository"
)

type Information struct {
}

func NewInfoRepository() repository.Information {
	return &Information{}
}

func (i *Information) GetInfo() (*domain.Information, error) {
	result := &domain.Information{
		Version: "0.1.0",
		Service: "weather",
		Author:  "a.antonov",
	}
	return result, nil
}
