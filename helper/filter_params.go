package helper

import (
	"crud-go/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FilterParams(ctx *gin.Context) *dto.FilterParams {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "5")
	search := ctx.Query("search")

	pageNumber, _ := strconv.Atoi(page)
	limitNumber, _ := strconv.Atoi(limit)
	offset := (pageNumber - 1) * limitNumber

	return &dto.FilterParams{
		Page:   pageNumber,
		Limit:  limitNumber,
		Offset: offset,
		Search: search,
	}
}
