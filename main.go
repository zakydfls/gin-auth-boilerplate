package main

import (
	"gin-auth-boilerplate/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to go server")
	})

	server := &http.Server{
		Addr:    ":3005",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}
