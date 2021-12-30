package controller

import (
	"monitoring-potensi-energi/reqres"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) Setup(ctx *gin.Context) {
	var req reqres.SetupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": req,
		})
		return
	}

	delim := " "
	for i, v := range req.Sensors {
		req.Sensors[i] = strings.Trim(v, delim)
	}
	req.Location.Desa = strings.Trim(req.Location.Desa, delim)
	req.Location.Kecamatan = strings.Trim(req.Location.Kecamatan, delim)
	req.Location.Provinsi = strings.Trim(req.Location.Provinsi, delim)

	idSensor, err := ctr.Repo.PlaceSensor(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message":   "accept",
		"id_sensor": idSensor,
	})
}
