package controller

import (
	"monitoring-potensi-energi/reqres"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) InputData(ctx *gin.Context) {
	var req reqres.InputValue
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}
	ctr.Repo.InputValue(ctx, req)
}
