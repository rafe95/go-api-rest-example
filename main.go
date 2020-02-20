package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"strconv"
)

var err error

type Laptop struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

func main() {

	const CnString = "host=localhost port=5432 user=gorm dbname=laptop_db sslmode=disable password=gormPassword"

	db, err := sql.Open("postgres", CnString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/laptop/", func(c *gin.Context) {
		sql := `select * from laptop`
		all, err := db.Query(sql)
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, all)
		}
	})

	r.GET("/laptop/:id", func(c *gin.Context) {

		id := c.Params.ByName("id")
		sql := `select * from laptop where id = ` + id
		laptop, err := db.Query(sql)
		if err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, laptop)
		}

	})

	r.POST("/laptop", func(c *gin.Context) {

		var laptop Laptop
		c.BindJSON(&laptop)
		id := strconv.Itoa(laptop.ID)
		sql := `insert into laptop values(` + id + `,` + laptop.Brand + `,` + laptop.Model + `)`
		db.Query(sql)
		c.JSON(200, laptop)

	})

	r.PUT("/laptop/:id", func(c *gin.Context) {

		id := c.Params.ByName("id")
		var laptop Laptop
		c.BindJSON(&laptop)

		sql := `update laptop set brand =` + laptop.Brand + `,` + `model =` + laptop.Model + ` where id =` + id

		db.Query(sql)
		c.JSON(200, laptop)

	})

	r.DELETE("/laptop/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		sql := `delete from laptop where id = ` + id
		db.Query(sql)
		c.JSON(200, gin.H{})
	})

	r.Run(":8080")

}
