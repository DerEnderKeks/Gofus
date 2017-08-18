package main

import (
	"github.com/getsentry/raven-go"
	"github.com/Fisado/Gofus/routes"
	"github.com/spf13/viper"
	"github.com/Fisado/Gofus/config"
	"github.com/Fisado/Gofus/database"
)


func init() {
	config.Init()
	raven.SetDSN(viper.GetString("sentry.dsn"))
}

func main() {
	database.Init()
	defer database.DB.Close()
	routes.Init()
	defer routes.SRV.Close()
}
