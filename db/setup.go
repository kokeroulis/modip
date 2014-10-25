package Db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/kokeroulis/modip/db/api"
	"github.com/kokeroulis/modip/db/schema"
	"github.com/kokeroulis/modip/config"
)

var Database *sql.DB

func Connect() {
	c := config.NewConfig()
	d, err := sql.Open("postgres", c.PgConnectionString())

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
		DbSchema.AssetType,
		DbSchema.Asset,
		DbSchema.Category,
		DbSchema.CategoryGroup,
	}

	apiList := []string{
		DbApi.TeacherLogin,
		DbApi.AssetAdd,
		DbApi.AssetRemove,
		DbApi.AssetMove,
		DbApi.AssetModify,
		DbApi.CategoryAdd,
		DbApi.CategoryGroupAdd,
		DbApi.CategorySave,
	}

	runQuery(schemaList)
	runQuery(apiList)
}
