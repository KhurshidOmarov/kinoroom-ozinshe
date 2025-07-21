package main

import (
	"projectOzinshe/handlers"
	"projectOzinshe/repositories"

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

	moviesRepository := repositories.NewMoviesRepository()
	genresRepositories := repositories.NewGenreRepositories()
	movieHendlers := handlers.NewMoviesHendler(moviesRepository, genresRepositories)

	r.POST("/movies", movieHendlers.Create)
	r.PUT("/movies/:id", movieHendlers.Update)
	r.GET("/movies", movieHendlers.FindAll)
	r.GET("/movies/:id", movieHendlers.FindById)

	r.DELETE("/movies/:id", movieHendlers.Delete)

	// genres

	genresHandlers := handlers.NewGenresHandlers(genresRepositories)

	r.GET("/genres", genresHandlers.FindAll)
	r.GET("/genres/:id", genresHandlers.FindById)
	r.POST("/genres", genresHandlers.Create)
	r.PUT("/genres/:id", genresHandlers.Update)
	r.DELETE("/genres/:id", genresHandlers.Delete)

	r.Run(":8083")
}
