package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Shakarang/Epirank/models"

	"github.com/Shakarang/Epirank/config"
	"github.com/Shakarang/Epirank/requests/urls"

	log "github.com/Sirupsen/logrus"
)

// RequestAllData start getting all ranking data
func RequestAllData(token string) error {

	// Iterate through different cities
	for _, city := range config.Cities {
		fmt.Printf("City : %v\n", city.Name)

		// Iterate through different promotions
		for _, promotion := range city.Promotions {
			fmt.Printf("\t%v\n", promotion)
			if data, err := retrievePromotion(token, city.ID, promotion); err != nil {
				log.WithFields(log.Fields{
					"City":      city.ID,
					"Promotion": promotion,
				}).Warning(err)
			} else {
				fmt.Printf("%v\n%v\n", data, len(data))
			}
			//break
		}
	}
	return nil
}

// retrievePromotion retrieve list of students by promotion/city
func retrievePromotion(token, cityID, promotion string) ([]models.Student, error) {

	var offset = 0
	var totalStudents = 0
	//var requestURL = fmt.Sprintf(urls.UserList, urls.EpitechIntranet, cityID, promotion, offset)

	//fmt.Printf("Request : %v\n", requestURL)

	var students []models.Student

	var yearbook models.Yearbook

	for len(students) < totalStudents || totalStudents == 0 {

		var requestURL = fmt.Sprintf(urls.UserList, urls.EpitechIntranet, cityID, promotion, offset)
		fmt.Printf("Request : %v\n", requestURL)
		// Create GET request with required header
		request, err := http.NewRequest("GET", requestURL, nil)
		request.Header.Add("Cookie", fmt.Sprintf("PHPSESSID=%v", token))

		if err != nil {
			return nil, err
		}

		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		// Read body data to []byte
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			log.WithFields(log.Fields{
				"City":      cityID,
				"Promotion": promotion,
				"Offset":    offset,
			}).Error(err)
		} else {

			// Unmarshal JSON to put it into yearbook object
			if err := json.Unmarshal(body, &yearbook); err != nil {
				log.WithFields(log.Fields{
					"City":      cityID,
					"Promotion": promotion,
					"Offset":    offset,
				}).Error(err)
			} else {

				totalStudents = yearbook.TotalStudents

				for _, student := range yearbook.Students {

					student.City = cityID
					student.Promotion = promotion

					students = append(students, student)
				}

				offset = len(students)

			}
		}
	}

	return students, nil
}
