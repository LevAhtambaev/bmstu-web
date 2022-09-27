package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		id := c.Query("id") // получаем из запроса query string
		create := c.Query("create")
		if create != "" {
			log.Printf("create received %s\n", create)
			createBool, err := strconv.ParseBool(create)
			if err != nil {
				log.Printf("cant convert create %v", err)
				c.Error(err)
				return
			}
			if createBool {
				a.repo.NewRandManga()
				c.JSON(http.StatusOK, gin.H{
					"status": "OK",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "Why u bulling me?",
				})
			}
		}
		if id != "" {
			log.Printf("id recived %s\n", id)
			intID, err := strconv.Atoi(id) // пытаемся привести это к чиселке
			if err != nil {                // если не получилось
				log.Printf("cant convert id %v", err)
				c.Error(err)
				return
			}

			product, err := a.repo.GetMangaByID(uint(intID))
			if err != nil { // если не получилось
				log.Printf("cant get product by id %v", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"manga_name":  product.Name,
				"manga_year":  product.Year,
				"manga_genre": product.Genre,
				"manga_price": product.Price,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.tmpl", gin.H{
			"title": "Main website",
			"test":  []string{"a", "b"},
		})
	})

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
