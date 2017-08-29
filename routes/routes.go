package routes

import (
	"github.com/Fisado/Gofus/routes/api"
	"net/http"
	"github.com/spf13/viper"
	"github.com/abiosoft/river"
	"github.com/Fisado/Gofus/log"
	"github.com/Fisado/Gofus/routes/auth"
	"github.com/Fisado/Gofus/handler"
	"github.com/Fisado/Gofus/routes/dashboard"
)

var (
	SRV *http.Server
)

func Init() {
	r := river.New()

	apiEndpoint := river.NewEndpoint()
	api.Init(apiEndpoint, r)
	r.Handle("/api", apiEndpoint)

	authEndpoint := river.NewEndpoint()
	auth.Init(authEndpoint)
	r.Handle("/auth", authEndpoint)

	dashboardEndpoint := river.NewEndpoint()
	dashboard.Init(dashboardEndpoint)
	r.Handle("/dashboard", dashboardEndpoint)

	r.NotFound(handler.NotFound) // This handles short urls like files and redirects too. Why? Because httprouter doesn't like named parameters when they could possibly override other routes.
	r.NotAllowed(handler.MethodNotAllowed)

	SRV = &http.Server{}
	SRV.Addr = viper.GetString("http.address")
	SRV.Handler = r
	log.Fatal(SRV.ListenAndServe())
}
