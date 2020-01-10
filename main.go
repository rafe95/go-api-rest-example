package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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

	db.AutoMigrate(&Laptop{})

	r := gin.Default()
	r.GET("/laptop/", GetAll)
	r.GET("/laptop/:id", GetLaptop)
	r.POST("/laptop", AddLaptop)
	r.PUT("/laptop/:id", UpdateLaptop)
	r.DELETE("/laptop/:id", DeleteLaptop)

	r.Run(":8080")

}

func DeleteLaptop(c *gin.Context) {
}

func UpdateLaptop(c *gin.Context) {
}

func AddLaptop(c *gin.Context) {
	var laptop Laptop
	c.BindJSON(&laptop)

	db.Create(&laptop)
	c.JSON(200, laptop)
}

func GetAll(c *gin.Context) {
}

func GetLaptop(c *gin.Context) {
}
