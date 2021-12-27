package controller

import (
	"monitoring-potensi-energi/reqres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) Setup(ctx *gin.Context) {
	var req reqres.SetupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := ctr.Repo.PlaceSensor(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "accept",
	})
}
