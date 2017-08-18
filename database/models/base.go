package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type UUIDModel struct {
	gorm.Model

	ID uuid.UUID `gorm:"primary_key"`
}

type IDModel struct {
	gorm.Model
}


