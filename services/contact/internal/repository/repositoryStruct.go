package repository

import "ass_4/services/contact/internal/domain"

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{
	}
}


func (r *repositoryImpl) AddContact(c domain.Contact) (domain.Contact, error) {
	return domain.Contact{}, nil
}

func (r *repositoryImpl) DeleteContact(c domain.Contact) (domain.Contact, error) {
	return domain.Contact{}, nil
}

func (r *repositoryImpl) CreateGroup(g domain.Group) (domain.Group, error) {
	return domain.Group{}, nil
}

func (r *repositoryImpl) GetGroup(g domain.Group) (domain.Group, error) {
	return domain.Group{}, nil
}

func (r *repositoryImpl) AddContactToGroup(c domain.Contact, g domain.Group) (domain.Contact, domain.Group, error) {
	return domain.Contact{}, domain.Group{}, nil
}