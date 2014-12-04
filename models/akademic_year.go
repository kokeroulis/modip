package models

import (
	"github.com/kokeroulis/modip/db"
)

type AkademicYear struct {
	Id            int    `schema:"id"`
	Name          string `schema:"name"`
}

func (a *AkademicYear) Create() {
	var alreadyExists bool
	query := `SELECT alreadyExists
			  FROM akademic_year_create($1)`

	err := Db.Database.QueryRow(query, a.Name).Scan(&alreadyExists)

	Db.CheckQuery(err, query)

	if alreadyExists {
		panic("The akademic year already exists!!")
	}
}

func (a *AkademicYear) Load() {
	if a.Id == 0 {
		panic("You can't use this func!!")
	}

	query := `SELECT name
			  FROM akademic_year
			  WHERE id = $1`
	err := Db.Database.QueryRow(query, a.Id).
		Scan(&a.Name)

	Db.CheckQuery(err, query)
}

func (a *AkademicYear) Update() {
	if a.Id == 0 {
		panic("You can't use this func!!")
	}

	query := `UPDATE akademic_year SET
			  name = $1
			  WHERE id = $2`

	_, err := Db.Database.Exec(query,
								a.Name,
								a.Id)

	Db.CheckQueryWithNoRows(err, query)
}

func ListAkademicYears() []AkademicYear {
	var akademicYearList []AkademicYear
	query := `SELECT id, name FROM akademic_year`

	rows, err := Db.Database.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		it := AkademicYear{}
		if err := rows.Scan(&it.Id, &it.Name); err != nil {
			panic(err)
		} else {
			akademicYearList = append(akademicYearList, it)
		}
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		panic(err)
	}

	return akademicYearList
}

