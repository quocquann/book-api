package models

type Book struct {
	Isbn    string `json:"isbn"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  string `json:"author"`
	Genre   string `json:"genre"`
}
