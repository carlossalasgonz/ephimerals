package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var envConf map[string]string

func main() {
	LoadEnv()
	router := gin.Default()
	router.GET("/time", func(c *gin.Context) {
		var timezone ?= 
			
		c.JSON(http.StatusOK, gin.H{
			"timezone": envConf["SYSTEM_TIMEZONE"],
			"date": time.Now().Format(time.DateOnly),
			"time": time.Now().Format(time.TimeOnly),
		})
	})
	router.Run()
}

func LoadEnv(){
	var err error
	config := make(map[string]string)

	//Load default values
	for _, entry := range os.ReadFile()
	for _, entry := range os.Environ() {
		key, value, found := strings.Cut(entry, "=")
		if found {
			config[key] = value
		}
	}

	envConf = config
	return err
}