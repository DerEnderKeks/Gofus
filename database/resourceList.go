package database

import (
	"github.com/Fisado/Gofus/database/models"
)

var resources = map[string][2]interface{}{
	"users": {models.User{}, },
	"files": models.File{},
	"groups": models.Group{},
	"redirects": models.Redirect{},
	"domains": models.Domain{},
	"settings": models.Setting{},
	"apikeys": models.ApiKey{},
}

func Get(resource string) interface{} {
	return resources[resource]
}
