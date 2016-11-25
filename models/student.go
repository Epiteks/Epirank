package models

// Student represents the user struct
type Student struct {
	Name      string  `json:"title"`
	Login     string  `json:"login"`
	Bachelor  float64 `json:"-"`
	Master    float64 `json:"-"`
	City      string  `json:"-"`
	Promotion string  `json:"-"`
}
