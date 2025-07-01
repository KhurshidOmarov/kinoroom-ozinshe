package models

type Movie struct {
	Id          int
	Title       string
	Description string
	ReleaseYear int
	Director    string
	Rating      int
	IsWatched   bool
	TrailerUrl  string
	WatchUrl    string
	PosterUrl   string  // not imlemented
	Genre       []Genre // not imlemented
}
