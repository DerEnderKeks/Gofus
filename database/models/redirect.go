package models

type Redirect struct {
	UUIDModel

	Target    string
	ShortName ShortName
	User      User
}
