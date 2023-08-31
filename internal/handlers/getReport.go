package handlers

import (
	"avitoTech/internal/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetReportResponse struct {
	DownloadLink string `json:"download_link" validate:"required,url"`
}

func GetReport(ctx *gin.Context) {
	month := ctx.Param("month")
	year := ctx.Param("year")
	monthInt, errMonth := strconv.Atoi(month)
	if errMonth != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "month error " + errMonth.Error()})
		return
	}
	yearInt, errYear := strconv.Atoi(year)
	if errYear != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "year error " + errYear.Error()})
		return
	}
	fileName := fmt.Sprintf("report_%d_%d.csv", yearInt, monthInt)
	// Создать ссылку на скачивание файла
	downloadLink := "/download/" + fileName

	ctx.JSON(http.StatusOK, &GetReportResponse{
		DownloadLink: downloadLink,
	})
}

func FileDownload(ctx *gin.Context) {

	filename := ctx.Param("filename")
	parts := strings.Split(filename, "_")
	year := parts[1]
	month := strings.TrimSuffix(parts[2], ".csv")
	reqCtx := ctx.Request.Context()
	service := services.Must(reqCtx).UserSegmentService
	fileData, err := service.GetReport(ctx, month, year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	// заголовки ответа для скачивания файла
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Type", "text/csv")

	// данные файла как тело ответа
	ctx.Data(http.StatusOK, "text/csv", fileData)
}
