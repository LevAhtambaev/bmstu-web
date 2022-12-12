package app

import (
	"WAD-2022/internal/app/role"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/comics", a.GetList)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/role", a.Role)
	r.GET("/logout", a.Logout)

	//r.GET("/comics/price/:uuid", a.GetMangaPrice)

	r.GET("/comics/:uuid", a.GetComics)

	r.POST("/cart", a.AddToCart)
	r.POST("/login", a.Login)
	r.POST("/sign_up", a.Register)

	r.DELETE("/cart/:uuid", a.DeleteFromCart)
	r.Use(a.WithAuthCheck(role.Buyer, role.Manager)).POST("/orders", a.AddOrder)
	r.Use(a.WithAuthCheck(role.Buyer, role.Manager, role.Admin)).GET("/cart", a.GetCart)
	r.Use(a.WithAuthCheck(role.Manager)).GET("/user/:uuid", a.GetUser)
	r.Use(a.WithAuthCheck(role.Manager)).PUT("/comics/:uuid", a.ChangeComics)
	r.Use(a.WithAuthCheck(role.Manager)).GET("/orders", a.GetOrders)
	r.Use(a.WithAuthCheck(role.Manager)).POST("/comics", a.AddComics)
	r.Use(a.WithAuthCheck(role.Manager)).DELETE("comics/:uuid", a.DeleteComics)
	r.Use(a.WithAuthCheck(role.Manager)).PUT("/orders/:uuid", a.ChangeStatus)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
