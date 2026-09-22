package main

import (
	"net/http"
	"os"
	"time"

	_ "time/tzdata"

	"github.com/gin-gonic/gin"

	"ephimeral/config"
)

var envConf map[string]string

func main() {
	var err error
	envConf, _ = config.LoadEnv("./.env.example", "./.env")
	if _, ok := envConf["SYSTEM_TIMEZONE"]; !ok {
		println("Missing timezone error")
		os.Exit(1)
	}
	_, err = time.LoadLocation(envConf["SYSTEM_TIMEZONE"])
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	router := gin.Default()
	router.GET("/time", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"timezone": envConf["SYSTEM_TIMEZONE"],
			"date":     time.Now().Format(time.DateOnly),
			"time":     time.Now().Format(time.TimeOnly),
		})
	})
	router.Run()
}
