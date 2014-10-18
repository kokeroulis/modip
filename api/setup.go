package api

import (
	"github.com/codegangsta/negroni"
	"github.com/kokeroulis/modip/models"
	"gopkg.in/unrolled/render.v1"
)

var Render *render.Render

func SetupApi(n *negroni.Negroni) {
	models.SetupRedis()

	Render = render.New(render.Options{})

	setupRoutes(n)
}

