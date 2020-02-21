package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
		sql := "SELECT * FROM laptops"
		rows, _ := db.Query(sql)
		all := []Laptop{}

		var single Laptop
		for rows.Next() {
			var (
				id    int
				model string
				brand string
			)
			rows.Scan(&id, &brand, &model)
			single = Laptop{
				ID:    id,
				Brand: brand,
				Model: model,
			}
			all = append(all, single)

		}

		if err != nil {
			c.AbortWithStatus(404)
		} else {
			if len(all) == 0 {
				c.JSON(404, "404 Not found!")
			} else {
				c.JSON(200, all)
			}
		}
	})

	r.GET("/laptop/:id", func(c *gin.Context) {

		queryId := c.Params.ByName("id")
		sql := "SELECT * FROM laptops WHERE id =" + queryId
		rows, err := db.Query(sql)
		defer rows.Close()

		if err != nil {
			c.AbortWithStatus(404)
		} else {
			var single Laptop
			var (
				id    int
				brand string
				model string
			)

			rows.Next()
			rows.Scan(&id, &brand, &model)
			single = Laptop{
				ID:    id,
				Brand: brand,
				Model: model,
			}
			if single.ID == 0 {
				c.JSON(404, "404 Not found")

			} else {
				c.JSON(200, single)
			}

		}
	})

	r.POST("/laptop/", func(c *gin.Context) {
		var laptop Laptop
		c.BindJSON(&laptop)
		sql := "INSERT INTO laptops (id, brand, model) VALUES ( $1, $2, $3)"
		db.Exec(sql, laptop.ID, laptop.Brand, laptop.Model)
		c.JSON(201, laptop)

	})

	r.PUT("/laptop/:id", func(c *gin.Context) {
		var laptop Laptop
		c.BindJSON(&laptop)
		sql := "UPDATE laptops SET brand = $1, model = $2 WHERE id = $3"
		db.Exec(sql, laptop.Brand, laptop.Model, laptop.ID)
		c.JSON(200, laptop)

	})

	r.DELETE("/laptop/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		sql := "DELETE FROM laptops WHERE id = $1"
		db.Exec(sql, id)
		c.JSON(200, gin.H{"msg": "Element with id:" + id + " has been deleted"})
	})

	r.Run(":8080")

}
