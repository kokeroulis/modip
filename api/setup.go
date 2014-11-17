package api

import (
	"github.com/codegangsta/negroni"
	"github.com/kokeroulis/modip/models"
)

func SetupApi(n *negroni.Negroni) {
	models.SetupRedis()

	setupRoutes(n)
}

