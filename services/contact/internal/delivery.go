package internal

import (
	"ass_4/services/contact/internal/repository"
)

type Delivery struct {
	useCase    UseCase
	repository repository.Repository
}

func NewDelivery(useCase UseCase, repository repository.Repository) *Delivery {
	return &Delivery{
		useCase:    useCase,
		repository: repository,
	}
}
