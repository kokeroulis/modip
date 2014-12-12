package main

import (
	"github.com/kokeroulis/modip/db"
	"github.com/kokeroulis/modip/models"
)

func createAkademicYears() {
	a1 := models.AkademicYear{
		Name: "0000-2013",
	}

	a1.Create()

	a2 := models.AkademicYear{
		Name: "2013-2014",
	}

	a2.Create()
}

func main() {
	Db.Connect()

//	connectDb()

	createAkademicYears()
	exportDepartments()
}
