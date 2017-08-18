package config

import (
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
	"github.com/Fisado/Gofus/log"
)

func Init() {
	pflag.String("config", "", "Path to the config file")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	if viper.GetString("config") != "" {
		viper.SetConfigFile(viper.GetString("config"))
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./data/config")
	viper.AddConfigPath("/etc/gofus")
	viper.AddConfigPath("/srv/gofus/config")

	setDefaults()

	err := viper.ReadInConfig()
	if err != nil {
		log.Warning(err)
	}

	viper.AutomaticEnv()
}

func setDefaults() {
	viper.SetDefault("sentry.dsn", "%%SENTRY_DSN%%")
	viper.SetDefault("postgres.host", "db")
	viper.SetDefault("postgres.user", "gofus")
	viper.SetDefault("postgres.password", "gofus")
	viper.SetDefault("postgres.dbname", "gofus")
	viper.SetDefault("postgres.sslmode", "disable")
	viper.SetDefault("http.address", ":3000")
	viper.SetDefault("loglevel", log.LevelInfo.ID)
}
