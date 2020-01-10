package main

import (
	"fmt"
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
	id := c.Params.ByName("id")
	var laptop Laptop
	d := db.Where("id = ?", id).Delete(&laptop)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateLaptop(c *gin.Context) {
	var laptop Laptop
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&laptop).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&laptop)

	db.Save(&laptop)
	c.JSON(200, laptop)
}

func AddLaptop(c *gin.Context) {
	var laptop Laptop
	c.BindJSON(&laptop)

	db.Create(&laptop)
	c.JSON(200, laptop)
}

func GetAll(c *gin.Context) {
	var all []Laptop
	if err := db.Find(&all).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, all)
	}
}

func GetLaptop(c *gin.Context) {
	id := c.Params.ByName("id")
	var laptop Laptop
	if err := db.Where("id = ?", id).First(&laptop).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, laptop)
	}
}
