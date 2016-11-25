package requests

import (
	"encoding/json"
	"fmt"
	"github.com/Shakarang/Epirank/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Shakarang/Epirank/config"
	"github.com/Shakarang/Epirank/requests/urls"

	log "github.com/Sirupsen/logrus"
)

// RequestAllData start getting all ranking data
func RequestAllData(token string) ([]models.Student, error) {

	students, err := requestStudentsList(token)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	requestStudentsGpas(token, &students)

	return students, nil
}

func requestStudentsList(token string) ([]models.Student, error) {
	var students []models.Student

	// Iterate through different cities
	for _, city := range config.Cities {

		log.Info("City : ", city.Name)

		// Iterate through different promotions (tek1,2,3,4,5)
		for _, promotion := range city.Promotions {
			log.Info("\t", promotion)
			if data, err := retrievePromotion(token, city.ID, promotion); err != nil {
				log.WithFields(log.Fields{
					"City":      city.ID,
					"Promotion": promotion,
				}).Warning(err)
			} else {
				// Concatenate new data with the current one
				students = append(students, data...)
			}
		}
	}

	fmt.Printf("%v\n%v\n", students, len(students))

	return students, nil
}

// retrievePromotion retrieve list of students by promotion/city
func retrievePromotion(token, cityID, promotion string) ([]models.Student, error) {

	var offset = 0
	var totalStudents = 0

	// Students data
	var students []models.Student

	// Yearbook containing temporary data and total students for promotion
	var yearbook models.Yearbook

	for len(students) < totalStudents || totalStudents == 0 {

		var requestURL = fmt.Sprintf(urls.UserList, urls.EpitechIntranet, cityID, promotion, offset)

		log.Info("Request : ", requestURL)

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

func requestStudentsGpas(token string, students *[]models.Student) error {

	log.Info("Hello world")

	for i, student := range *students {

		// Create GET request with required header
		request, err := http.NewRequest("GET", urls.EpitechAPIProfile, nil)
		request.Header.Add("token", token)
		request.Header.Add("login", student.Login)

		fmt.Printf("Request gpa\n")

		if err != nil {
			return err
		}

		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		// Read body data to []byte
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			log.WithFields(log.Fields{
				"Student": student.Name,
			}).Error(err)
		} else {

			fmt.Printf("Body : \n%v\n", string(body))

			var profileData models.ProfileData

			// Unmarshal JSON to put it into profileData object
			if err := json.Unmarshal(body, &profileData); err != nil {
				log.WithFields(log.Fields{
					"Student": student.Name,
				}).Error(err)
			} else {

				for _, gpa := range profileData.Gpa {
					if gpa.Cycle == "bachelor" {
						(*students)[i].Bachelor, _ = strconv.ParseFloat(gpa.Value, 64)
					} else if gpa.Cycle == "master" {
						(*students)[i].Master, _ = strconv.ParseFloat(gpa.Value, 64)
					}
				}
				fmt.Println(student)
			}
		}
	}
	fmt.Printf("---------\n%v\n----------\n", students)
	return nil

}
