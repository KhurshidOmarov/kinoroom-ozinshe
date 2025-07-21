package handlers

import (
	"net/http"
	"projectOzinshe/models"
	"projectOzinshe/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GenresHandlers struct {
	genreRepo *repositories.GenreRepositories
}

func NewGenresHandlers(genreRepo *repositories.GenreRepositories) *GenresHandlers {
	return &GenresHandlers{genreRepo: genreRepo}
}

func (h *GenresHandlers) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid ID"))
		return
	}

	genre, err := h.genreRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, genre)
}

func (h *GenresHandlers) FindAll(c *gin.Context) {
	genres := h.genreRepo.FindAll(c)
	c.JSON(http.StatusOK, genres)
}

func (h *GenresHandlers) Create(c *gin.Context) {
	var genre models.Genre
	if err := c.BindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid data"))
		return
	}
	id := h.genreRepo.Create(c, genre)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *GenresHandlers) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid ID"))
		return
	}

	var genre models.Genre
	if err := c.BindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid data"))
		return
	}

	err = h.genreRepo.Update(c, id, genre)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewApiError(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *GenresHandlers) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid ID"))
		return
	}

	h.genreRepo.Delete(c, id)
	c.Status(http.StatusOK)
}
