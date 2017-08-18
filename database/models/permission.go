package models

type Permission struct {
	IDModel

	PermissionName string
	Value          string
	Group          Group
}
