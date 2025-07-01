package main

import (
	"projectOzinshe/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
	}
	r.Use(cors.New(corsConfig))

	movieHendlers := handlers.NewMoviesHendler()

	r.POST("/movie", movieHendlers.Create)
	r.PUT("/movie/:id", movieHendlers.Update)
	r.GET("/movie", movieHendlers.FindAll)

	r.Run("0.0.0.0:8082")
}
