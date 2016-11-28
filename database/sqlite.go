package database

import (
	"database/sql"
	//"fmt"
	"github.com/Shakarang/Epirank/models"
	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3" // SQLITE
)

// Init database
func Init(path string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateTable creates student table if not existing in current database
func CreateTable(db *sql.DB) {

	sqlTable := `CREATE TABLE IF NOT EXISTS students(Login TEXT NOT NULL PRIMARY KEY,
			Name TEXT NOT NULL,
			Bachelor FLOAT NOT NULL,
			Master FLOAT NOT NULL,
			City TEXT NOT NULL,
			Promotion TEXT NOT NULL);
			`

	if _, err := db.Exec(sqlTable); err != nil {
		log.Panic("Create table : ", err)
	}
}

// InsertData insert students in database
// Replace existing value with new one if present
func InsertData(db *sql.DB, students []models.Student) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT OR REPLACE INTO students(Login, Name, Bachelor, Master, City, Promotion) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, student := range students {
		_, err = stmt.Exec(student.Login, student.Name, student.Bachelor, student.Master, student.City, student.Promotion)
		if err != nil {
			log.Error("Inserting : ", err)
			return err
		}
	}
	tx.Commit()
	return nil
}

// GetStudentsFrom retrieve students by their Promotion and City
// Orders them by their Bachelor's gpa if they are in Tek[1,2,3,4]
// and the Master one if they are in Tek5
func GetStudentsFrom(db *sql.DB, city, promotion *string) []models.Student {

	var students []models.Student

	if promotion == nil || len(*promotion) == 0 {
		*promotion = "tek1"
	}

	var sqlQuery = "select Name, Login, Bachelor, Master, City from students WHERE Promotion = ? "

	var rows *sql.Rows
	var err error

	var orderByQuery = "ORDER BY "

	if *promotion == "tek5" {
		orderByQuery += "Master"
	} else {
		orderByQuery += "Bachelor"
	}

	orderByQuery += " DESC"

	if city != nil && len(*city) > 0 {
		sqlQuery += "AND City = ? "
		sqlQuery += orderByQuery
		rows, err = db.Query(sqlQuery, promotion, city)
	} else {
		sqlQuery += orderByQuery
		rows, err = db.Query(sqlQuery, promotion)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		var student = models.Student{
			Promotion: *promotion,
		}

		err = rows.Scan(&student.Name, &student.Login, &student.Bachelor, &student.Master, &student.City)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return students
}
