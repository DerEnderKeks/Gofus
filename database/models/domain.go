package models

type Domain struct {
	UUIDModel

	DomainName string
	Groups     []Group
	Users      []User
}
