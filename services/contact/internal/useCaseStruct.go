package internal

import (
	"ass_4/services/contact/internal/domain"
	"ass_4/services/contact/internal/repository"
)

type useCaseImpl struct {
	repository repository.Repository
}

func (uc *useCaseImpl) AddContact(c domain.Contact) (domain.Contact, error) {
	return domain.Contact{}, nil
}

func (uc *useCaseImpl) DeleteContact(c domain.Contact) (domain.Contact, error) {
	return domain.Contact{}, nil
}

func (uc *useCaseImpl) CreateGroup(g domain.Group) (domain.Group, error) {
	return domain.Group{}, nil
}

func (uc *useCaseImpl) GetGroup(g domain.Group) (domain.Group, error) {
	return domain.Group{}, nil
}

func (uc *useCaseImpl) AddContactToGroup(c domain.Contact, g domain.Group) (domain.Contact, domain.Group, error) {
	return domain.Contact{}, domain.Group{}, nil
}

func NewUseCase(repository repository.Repository) UseCase {
	return &useCaseImpl{
		repository: repository,
	}
}