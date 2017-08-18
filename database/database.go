package database

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/Fisado/Gofus/log"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB *gorm.DB
)

func Init() {
	var err error
	DB, err = gorm.Open("postgres", "host=" + viper.GetString("postgres.host") +
		" user=" + viper.GetString("postgres.user") +
		" dbname=" + viper.GetString("postgres.dbname") +
		" sslmode=" + viper.GetString("postgres.sslmode") +
		" password=" + viper.GetString("postgres.password"))

	// TODO: Auto Migration (http://jinzhu.me/gorm/database.html#migration)

	log.Fatal(err)
}
