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
// @Description  Get a list of all comics
// @Tags         Info
// @Produce      json
// @Success      200  {object}  ds.Comics
// @Failure 500 {object} models.ModelError
// @Router       /comics/ [get]
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

// GetComics  godoc
// @Summary      Get comics with corresponding name
// @Description  Get a comics via uuid
// @Tags         Info
// @Produce      json
// @Param UUID query string true "UUID комикса"
// @Success      200  {object}  models.ModelComicsDesc
// @Failure 	 500 {object} models.ModelError
// @Router       /comics/{uuid} [get]
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

// ChangeComics  godoc
// @Summary      Change comics
// @Description  Change a description of comics via its uuid
// @Tags         Change
// @Produce      json
// @Param UUID query string true "UUID комикса"
// @Param Description query string true "Новое описание"
// @Success      200  {object}  models.ModelDescChanged
// @Failure 	 500 {object} models.ModelError
// @Router       /comics/{uuid} [put]
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

// DeleteComics   godoc
// @Summary      Delete a comics
// @Description  Delete a comics via its uuid
// @Tags         Change
// @Produce      json
// @Param UUID query string true "UUID комикса"
// @Success      200  {object}  models.ModelComicsDeleted
// @Failure 	 500 {object} models.ModelError
// @Router       /comics/{uuid} [delete]
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
		&models.ModelComicsDeleted{
			Success: true,
		})

}

// AddComics godoc
// @Summary      Add a new comics
// @Description  Adding a new comics to database
// @Tags         Add
// @Produce      json
// @Param Name body string true "Название"
// @Param Rate body uint64 true "Рейтинг"
// @Param Year body uint64 true "Год производства"
// @Param Genre body string true "Жанр"
// @Param Price body uint64 true "Цена"
// @Param Episodes body uint64 true "Количество серий"
// @Param Description body string false "Описание"
// @Success      201  {object}  models.ModelComicsCreated
// @Failure 500 {object} models.ModelError
// @Router       /comics/ [Post]
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
		&models.ModelComicsCreated{
			Success: true,
		})
}
