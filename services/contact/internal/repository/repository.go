package repository

import (
	"ass_4/services/contact/internal/domain"
)

type Repository interface {
	AddContact(c domain.Contact) (domain.Contact, error)
	DeleteContact(c domain.Contact) (domain.Contact, error)
	CreateGroup(g domain.Group) (domain.Group, error)
	GetGroup(g domain.Group) (domain.Group, error)
	AddContactToGroup(c domain.Contact, g domain.Group) (domain.Contact, domain.Group, error)
}