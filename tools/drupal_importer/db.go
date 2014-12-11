package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func connectDb() {
	d, err := sql.Open("mysql", "mod-teliar-1:@#gr2012adipLpidom@/modip-v1")

	if err != nil {
		fmt.Println("Can't connect to postgresql")
		panic(err)
	}

	pingErr := d.Ping()
	if pingErr != nil {
		fmt.Println("Can't connect to postgresql")
		panic(pingErr)
	}

	MySQLDb = d
}

