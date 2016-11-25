package models

// ProfileData represents the json get on profile
type ProfileData struct {
	Gpa []Gpa `json:"gpa"`
}

// Gpa represents the gpa data
type Gpa struct {
	Value string `json:"gpa"`
	Cycle string `json:"cycle"`
}
