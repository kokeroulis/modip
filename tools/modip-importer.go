package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kokeroulis/modip/db"
	"github.com/kokeroulis/modip/models"
	"io/ioutil"
)

type teacher struct {
	Name     string
	Email    string
	Password string
    Username string
    Type     int
}

type teacherDb struct {
	Id           int
	Name         string
	Email        string
	Password     string
	DepartmentId int
}

type department struct {
	Name     string
	Teachers []teacher
}

type jsonData struct {
	Departments []department
}

func checkQuery(err error) bool {
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
		return false
	default:
		return true
	}
}

func populateDb(data jsonData) {
	Db.Connect()
	db := Db.Database

	for _, department := range data.Departments {

		var departmentId int
		err := db.QueryRow("SELECT id FROM department WHERE name=$1", department.Name).Scan(&departmentId)
		if !checkQuery(err) {
			// department doesn't exist, create one
			err := db.QueryRow("INSERT INTO department (name) VALUES ($1) RETURNING id", department.Name).Scan(&departmentId)
			if !checkQuery(err) {
				fmt.Println("Error %s", err)
				panic("WRONG QUERY!")
			}
		}

		createTeachers(db, department.Teachers, departmentId)
	}
}

func createTeachers(db *sql.DB, teachers []teacher, departmentId int) {
	for _, teacher := range teachers {

		var teacherName string
		var teacherEmail string
		var teacherPassword string
		var teacherDepartment int
		var teacherId int
        var teacherUsername string
        var teacherType string

		query := `SELECT id, name, email, password, department, username, type
				  FROM teacher WHERE name=$1`

		err := db.QueryRow(query, teacher.Name).Scan(&teacherId, &teacherName, &teacherEmail, &teacherPassword, &teacherDepartment, &teacherUsername, &teacherType)

		if !checkQuery(err) {
			// teacher doesn't exist, create one

			d := models.Department{
				Id: departmentId,
			}

			t := models.Teacher{
				Name:       teacher.Name,
				Email:      teacher.Email,
                Username:   teacher.Username,
                Type:       teacher.Type,
				Department: d,
			}

			t.Create(teacher.Password)
		}
	}
}

func main() {
	jsonFile := flag.String("data", "", "A json data file")
	flag.Parse()

	if *jsonFile == "" {
		fmt.Println("File is missing")
		fmt.Println("Execute it like, `modip-importer --data=testdata.json`")
		return
	}

	fileData, fileErr := ioutil.ReadFile(*jsonFile)
	if fileErr != nil {
		fmt.Println("Can't read file %s", *jsonFile)
	}

	fmt.Println(string(fileData))

	j := jsonData{}
	err := json.Unmarshal(fileData, &j)

	if err != nil {
		fmt.Println(("Can't decode json"))
		fmt.Println(err)
		return
	}

	populateDb(j)
}
