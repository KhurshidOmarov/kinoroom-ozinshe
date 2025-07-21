package repositories

import (
	"context"
	"errors"
	"projectOzinshe/models"
)

type GenreRepositories struct {
	bd map[int]models.Genre
}

func NewGenreRepositories() *GenreRepositories {
	return &GenreRepositories{
		bd: make(map[int]models.Genre),
	}
}

func (r *GenreRepositories) FindByAllIds(ctx context.Context, ids []int) []models.Genre {
	genres := make([]models.Genre, 0, len(r.bd))
	for _, v := range r.bd {
		for _, id := range ids {
			if v.Id == id {
				genres = append(genres, v)
			}
		}
	}
	return genres
}

func (r *GenreRepositories) FindById(ctx context.Context, id int) (models.Genre, error) {
	genre, ok := r.bd[id]
	if !ok {
		return models.Genre{}, errors.New("genre not found")
	}
	return genre, nil
}

func (r *GenreRepositories) FindAll(ctx context.Context) []models.Genre {
	genres := make([]models.Genre, 0, len(r.bd))
	for _, v := range r.bd {
		genres = append(genres, v)
	}
	return genres
}

func (r *GenreRepositories) Create(ctx context.Context, genre models.Genre) int {
	id := len(r.bd) + 1
	genre.Id = id
	r.bd[id] = genre
	return id
}

func (r *GenreRepositories) Update(ctx context.Context, id int, updatedGenre models.Genre) error {
	_, ok := r.bd[id]
	if !ok {
		return errors.New("genre not found")
	}
	updatedGenre.Id = id
	r.bd[id] = updatedGenre
	return nil
}

func (r *GenreRepositories) Delete(ctx context.Context, id int) {
	delete(r.bd, id)
}
