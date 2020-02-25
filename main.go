package main

import (
	"database/sql"
	"flag"
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

	var (
		host     string
		port     int
		user     string
		dbname   string
		sslmode  string
		password string
		appPort  int
	)

	flag.StringVar(&host, "host", "localhost", "Db host address")
	flag.IntVar(&port, "port", 5432, "Host connection port")
	flag.StringVar(&user, "user", "gorm", "The username of the database owner")
	flag.StringVar(&dbname, "db", "laptop_db", "Database name on host")
	flag.StringVar(&sslmode, "ssl", "disable", "SSL connection mode, default disable")
	flag.StringVar(&password, "pwd", "gormPassword", "Password for the database user owner")
	flag.IntVar(&appPort, "apprt", 8080, "Port for running LaptopAPI")
	flag.Usage()
	flag.Parse()

	var CnString = "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " dbname=" + dbname + " sslmode=" + sslmode + " password=" + password

	db, err := sql.Open("postgres", CnString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/laptop/", func(c *gin.Context) {
		selectAll := "SELECT * FROM laptops"
		rows, _ := db.Query(selectAll)
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
		selectById := "SELECT * FROM laptops WHERE id = $1"
		rows, err := db.Query(selectById, queryId)
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
		insert := "INSERT INTO laptops (id, brand, model) VALUES ( $1, $2, $3)"
		db.Exec(insert, laptop.ID, laptop.Brand, laptop.Model)
		c.JSON(201, laptop)

	})

	r.PUT("/laptop/:id", func(c *gin.Context) {
		var laptop Laptop
		c.BindJSON(&laptop)
		update := "UPDATE laptops SET brand = $1, model = $2 WHERE id = $3"
		db.Exec(update, laptop.Brand, laptop.Model, laptop.ID)
		c.JSON(200, laptop)

	})

	r.DELETE("/laptop/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		delete := "DELETE FROM laptops WHERE id = $1"
		db.Exec(delete, id)
		c.JSON(200, gin.H{"msg": "Element with id:" + id + " has been deleted"})
	})

	r.Run(":" + strconv.Itoa(appPort))

}
