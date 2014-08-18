package types

type Teacher struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Department Department `json:"department"`
}

type TeacherJson struct {
	Common  CommonJson
	Teacher Teacher `json:"teacher"`
}

