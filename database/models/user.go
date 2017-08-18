package models

type User struct {
	UUIDModel

	UserName      string `gorm:"unique_index"`
	PasswordHash  string
	Email         string
	Files         []File
	Groups        []Group
	Domains       []Domain
	DefaultDomain Domain
}
