package api

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	_ "github.com/lib/pq"
	"gopkg.in/unrolled/render.v1"
)

var Db *sql.DB
var Render *render.Render

func SetupApi(n *negroni.Negroni) {
	setupRedis(n)
	setupDb(n)

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

func setupDb(n *negroni.Negroni) {
	d, err := sql.Open("postgres", "postgres://tsiapaliokas:@localhost/modip_db?sslmode=disable")

	if err != nil {
		fmt.Println("Can't connect to postgresql")
		panic(err)
	}

	pingErr := d.Ping()
	if pingErr != nil {
		fmt.Println("Can't connect to postgresql")
		panic(pingErr)
	}

	Db = d
}
