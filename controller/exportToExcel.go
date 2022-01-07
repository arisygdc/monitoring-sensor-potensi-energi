package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func (ctr Controller) ExportToexcel(ctx *gin.Context) {
	id := ctx.Param("id_sensor")

	idSensor, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filename := fmt.Sprintf("Export sensor id %v-%v.xlsx", id, time.Now().String())
	dir, err := os.Getwd()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	file := fmt.Sprintf("%v/exportfile/%v", dir, filename)
	f, err := excelize.OpenFile(file)
	if err != nil {
		f = excelize.NewFile()
	}

	values, err := ctr.Repo.GetAllValueSensor(ctx, int32(idSensor))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sheet := "Sheet1"
	for i, v := range values {
		row := i + 1
		axis := fmt.Sprintf("A%d", row)
		f.SetCellValue(sheet, axis, row)
		axis = fmt.Sprintf("B%d", row)
		f.SetCellValue(sheet, axis, v.Data)
		axis = fmt.Sprintf("C%d", row)
		f.SetCellValue(sheet, axis, v.DibuatPada.String())
	}

	err = f.SaveAs(file)
	if err != nil {
		if err := f.Save(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

}
