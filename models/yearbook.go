package models

import (
//"encoding/json"
//"fmt"
)

// Yearbook represent data returned from yearbook requests
type Yearbook struct {
	TotalStudents int       `json:"total"`
	Students      []Student `json:"items"`
}

// // UnmarshalJSON lol
// func (y *Yearbook) UnmarshalJSON(data []byte) error {
// 	if err := json.Unmarshal(data, &y); err != nil {
// 		return err
// 	}
// 	return nil
// }
