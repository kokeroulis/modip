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

	connectDb()

	//query := "SELECT mail FROM dep_users where name = ?;"
	//var email string
	//err := MySQLDb.QueryRow(query, "hartonas").
	//Scan(&email)

	//Db.CheckQuery(err, query)
	createAkademicYears()
	exportDepartments()
}
