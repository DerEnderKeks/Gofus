package models

type Group struct {
	UUIDModel

	GroupName       string
	Permissions     []Permission
	InheritedGroups []Group
	Domains         []Domain
}
