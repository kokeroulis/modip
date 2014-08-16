package main

import (
	"github.com/codegangsta/negroni"
	"api"
)

func main() {
	n := negroni.Classic()

	api.SetupApi(n)
	n.Run(":3000")
}

