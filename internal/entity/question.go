package entity

type Question struct {
	ID         int
	Title      string
	Slug       string
	Content    string
	Difficulty string
	Tags       []string
}
