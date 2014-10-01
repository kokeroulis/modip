package models

type Book struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Teacher Teacher `json:"teacher"`
}

