package controller

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func (ctr Controller) ExportToexcel(ctx *gin.Context) {
	mulai := ctx.Query("mulai")
	sampai := ctx.Query("sampai")

	tMulai, tSampai, err := timeBetweenParse(mulai, sampai)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filename := fmt.Sprintf("%v-%v.xlsx", mulai, sampai)
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

	values, err := ctr.Repo.GetMonitoringDataBetween(ctx, tMulai, tSampai)
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
		f.SetCellValue(sheet, axis, v.ID)
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

func timeBetweenParse(mulai string, sampai string) (tMulai time.Time, tSampai time.Time, err error) {
	tMulai, err = time.Parse("2006-1-2", strings.Trim(mulai, " "))
	if err != nil {
		return
	}
	tSampai, err = time.Parse("2006-1-2", strings.Trim(sampai, " "))
	return
}
