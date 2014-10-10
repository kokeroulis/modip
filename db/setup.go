package Db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/kokeroulis/modip/db/api"
	"github.com/kokeroulis/modip/db/schema"
)

var Database *sql.DB

func Connect() {
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

	Database = d
}

func runQuery(list []string) {
	for _, it := range list {
		_, err := Database.Exec(it)
		if err != nil {
			fmt.Println("Error at: ", it)
			panic(err)
		}
	}
}

func CreateDb() {
	schemaList := []string{
		DbSchema.Department,
		DbSchema.Teacher,
		DbSchema.Paper,
		DbSchema.Book,
	}

	apiList := []string{
		DbApi.BookAdd,
		DbApi.PaperAdd,
		DbApi.TeacherLogin,
	}

	runQuery(schemaList)
	runQuery(apiList)
}
