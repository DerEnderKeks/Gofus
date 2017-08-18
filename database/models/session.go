package models

type session struct {
	UUIDModel

	Token      string
	UserAgent  string
}
