package datasource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"models"

	"github.com/jinzhu/gorm"
	/* Postgres dialect is used */
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

type queries struct {
	Extensions  []string
	Constraints []string
}

func InitDB() {
	var err error

	connectionString := "host=192.168.2.45 port=5432 user=postgres password=ubuntu dbname=school_management sslmode=disable"
	driverName := "postgres"

	Db, err = gorm.Open(driverName, connectionString)

	if err != nil {
		log.Fatal("Failed to connect database !!!", err)
	}
	defer Db.Close()

	jsonData, err := ioutil.ReadFile("src/datasource/queries.json")
	if err != nil {
		fmt.Println("Error reading JSON data", err)
		return
	}

	queries := new(queries)

	err = json.Unmarshal(jsonData, &queries)
	if err != nil {
		fmt.Println("Error while unmarshaling data", err)
		return
	}

	for item := range queries.Extensions {
		Db.Exec(queries.Extensions[item])
	}

	Db.AutoMigrate(&models.Group{}, &models.Mark{}, &models.Student{},
		&models.SubjectTeacher{}, &models.Subject{}, &models.Teacher{})

	for item := range queries.Constraints {
		Db.Exec(queries.Constraints[item])
	}

}
