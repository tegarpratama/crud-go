package main

import (
	"crud-go/config"
	"crud-go/dto"
	"crud-go/helper"
	"crud-go/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/check-health", func(ctx *gin.Context) {
		res := helper.Response(dto.ResponseParams{
			StatusCode: http.StatusOK,
			Message: "It's work",
		})

		ctx.JSON(http.StatusOK, res)
	})

	router.PostRouter(api)
	
	r.Run(fmt.Sprintf("127.0.0.1:%v", config.ENV.PORT))
}