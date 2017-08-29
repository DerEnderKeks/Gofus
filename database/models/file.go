package models

type File struct {
	UUIDModel

	FileName          string
	ShortName         ShortName
	FileSize          uint
	MagicMimeType     string
	ExtensionMimeType string
	User              User
}
