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

	r.POST("/movies", movieHendlers.Create)
	r.PUT("/movies/:id", movieHendlers.Update)
	r.GET("/movies", movieHendlers.FindAll)
	r.GET("/movies/:id", movieHendlers.FindByID)
	r.DELETE("/movies/:id", movieHendlers.Delete)

	r.Run(":8083")
}
