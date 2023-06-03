package internal

interface UseCase {
	addContact(Contact c) (Contact, error)
	getContact(Contact c) (Contact, error)
	deleteContact(Contact c) (Contact, error)
	createGroup(Group g) (Group, error)
	getGroup(Group g) (Group, error)
	addContactToGroup(Contact c, Group g) (Contact, Group, error)

}