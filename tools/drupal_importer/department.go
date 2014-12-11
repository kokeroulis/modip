package main

import (
	"github.com/kokeroulis/modip/models"
)

func exportDepartments() {

	query := ""

	rows, err := MySQLDb.Database.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := &models.Department{}

		if err := rows.Scan(&it.Name); err != nil {
			panic(err)
		} else {
			departments = append(departments, it)
			exportTeachers(it)
			exportLessons(it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}
}

func exportTeachers(d *models.Department) {

}

func exportLessons(d *models.Department) {

}

