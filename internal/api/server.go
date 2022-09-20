package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Gift struct {
	Name     string
	Priority int
}

func StartServer() {
	giftList := []Gift{
		{
			"dildo",
			1,
		},
		{
			"hentai manga",
			3,
		},
		{
			"wmal'",
			5,
		},
		{
			"katana",
			2,
		},
	}

	quickSort(giftList, 0, len(giftList)-1)
	log.Println("Server starts up")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/wish_list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "giftList.tmpl", gin.H{
			"title":    "Wish List",
			"giftList": giftList,
		})
	})

	r.Static("/image", "./resourses")

	r.Run()
	log.Println("Server is down")
}
