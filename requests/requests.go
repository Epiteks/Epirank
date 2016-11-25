package requests

import (
	"encoding/json"
	"fmt"
	"github.com/Shakarang/Epirank/config"
	"github.com/Shakarang/Epirank/models"
	"github.com/Shakarang/Epirank/requests/urls"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

// RequestAllData start getting all ranking data
func RequestAllData(token string) ([]models.Student, error) {

	students, err := requestStudentsList(token)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//requestStudentsGpas(token, &students)
	requestsPool(token, &students)
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

func requestsPool(token string, students *[]models.Student) {

	var studentsNumber = len(*students)

	jobs := make(chan *models.Student, studentsNumber)
	results := make(chan *models.Student, studentsNumber)

	// This starts up 5 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= 5; w++ {
		go requestGpa(w, token, jobs, results)
	}

	// Here we send 9 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 0; j < studentsNumber; j++ {
		jobs <- &(*students)[j]
	}
	close(jobs)

	// Finally we collect all the results of the work.
	for a := 0; a < studentsNumber; a++ {
		student := <-results
		log.Info("Student :", student)
	}
}

func requestGpa(id int, token string, jobs <-chan *models.Student, results chan<- *models.Student) {

	for student := range jobs {

		// Create GET request with required header
		request, err := http.NewRequest("GET", urls.EpitechAPIProfile, nil)
		request.Header.Add("token", token)
		request.Header.Add("login", student.Login)

		if err != nil {
			log.Error(err)
		}

		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Error(err)
		}

		defer resp.Body.Close()

		// Read body data to []byte
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			log.WithFields(log.Fields{
				"Student": student.Name,
			}).Error(err)
		} else {
			var profileData models.ProfileData

			// Unmarshal JSON to put it into profileData object
			if err := json.Unmarshal(body, &profileData); err != nil {
				log.WithFields(log.Fields{
					"Student": student.Name,
				}).Error(err)
			} else {

				for _, gpa := range profileData.Gpa {
					if gpa.Cycle == "bachelor" {
						student.Bachelor, _ = strconv.ParseFloat(gpa.Value, 64)
					} else if gpa.Cycle == "master" {
						student.Master, _ = strconv.ParseFloat(gpa.Value, 64)
					}
				}
			}
		}
		results <- student
	}
}
