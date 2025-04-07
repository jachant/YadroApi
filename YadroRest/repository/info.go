package repository

import "yadro/domain"

type Information interface {
	GetInfo() (*domain.Information, error)
}
