package models

type ApiKey struct {
	UUIDModel

	Token string
	Scope string
}
