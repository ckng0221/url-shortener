package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"url-shortener/initializers"
	"url-shortener/models"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SynDatabase()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello URL shortener")

	})

	r.GET("/url-shortener", func(c *gin.Context) {
		var url models.Url
		queryUrl := c.Query("url")
		err := initializers.Db.Where("url = ?", queryUrl).Find(&url).Error
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, gin.H{"shorten_url": url.ShortenUrl})
	})

	r.POST("/url-shortener", func(c *gin.Context) {
		// Create url in db
		var url models.Url
		err := c.ShouldBind(&url)
		if err != nil {
			c.AbortWithStatusJSON(400, map[string]any{"error": "invalid request body"})
			return
		}
		err = initializers.Db.Clauses(clause.Returning{}).Create(&url).Error
		if err != nil {
			log.Println(err.Error())
			var mysqlErr *mysql.MySQLError
			if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
				c.AbortWithStatusJSON(400, gin.H{"error": "URL already exists"})
				return
			}

			c.AbortWithStatus(500)
			return
		}

		baseUrl := os.Getenv("BASE_URL")
		shortenUrlPath := convertBase62(int(url.ID))
		url.ShortenUrl = fmt.Sprintf("%s/urls/%s", baseUrl, shortenUrlPath)
		err = initializers.Db.Updates(&url).Error
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(201, url)
	})

	// Redirect
	r.GET("/url/:path", func(c *gin.Context) {
		//TODO: reverse path to integer ID
		//TODO: Get id and return
	})

	r.Run()
}
