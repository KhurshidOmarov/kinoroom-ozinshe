package handlers

import (
	"context"
	"net/http"
	"projectOzinshe/models"
	"projectOzinshe/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHendlers struct {
	moviesRepo *repositories.MoviesRepository
	genreRepo  *repositories.GenreRepositories
}

type createMovieRequest struct {
	Title       string
	Description string
	ReleaseYear int
	Director    string
	TrailerUrl  string
	GenreIds    []int
}

type updatedMuvieRequest struct {
	Title       string
	Description string
	ReleaseYear int
	Director    string
	TrailerUrl  string
	GenreIds    []int
}

func NewMoviesHendler(
	moviesRepo *repositories.MoviesRepository,
	genreRepo *repositories.GenreRepositories) *MovieHendlers {

	return &MovieHendlers{
		moviesRepo: moviesRepo,
		genreRepo:  genreRepo,
	}
}

func (h *MovieHendlers) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid ID format"))
		return
	}

	movie, err := h.moviesRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHendlers) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid ID format"))
		return
	}

	_, err = h.moviesRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError("Movie not found"))
		return
	}

	h.moviesRepo.Delete(c, id)

	c.Status(http.StatusOK)
}

func (h *MovieHendlers) FindAll(c *gin.Context) {
	movies := h.moviesRepo.FindAll(c)
	c.JSON(http.StatusOK, movies)
}

func (h *MovieHendlers) Create(c *gin.Context) {
	var request createMovieRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("not worked"))
		return
	}

	genres := h.genreRepo.FindByAllIds(context.Background(), request.GenreIds)

	movie := models.Movie{
		Title:       request.Title,
		Description: request.Description,
		ReleaseYear: request.ReleaseYear,
		Director:    request.Director,
		TrailerUrl:  request.TrailerUrl,
		Genres:      genres,
	}

	id := h.moviesRepo.Create(c, movie)

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}

func (h *MovieHendlers) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("progrram not working !"))
		return
	}

	_, err = h.moviesRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError(err.Error()))
		return
	}

	var request updatedMuvieRequest
	err = c.BindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("not worked"))
		return
	}

	genres := h.genreRepo.FindByAllIds(c.Request.Context(), request.GenreIds)
	movie := models.Movie{
		Title:       request.Title,
		ReleaseYear: request.ReleaseYear,
		Director:    request.Director,
		TrailerUrl:  c.Request.RequestURI,
		Genres:      genres,
	}
	h.moviesRepo.Update(c, id, movie)
	c.Status(http.StatusOK)

}
