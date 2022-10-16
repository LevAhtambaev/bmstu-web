package app

import (
	"WAD-2022/internal/app/ds"
	"WAD-2022/swagger/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
)

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/manga/all", a.GetList)

	r.GET("/manga/find", a.GetManga)

	r.POST("/manga/create", a.AddManga)

	r.PUT("manga/changeDescription", a.ChangeDesc)

	r.DELETE("manga/delete", a.DeleteManga)

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
// @Router       /manga/all [get]
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
// @Description  Get a manga via name
// @Tags         Info
// @Produce      json
// @Param Name query string true "Название манги"
// @Success      200  {object}  models.ModelMangaDesc
// @Failure 	 500 {object} models.ModelError
// @Router       /manga/find [get]
func (a *Application) GetManga(gCtx *gin.Context) {
	name := gCtx.Query("Name")
	resp, err := a.repo.GetMangaByName(name)
	if err != nil {
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
// @Param String query int true "Новое описание"
// @Success      200  {object}  models.ModelDescChanged
// @Failure 	 500 {object} models.ModelError
// @Router       /manga/changeDescription [put]
func (a *Application) ChangeDesc(gCtx *gin.Context) {
	inputUuid, _ := uuid.Parse(gCtx.Query("UUID"))
	newDesc := gCtx.Query("Description")
	err := a.repo.ChangeDescription(inputUuid, newDesc)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "update failed",
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
// @Router       /manga/delete [delete]
func (a *Application) DeleteManga(gCtx *gin.Context) {
	uuid := gCtx.Query("UUID")
	err := a.repo.DeleteManga(uuid)
	if err != nil {
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
// @Param Name query string true "Название манги"
// @Param Rate query uint64 true "Рейтинг манги"
// @Param Year query uint64 true "Год производства"
// @Param Genre query string true "Жанр"
// @Param Price query uint64 true "Цена"
// @Param Episodes query uint64 true "Количество серий"
// @Param Description query string false "Описание"
// @Success      201  {object}  models.ModelMangaCreated
// @Failure 500 {object} models.ModelError
// @Router       /manga/create [Post]
func (a *Application) AddManga(gCtx *gin.Context) {
	rate, _ := strconv.ParseUint(gCtx.Query("Rate"), 10, 64)
	year, _ := strconv.ParseUint(gCtx.Query("Year"), 10, 64)
	price, _ := strconv.ParseUint(gCtx.Query("Price"), 10, 64)
	ep, _ := strconv.ParseUint(gCtx.Query("Episodes"), 10, 64)

	manga := ds.Manga{
		Name:        gCtx.Query("Name"),
		Rate:        rate,
		Year:        year,
		Genre:       gCtx.Query("Genre"),
		Price:       price,
		Episodes:    ep,
		Description: gCtx.Query("Description"),
	}

	err := a.repo.CreateManga(manga)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "adding failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelMangaCreated{
			Success: true,
		})

}
