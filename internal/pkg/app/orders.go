package app

import (
	"WAD-2022/internal/app/ds"
	"WAD-2022/swagger/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (a *Application) AddOrder(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	userUUID := a.GetUserByToken(jwtStr)
	order := ds.Order{}
	order.UserUUID = userUUID
	err := gCtx.BindJSON(&order)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid parameters",
				Error:       "400",
			})
		return
	}
	err = a.repo.AddOrder(order)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "Create failed",
			})
		return
	}
	gCtx.JSON(http.StatusOK, "Success: true")

}

func (a *Application) GetOrders(gCtx *gin.Context) {
	stDate := gCtx.Query("start_date")
	endDate := gCtx.Query("end_date")
	status := gCtx.Query("status")
	resp, err := a.repo.GetOrders(stDate, endDate, status)

	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&models.ModelError{
				Description: "can`t get a list",
			})
		return
	}
	gCtx.JSON(http.StatusOK, resp)

}

func (a *Application) ChangeStatus(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "Invalid UUID format",
			})
		return
	}
	order := ds.Order{}
	err = gCtx.BindJSON(&order)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&models.ModelError{
				Description: "The price is negative or not int",
			})
		return
	}
	resp, err := a.repo.ChangeStatus(UUID, order.Status)
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

	gCtx.JSON(http.StatusOK, "Success: true")

}
