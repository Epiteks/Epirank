package routes

import (
	"database/sql"
	"github.com/Shakarang/Epirank/config"
	"github.com/Shakarang/Epirank/database"
	"github.com/Shakarang/Epirank/models"
	"github.com/Shakarang/Epirank/ranking"
	"github.com/gin-gonic/gin"
	"net/http"
)

type htmlModel struct {
	Students    []models.Student
	CurrentCity string
	Cities      []models.City
	Promotions  []string
}

// GetStudents returns students from :
// city : in Query
// promotion : in Query
func GetStudents(c *gin.Context) {

	format := c.Query("format")
	promo := c.Query("promotion")
	cityID := c.Query("city")
	db, _ := c.Get("database")
	students := database.GetStudentsFrom(db.(*sql.DB), &cityID, &promo)

	var promotions []string

	for _, city := range config.Cities {
		if city.ID == cityID {
			promotions = city.Promotions
		}
	}

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

		model := htmlModel{
			Students:    students,
			CurrentCity: cityID,
			Cities:      config.Cities,
			Promotions:  promotions,
		}
		c.HTML(http.StatusOK, "body.html", &model)
	}
}
