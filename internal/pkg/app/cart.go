package app

import (
	"WAD-2022/internal/app/ds"
	"WAD-2022/swagger/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (a *Application) AddToCart(gCtx *gin.Context) {
	cart := ds.Cart{}
	err := gCtx.BindJSON(&cart)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid parameters",
				Error:       "Bad request",
			})
		return
	}
	err = a.repo.AddToCart(cart)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "Create failed",
				Error:       "Internal",
			})
		return
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelCartCreated{
			Success: true,
		})

}

func (a *Application) DeleteFromCart(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid UUID format",
				Error:       "Bad request",
			})
		return
	}
	resp, err := a.repo.DeleteFromCart(UUID)
	if err != nil {
		if resp == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&models.ModelError{
					Description: "UUID Not Found",
					Error:       "Bad request",
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&models.ModelError{
					Description: "Delete failed",
					Error:       "Internal",
				})
			return
		}
	}
	gCtx.JSON(
		http.StatusOK,
		&models.ModelCartDeleted{
			Success: true,
		})

}

func (a *Application) GetCart(gCtx *gin.Context) {
	resp, err := a.repo.GetCart()
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a list",
				Error:       "Internal",
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)

}
