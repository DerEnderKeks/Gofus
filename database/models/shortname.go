package models

type ShortName struct {
	UUIDModel

	ShortName string
	File      File
	Redirect  Redirect
}
