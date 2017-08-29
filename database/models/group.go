package models

type Group struct {
	UUIDModel

	GroupName       string
	InheritedGroups []Group
	Domains         []Domain
	Permissions     int64
}
