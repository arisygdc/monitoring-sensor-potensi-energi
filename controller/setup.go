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
			"message": err.Error(),
		})
		return
	}

	delim := " "
	req.Desa = strings.Trim(req.Desa, delim)
	req.Identity = strings.Trim(req.Identity, delim)
	req.Kecamatan = strings.Trim(req.Kecamatan, delim)
	req.NamaLokasi = strings.Trim(req.NamaLokasi, delim)
	req.Provinsi = strings.Trim(req.Provinsi, delim)
	req.TipeSensor = strings.Trim(req.TipeSensor, delim)

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
