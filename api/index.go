package api

import (
	"net/http"
	"github.com/kokeroulis/modip/models"
	"github.com/gorilla/schema"
)

type LoginForm struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
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

