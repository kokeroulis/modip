package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"github.com/kokeroulis/modip/models"
)

func departmentAlreadyExists(candidate string) bool {
	for _, it := range departments {
		if candidate == it.Name {
			return true
		}
	}

	return false
}

func findDepartmentByName(departmentName string) *models.Department {
	for _, it := range departments {
		if departmentName == it.Name {
			return it
		}
	}

	panic("UNREACHABLE")
}

func readFile(fileName string) []string {
	path := "modip-data/" + fileName
	b, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	contents := string(b)
	l := strings.Split(contents, "\n")
	return l
}

func exportDepartments() {
	for _, it := range readFile(departmentsFile) {
		l := strings.Split(it, ";")

		var d *models.Department

		if len(l) > 5 {
			departmentName := strings.TrimSpace(l[5])

			if !departmentAlreadyExists(departmentName) {

				d = &models.Department {
					Name: departmentName,
				}
				departments = append(departments, d)
			} else {
				d = findDepartmentByName(departmentName)
			}

			exportTeachers(d, it)
		}
	}

	exportLessons()

	for _, it :=range departments {
		fmt.Println(it)
	}
}

func exportTeachers(d *models.Department, data string) {
	l := strings.Split(data, ";")

	t := &models.Teacher{
		Name:  l[1],
//		Attribute: l[2]
		Email: l[3],
	}

	fmt.Println(t)
	d.AddTeacher(t)
}

func exportLessons() {
	for _, it := range readFile(lessonsFile) {
		l := strings.Split(it, ";")

		if len(l) >= 7 {
			departmentName := strings.TrimSpace(l[2])

			var d *models.Department
			d = findDepartmentByName(departmentName)

			lesson := &models.Lesson{
				Name:       l[5],
				CourseCode: l[4],
			}

			d.AddLesson(lesson)
		}
	}
}

