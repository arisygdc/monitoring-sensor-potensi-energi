package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) Monitoring(ctx *gin.Context) {
	var (
		param    = ctx.Param("id_sensor")
		idSensor int64
	)

	idSensor, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	metrics, err := ctr.Repo.GetMetrics(ctx, int32(idSensor))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"metrics": metrics,
	})
}
