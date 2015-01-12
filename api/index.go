package api

import (
	"net/http"
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
    "encoding/json"
)

type LoginForm struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
}

type UserTypeForm struct {
    IsTeacher bool `json:"isTeacher"`
    IsSecretary bool `json:"isSecretary"`
}

func Index(resp http.ResponseWriter, req *http.Request) {
	teacherId := models.GetTeacherFromSession(req).Id

	if teacherId != 0 {
		http.Redirect(resp, req, "/teacher/report/list", http.StatusMovedPermanently)
		return
	}

	var helpers []string
	data := make(map[string]string)
	data["error"] = "NoError"

	RenderTemplate("home", helpers, resp, data)
}

func Login(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	decoder := schema.NewDecoder()
	form := &LoginForm{}
	err := decoder.Decode(form, req.PostForm)
	if err != nil {
		panic(err)
	}

	teacherLogin(form, resp, req)

}

func UserType(resp http.ResponseWriter, req *http.Request) {
    //we need this in order to manipulate the sidebar
    teacherId := models.GetTeacherFromSession(req).Id
    teacherType := models.GetTeacherFromSession(req).Type
    typeForm := UserTypeForm {false, false}

    //unauthorized
    if teacherId == 0 {
        j, _ := json.Marshal(typeForm)
        resp.Write(j)
        return
    } else if teacherType != 7 && teacherType != 8 && teacherType != 9 && teacherType != 11 && teacherType != 12 {
        //type is teacher here
        typeForm.IsTeacher = true
    } else if teacherType != 1 && teacherType != 2 && teacherType != 3 && teacherType != 4 && teacherType != 5 && teacherType != 6 && teacherType != 10 && teacherType != 12 {
        //type is secretary here
        typeForm.IsSecretary = true
    } else if teacherType == 12 {
        //type is admin here, so show both.
        typeForm.IsTeacher = true
        typeForm.IsSecretary = true
    }

    //return the result as json
    j, err := json.Marshal(typeForm)

    if err != nil {
        panic(err)
    }

    //send the json
    resp.Write(j)
}

func teacherLogin(l *LoginForm, resp http.ResponseWriter, req *http.Request) {
	t := models.Teacher{}
	auth := t.Login(l.Username, l.Password)

	if auth {
		session := models.GetSession(req)
		session.Values["teacher"] = &t
		session.Save(req, resp)
		http.Redirect(resp, req, "/teacher/report/list", http.StatusMovedPermanently)
	} else {
		var helpers []string
		data := make(map[string]string)
		data["error"] = "Ο λογαριασμός δεν βρέθηκε. Το όνομα χρήστη ή το συνθηματικό είναι λάθος."

		RenderTemplate("home", helpers, resp, data)
	}
}

func Logout(resp http.ResponseWriter, req *http.Request) {
	session := models.GetSession(req)
	session.Options.MaxAge = -1
	session.Save(req, resp)
	http.Redirect(resp, req, "/", http.StatusTemporaryRedirect)
}

