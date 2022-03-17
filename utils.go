package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func httpErrorSend(err error, c *gin.Context) {
	log.Println(err)
	c.JSON(http.StatusFailedDependency, gin.H{
		"status": "failed",
	})
}
