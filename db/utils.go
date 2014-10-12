package Db

import (
	"database/sql"
	"fmt"
	"reflect"
)

func CheckQuery(err error, query string) {
	if err == sql.ErrNoRows {
		fmt.Println(query)
		panic("No rows in query! Something is wrong with the db")
	} else if err != nil {
		// print the error
		// TODO use logger
		fmt.Println(query)
		fmt.Println(err)
		panic("Db Error")
	}
}

func SqlRowsToArray(query string, args []interface{}, iterator interface{}) []interface{} {
	var results []interface{}

	rows, err := Database.Query(query, args...)

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
