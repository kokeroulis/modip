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

type TeacherInfo struct {
	Teacher Teacher           `json:"teacher"`
	Books   []BookOrPaperInfo `json:"books"`
	Papers  []BookOrPaperInfo `json:"papers"`
}

type TeacherInfoJson struct {
	CommonJson  CommonJson
	TeacherInfo TeacherInfo `json:"teacher"`
}
