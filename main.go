package main

import "github.com/jinzhu/gorm"

var db *gorm.DB
var err error

type Laptop struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

func main() {

	const CnString = "host=localhost port=5432 user=gorm dbname=person_db sslmode=disable password=gormPassword"
	db, err = gorm.Open("postgres", CnString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

}
