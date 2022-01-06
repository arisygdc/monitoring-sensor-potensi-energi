package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) GetSensors(ctx *gin.Context) {
	sensors, err := ctr.Repo.GetAllSensor(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"sensors": sensors,
	})
}
