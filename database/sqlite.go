package database

import (
	"database/sql"
	"fmt"
	"github.com/Shakarang/Epirank/models"
	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3" // Hello
)

// Init database
func Init(path string) (*sql.DB, error) {

	if db, err := sql.Open("sqlite3", path); err != nil {
		return nil, err
	} else {
		return db, nil
	}
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

// GetStudentsFrom retrieve students by their GPA and city
func GetStudentsFrom(db *sql.DB, city, promotion *string) []models.Student {

	var students []models.Student

	var sqlQuery = "select Name, Login, Bachelor, Master from students WHERE Promotion = ? "

	if promotion != nil {
		sqlQuery += "AND City = ? "
	}
	log.Info(sqlQuery)
	rows, err := db.Query(sqlQuery, promotion, city)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		var student = models.Student{
			City:      *city,
			Promotion: *promotion,
		}

		err = rows.Scan(&student.Name, &student.Login, &student.Bachelor, &student.Master)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(student)
		students = append(students, student)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return students
}
