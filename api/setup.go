package api

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/kokeroulis/modip/config"
	_ "github.com/lib/pq"
	"gopkg.in/unrolled/render.v1"
)

var Render *render.Render

func SetupApi(n *negroni.Negroni) {
	setupRedis(n)

	Render = render.New(render.Options{})

	setupRoutes(n)
}

func setupRedis(n *negroni.Negroni) {
	store, err := sessions.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))

	if err != nil {
		fmt.Println("Can't connect to redis")
		panic(err)
	}

	n.Use(sessions.Sessions("my_session", store))
}

