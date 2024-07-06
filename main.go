package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"url-shortener/initializers"
	"url-shortener/models"
	"url-shortener/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SynDatabase()
}

func main() {
	hashTable := make(map[int]int)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello URL shortener")
	})

	r.GET("/shorten-urls", func(c *gin.Context) {
		var url models.Url
		queryUrl := c.Query("url")
		err := initializers.Db.Where("url = ?", queryUrl).Find(&url).Error
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatus(500)
			return
		}
		if url.ID == "" {
			c.Status(200)
			return
		}
		baseUrl := os.Getenv("BASE_URL")
		shortenUrlWithBase := fmt.Sprintf("%s/urls/%s", baseUrl, url.ShortenUrl)

		c.JSON(200, gin.H{"shorten_url": shortenUrlWithBase})
	})

	r.GET("/shorten-urls/:id", func(c *gin.Context) {
		var url models.Url
		id := c.Param("id")
		err := initializers.Db.First(&url, id).Error
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(404, gin.H{"error": "ID not found"})
				return
			}
			c.AbortWithStatus(500)
			return
		}
		baseUrl := os.Getenv("BASE_URL")
		shortenUrlWithBase := fmt.Sprintf("%s/urls/%s", baseUrl, url.ShortenUrl)
		url.ShortenUrl = shortenUrlWithBase

		c.JSON(200, url)
	})

	r.POST("/shorten-urls", func(c *gin.Context) {
		// Create url in db
		var url models.Url
		err := c.ShouldBind(&url)
		if err != nil {
			log.Print(err.Error())
			c.AbortWithStatusJSON(400, map[string]any{"error": "invalid request body"})
			return
		}
		idString := utils.IdGenerator(&hashTable)
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Print(err.Error())
			c.AbortWithStatus(500)
			return
		}

		shortenUrlPath := utils.ConvertIntegerToBase62(id)
		url.ShortenUrl = shortenUrlPath

		url.ID = idString

		err = initializers.Db.Create(&url).Error
		if err != nil {
			log.Println(err.Error())
			// var mysqlErr *mysql.MySQLError
			// if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// 	c.AbortWithStatusJSON(400, gin.H{"error": "URL already exists"})
			// 	return
			// }

			c.AbortWithStatus(500)
			return
		}

		baseUrl := os.Getenv("BASE_URL")
		shortenUrlWithBase := fmt.Sprintf("%s/urls/%s", baseUrl, shortenUrlPath)
		url.ShortenUrl = shortenUrlWithBase

		c.JSON(201, url)
	})

	// Redirect
	r.GET("/urls/:path", func(c *gin.Context) {
		var url models.Url
		path := c.Param("path")

		id := utils.ConvertBase62ToInteger(path)
		// fmt.Println("id", id)

		err := initializers.Db.First(&url, id).Error
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(404, gin.H{"error": "url not found"})
				return
			}
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// update usage count
		expression := "usage_count + 1"
		initializers.Db.Model(&url).UpdateColumn("usage_count", gorm.Expr(expression))
		initializers.Db.Model(&url).Update("last_accessed", time.Now())

		c.Redirect(http.StatusTemporaryRedirect, url.Url)
	})

	go utils.ResetHashTable(&hashTable)
	r.Run()
}
