package models

// Student represents the user struct
type Student struct {
	Name      string `json:"title"`
	Login     string `json:"login"`
	Gpa       string `json:"-"`
	City      string `json:"-"`
	Promotion string `json:"-"`
}
