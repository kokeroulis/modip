package models

import (
	"database/sql"
	"fmt"
	"reflect"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
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

func checkQuery(err error, query string) {
	if err == sql.ErrNoRows {
		panic("No rows in query! Something is wrong with the db")
	} else if err != nil {
		// print the error
		// TODO use logger
		fmt.Println(err)
		panic("Db Error")
	}
}

func sqlRowsToArray(query string, args []interface{}, iterator interface{}) []interface{} {
	var results []interface{}

	rows, err := Db.Query(query, args...)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {

		v := reflect.ValueOf(iterator)
		var fieldArray []interface{}

		elem := v.Elem()
		for i:=0; i < elem.NumField(); i++ {
			fieldArray = append(fieldArray, elem.Field(i).Addr().Interface())
		}

		if err := rows.Scan(fieldArray...); err != nil {
			fmt.Println(err)
		} else {
			results = append(results, iterator)
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	return results
}
