package types

type Book struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Teacher Teacher `json:"teacher"`
}

type BookJson struct {
	Common CommonJson
	Book   Book       `json:"book"`
}

