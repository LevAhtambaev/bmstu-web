package app

import (
	"WAD-2022/internal/app/ds"
	"WAD-2022/swagger/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

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

	r.GET("/manga/", a.GetList)

	r.GET("/manga/:uuid", a.GetManga)

	r.POST("/manga/", a.AddManga)

	r.PUT("manga/:uuid", a.ChangeDesc)

	r.DELETE("manga/:uuid", a.DeleteManga)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

type inter struct {
	Status string `json:"status"`
}

// GetList godoc
// @Summary      Get all records
// @Description  Get a list of all mangas
// @Tags         Info
// @Produce      json
// @Success      200  {object}  ds.Manga
// @Failure 500 {object} models.ModelError
// @Router       /manga/ [get]
func (a *Application) GetList(gCtx *gin.Context) {
	resp, err := a.repo.GetAllManga()
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a list",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)
}

// GetManga  godoc
// @Summary      Get manga with corresponding name
// @Description  Get a manga via uuid
// @Tags         Info
// @Produce      json
// @Param UUID query string true "UUID манги"
// @Success      200  {object}  models.ModelMangaDesc
// @Failure 	 500 {object} models.ModelError
// @Router       /manga/{uuid} [get]
func (a *Application) GetManga(gCtx *gin.Context) {
	uuid := gCtx.Param("uuid")
	resp, err := a.repo.GetMangaByName(uuid)
	if err != nil {
		if resp == nil {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Error: "not found",
				})
			return
		}
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a manga",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)
}

// ChangeDesc  godoc
// @Summary      Change manga description
// @Description  Change a description of manga via its uuid
// @Tags         Change
// @Produce      json
// @Param UUID query string true "UUID манги"
// @Param Description query string true "Новое описание"
// @Success      200  {object}  models.ModelDescChanged
// @Failure 	 500 {object} models.ModelError
// @Router       /manga/{uuid} [put]
func (a *Application) ChangeDesc(gCtx *gin.Context) {
	inputUuid, _ := uuid.Parse(gCtx.Param("uuid"))
	newDesc := gCtx.Query("Description")
	err := a.repo.ChangeDescription(inputUuid, newDesc)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		gCtx.JSON(
			http.StatusNotFound,
			&models.ModelError{
				Description: err.Error(),
				Error:       "db error",
				Type:        "internal",
			})
		return
	}

	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: err.Error(),
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelDescChanged{
			Success: true,
		})

}

// DeleteManga   godoc
// @Summary      Delete a manga
// @Description  Delete a manga via its uuid
// @Tags         Change
// @Produce      json
// @Param UUID query string true "UUID манги"
// @Success      200  {object}  models.ModelMangaDeleted
// @Failure 	 500 {object} models.ModelError
// @Router       /manga/{uuid} [delete]
func (a *Application) DeleteManga(gCtx *gin.Context) {
	uuid := gCtx.Param("uuid")
	msg, err := a.repo.DeleteManga(uuid)
	if err != nil {
		if msg == "no manga" {
			gCtx.JSON(
				http.StatusBadRequest,
				&models.ModelError{
					Description: "No product found with this uuid",
					Error:       "uuid error",
					Type:        "client",
				})
			return
		}
		if msg == "no such rows" {
			gCtx.JSON(
				http.StatusBadRequest,
				&models.ModelError{
					Description: msg,
					Error:       "uuid error",
					Type:        "client",
				})
			return
		}
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "delete failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelMangaDeleted{
			Success: true,
		})

}

// AddManga godoc
// @Summary      Add a new manga
// @Description  Adding a new manga to database
// @Tags         Add
// @Produce      json
// @Param Name body string true "Название манги"
// @Param Rate body uint64 true "Рейтинг манги"
// @Param Year body uint64 true "Год производства"
// @Param Genre body string true "Жанр"
// @Param Price body uint64 true "Цена"
// @Param Episodes body uint64 true "Количество серий"
// @Param Description body string false "Описание"
// @Success      201  {object}  models.ModelMangaCreated
// @Failure 500 {object} models.ModelError
// @Router       /manga/ [Post]
func (a *Application) AddManga(gCtx *gin.Context) {
	manga := ds.Manga{}

	if err := gCtx.BindJSON(&manga); err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "unmarshal failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	if manga.Price < 0 || manga.Year < 0 || manga.Rate < 1 {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Write correct data",
				Error:       "Price error",
			})
		return
	}
	manga.UUID = uuid.New()
	err := a.repo.CreateManga(manga)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Bad Request",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelMangaCreated{
			Success: true,
		})
}
