package api

import (
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
	"net/http"
    "strconv"
)

func GetStuffList(resp http.ResponseWriter, req *http.Request) {
    var helpers []string
    d := models.ListAllStuff()
    RenderTemplate("stuff/list", helpers, resp, d)
}

func GetStuffCreate(resp http.ResponseWriter, req *http.Request) {
    var helpers []string
    d := models.ListAllDepartments(false)

    teacherType := models.TeacherType

    data := map[string]interface{}{
        "d": d,
        "teacherType": teacherType,
    }

    RenderTemplate("stuff/create", helpers, resp, data)
}

func GetStuffEdit(resp http.ResponseWriter, req *http.Request) {
    var helpers []string

	id, _ := getId(req)

	t := models.Teacher{
		Id: id,
	}

	t.Load()
    teacherType := models.TeacherType

    data := map[string]interface{}{
        "t": t,
        "teacherType": teacherType,
    }

    RenderTemplate("stuff/edit", helpers, resp, data)
}

func StuffEdit(resp http.ResponseWriter, req *http.Request) {
	t := &models.Teacher{}

	req.ParseForm()
	decoder := schema.NewDecoder()

	err := decoder.Decode(t, req.PostForm)

	if err != nil {
		panic(err)
	}

    t.Update()

	url := "/stuff/edit/" + strconv.Itoa(t.Id)
	http.Redirect(resp, req, url, http.StatusMovedPermanently)
}

type TeacherCreateForm struct {
	Id           int     `schema:"id"`
	Name         string  `schema:"name"`
	Email        string  `schema:"email"`
	DepartmentId string  `schema:"department_id"`
    Password     string  `schema:"password"`
    Username   string     `schema:"username"`
    Type       int        `schema:"type"`
}

func StuffCreate(resp http.ResponseWriter, req *http.Request) {
	f := &TeacherCreateForm{}

	req.ParseForm()
	decoder := schema.NewDecoder()

	err := decoder.Decode(f, req.PostForm)

	if err != nil {
		panic(err)
	}

    departmentId,errParse := strconv.Atoi(f.DepartmentId)

    if errParse != nil {
        panic(err)
    }

    d := models.Department{
        Id: departmentId,
    }

    t := models.Teacher{
        Id:         f.Id,
        Name:       f.Name,
        Email:      f.Email,
        Department: d,
        Username:   f.Username,
        Type:       f.Type,
    }

	t.Create(f.Password)

	url := "/stuff/list"
	http.Redirect(resp, req, url, http.StatusMovedPermanently)
}
