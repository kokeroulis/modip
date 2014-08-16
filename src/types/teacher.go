package types

type TeacherJson struct {
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Department DepartmentJson `json:"department"`
}

