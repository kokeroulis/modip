package types

type Paper struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Teacher Teacher `json:"teacher"`
}

