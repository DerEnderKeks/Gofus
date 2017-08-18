package models

type File struct {
	UUIDModel

	FileName          string
	ShortName         string
	FileSize          uint
	MagicMimeType     string
	ExtensionMimeType string
	User              User
}
