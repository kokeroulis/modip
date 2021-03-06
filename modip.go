package main

import (
	"github.com/codegangsta/negroni"
	"github.com/kokeroulis/modip/api"
	"github.com/kokeroulis/modip/db"
	"github.com/kokeroulis/modip/models/categories"
	"flag"
)

func main() {

	var setupDb = flag.Bool("setup-db", false, "Setup db")
	var start = flag.Bool("start", false, "start")

	flag.Parse()

	Db.Connect()

	if *setupDb == true {
		Db.CreateDb()
		categories.CreateCategories()
	}

	if *start {
		n := negroni.Classic()

		api.SetupApi(n)
		n.Run(":3001")
	}
}
