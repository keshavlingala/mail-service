package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/keshavlingala/mail-service/controllers"
	"log"
	"os"
	"runtime"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

func StartWorkers() {
	go statsWorker()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://keshav.codes", "https://swasthomeo.com", "https://swasthahomeopathy.web.app"}
	router.Use(cors.New(config))

	router.Use(rateLimit, gin.Recovery())
	router.GET("/", controllers.IndexPage)
	router.GET("/google_login", controllers.GoogleLogin)
	router.GET("/google_callback", controllers.GoogleCallback)
	router.POST("/keshav", controllers.KeshavEmail)
	router.POST("/swastha", controllers.SwasthaContactUs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
