package app

import (
	"WAD-2022/internal/app/ds"
	"WAD-2022/internal/app/role"
	"WAD-2022/swagger/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

const jwtPrefix = "Bearer "

func (a *Application) WithAuthCheck(assignedRoles ...role.Role) func(ctx *gin.Context) {
	return func(gCtx *gin.Context) {
		jwtStr := gCtx.GetHeader("Authorization")
		log.Println(jwtStr)
		if !strings.HasPrefix(jwtStr, jwtPrefix) { // если нет префикса то нас дурят!
			gCtx.AbortWithStatus(http.StatusForbidden) // отдаем что нет доступа

			return // завершаем обработку
		}

		// отрезаем префикс
		jwtStr = jwtStr[len(jwtPrefix):]
		// проверяем jwt в блеклист редиса
		err := a.redis.CheckJWTInBlacklist(gCtx.Request.Context(), jwtStr)
		if err == nil { // значит что токен в блеклисте
			gCtx.AbortWithStatus(http.StatusForbidden)

			return
		}
		if !errors.Is(err, redis.Nil) { // значит что это не ошибка отсуствия - внутренняя ошибка
			gCtx.AbortWithError(http.StatusInternalServerError, err)

			return
		}

		token, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(a.config.JWT.Token), nil
		})
		if err != nil {
			gCtx.AbortWithStatus(http.StatusForbidden)
			log.Println(err)

			return
		}

		myClaims := token.Claims.(*ds.JWTClaims)
		log.Println(myClaims)
		for _, oneOfAssignedRole := range assignedRoles {
			if myClaims.Role == oneOfAssignedRole {
				gCtx.Next()
				return
			}
		}
		gCtx.AbortWithStatus(http.StatusForbidden)
		log.Printf("role %s is not assigned in %s", myClaims.Role, assignedRoles)

		return

	}

}

func (a *Application) GetUserByToken(jwtStr string) (userUUID uuid.UUID) {
	if !strings.HasPrefix(jwtStr, jwtPrefix) { // если нет префикса то нас дурят!
		return // завершаем обработку
	}
	// отрезаем префикс
	jwtStr = jwtStr[len(jwtPrefix):]

	token, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.JWT.Token), nil
	})
	if err != nil {
		log.Println(err)

		return
	}

	myClaims := token.Claims.(*ds.JWTClaims)
	log.Println(myClaims)

	return myClaims.UserUUID
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
	resp, err := a.repo.GetAllComics()
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
func (a *Application) GetComics(gCtx *gin.Context) {
	uuid := gCtx.Param("uuid")
	resp, err := a.repo.GetComicsByName(uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		gCtx.JSON(
			http.StatusNotFound,
			&models.ModelError{
				Error: "not found",
			})
		return
	}
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a comics",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)
}

// ChangeManga  godoc
// @Summary      Change manga
// @Description  Change a description of manga via its uuid
// @Tags         Change
// @Produce      json
// @Param UUID query string true "UUID манги"
// @Param Description query string true "Новое описание"
// @Success      200  {object}  models.ModelDescChanged
// @Failure 	 500 {object} models.ModelError
// @Router       /manga/{uuid} [put]
func (a *Application) ChangeComics(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))

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
	comics := ds.Comics{}
	err = gCtx.BindJSON(&comics)
	log.Println(comics)
	resp, err := a.repo.ChangeComics(UUID, comics)
	if err != nil {
		if resp == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Change failed",
				})
			return
		}
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
func (a *Application) DeleteComics(gCtx *gin.Context) {
	uuid := gCtx.Param("uuid")
	msg, err := a.repo.DeleteComics(uuid)
	if err != nil {
		if msg == "no comics" {
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
func (a *Application) AddComics(gCtx *gin.Context) {
	comics := ds.Comics{}
	err := gCtx.BindJSON(&comics)
	log.Println(comics)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "unmarshal failed",
				Error:       "db error",
				Type:        "internal",
			})
		return
	}
	err = a.repo.CreateComics(comics)
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
