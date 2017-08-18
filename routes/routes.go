package routes

import (
	"github.com/Fisado/Gofus/routes/api"
	"net/http"
	"github.com/spf13/viper"
	"github.com/Fisado/Gofus/log"
	"github.com/nmaggioni/goat"
	"github.com/Fisado/Gofus/middleware/auth"
	auth2 "github.com/Fisado/Gofus/routes/auth"
)

var (
	SRV *http.Server
)

func Init() {
	r := goat.New()

	r.Use(auth.Auth)

	api.Init(r.Subrouter("/api"))

	r.Get("/oauth/auth", "oauth_auth", auth2.Auth)
	r.Post("/oauth/token", "oauth_token", auth2.Token)

	SRV = &http.Server{}
	SRV.Addr = viper.GetString("http.address")
	SRV.Handler = r
	log.Fatal(SRV.ListenAndServe())
}
