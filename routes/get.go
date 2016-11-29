package routes

import (
	"database/sql"
	"github.com/Epiteks/Epirank/config"
	"github.com/Epiteks/Epirank/database"
	"github.com/Epiteks/Epirank/models"
	"github.com/Epiteks/Epirank/ranking"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type htmlModel struct {
	Students              []models.Student
	CurrentCity           string
	Cities                []models.City
	GpaType               string
	CurrentData           string
	StudentsUpdateMessage string
}

// GetStudents returns students from :
// city : in Query
// promotion : in Query
func GetStudents(c *gin.Context) {

	format := c.Query("format")
	promo := c.Query("promotion")
	cityID := c.Query("city")
	db, _ := c.Get("database")

	if len(promo) == 0 {
		promo = "tek1"
	}

	var studentsUpdateMessage = "The ranking is updated every day around 3am. The last one was on : "

	studentsUpdateMessage += ranking.LastRankUpdate.Format("2 January 2006 15:04")

	students := database.GetStudentsFrom(db.(*sql.DB), &cityID, &promo)

	if format == "json" {
		if len(students) == 0 {
			c.Status(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"updatedAt": ranking.LastRankUpdate,
				"students":  students,
			})
		}
	} else {

		var gpaType = "Bachelor"

		if promo == "tek5" {
			gpaType = "Master"
		}

		var data string

		for _, city := range config.Cities {
			if city.ID == cityID {
				data = city.Name
				break
			}
		}

		data += " "
		data += strings.ToUpper(promo)

		model := htmlModel{
			Students:              students,
			Cities:                config.Cities,
			GpaType:               gpaType,
			CurrentData:           data,
			StudentsUpdateMessage: studentsUpdateMessage,
		}
		c.HTML(http.StatusOK, "body.html", model)
	}
}
