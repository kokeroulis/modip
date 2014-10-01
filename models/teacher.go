package models

type Teacher struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Department Department `json:"department"`
}

type TeacherInfo struct {
	Teacher Teacher           `json:"teacher"`
	Books   []BookOrPaperInfo `json:"books"`
	Papers  []BookOrPaperInfo `json:"papers"`
}

