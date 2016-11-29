package main

import (
	"database/sql"
	"github.com/Epiteks/Epirank/config"
	"github.com/Epiteks/Epirank/database"
	"github.com/Epiteks/Epirank/ranking"
	"github.com/Epiteks/Epirank/routes"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// APIMiddleware will add the db connection to the context
func APIMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
		c.Next()
	}
}

func init() {

	// Log as ASCII Formatter
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout, could also be a file.
	log.SetOutput(os.Stdout)

	// Log everything
	log.SetLevel(log.InfoLevel)
}

func main() {

	if db, err := database.Init(config.DatabasePath); err != nil {
		log.Fatal(err)
	} else {

		defer db.Close()
		database.CreateTable(db)

		if err := ranking.InitRanking(db); err != nil {
			log.Error(err)
			os.Exit(-1)
		}

		// Webservice
		router := gin.Default()

		router.Use(APIMiddleware(db))

		gopath := os.Getenv("GOPATH")
		gopath += "/src/github.com/Epiteks/Epirank"

		router.LoadHTMLGlob(gopath + "/front/html/*")
		router.Static("/css", gopath+"/front/css")
		router.Static("/js", gopath+"/front/js")
		router.StaticFile("/favicon.ico", gopath+"/front/icons/favicon-96x96.png")

		router.GET("/", routes.GetStudents)

		router.Run()
	}
}
