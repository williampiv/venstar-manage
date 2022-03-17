package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	api "github.com/williampiv/venstar-manage/internal/api"
)

var router *gin.Engine
var thermostatIP string

func main() {
	log.Println("Venstar Manage")
	router = gin.Default()
	thermostatIP = os.Getenv("THERMOSTAT_IP")
	router.GET("/", showIndexPage)

	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}
}

func showIndexPage(c *gin.Context) {
	vStatus, err := api.GetThermostatInfo(thermostatIP)
	if err != nil {
		httpErrorSend(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"currentTemp": vStatus.SpaceTemp,
		"coolTemp":    vStatus.CoolTemp,
	})
}
