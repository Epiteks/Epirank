package models

// Student represents the user struct
type Student struct {
	Name      string  `json:"title"`
	Login     string  `json:"login"`
	Bachelor  float64 `json:"bachelor"`
	Master    float64 `json:"master"`
	City      string  `json:"city"`
	Promotion string  `json:"promotion"`
	Position  int     `json:"-"`
}

// GpaToShow returns which gpa we must show for webpage
func (s Student) GpaToShow() float64 {
	if s.Promotion == "tek5" {
		return s.Master
	}
	return s.Bachelor
}
