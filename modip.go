package main

import (
	"github.com/codegangsta/negroni"
	"github.com/kokeroulis/modip/api"
)

func main() {
	n := negroni.Classic()

	api.SetupApi(n)
	n.Run(":3001")
}

