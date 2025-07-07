package handlers

import (
	"net/http"
	"projectOzinshe/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHendlers struct {
	bd map[int]models.Movie
}

func NewMoviesHendler() *MovieHendlers {
	return &MovieHendlers{
		bd: map[int]models.Movie{
			1: {
				Id:          1,
				Title:       "1+1",
				Description: "Пострадав в результате несчастного случая, богатый аристократ Филипп нанимает в помощники человека, который менее всего подходит для этой работы, – молодого жителя предместья Дрисса, только что освободившегося из тюрьмы. Несмотря на то, что Филипп прикован к инвалидному креслу, Дриссу удается привнести в размеренную жизнь аристократа дух приключений.",
				ReleaseYear: 2011,
				Director:    "Оливье Накаш",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://flcksbr.top/film/535341/",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},
			2: {
				Id:          2,
				Title:       "Интерстеллар",
				Description: "Когда засуха, пыльные бури и вымирание растений приводят человечество к продовольственному кризису, коллектив исследователей и учёных отправляется сквозь червоточину (которая предположительно соединяет области пространства-времени через большое расстояние) в путешествие, чтобы превзойти прежние ограничения для космических путешествий человека и найти планету с подходящими для человечества условиями.",
				ReleaseYear: 2014,
				Director:    "Кристофер Нолан",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=6ybBuTETr3U",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},
			3: {
				Id:          3,
				Title:       "Побег из Шоушенка",
				Description: "Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения.",
				ReleaseYear: 1994,
				Director:    "Фрэнк Дарабонт",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=kgAeKpAPOYk&ab_channel=%D0%A2%D1%80%D0%B5%D0%B9%D0%BB%D0%B5%D1%80%D1%8B%D0%BA%D1%84%D0%B8%D0%BB%D1%8C%D0%BC%D0%B0%D0%BC",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},
			4: {
				Id:          4,
				Title:       "Бойцовский клуб",
				Description: "Сотрудник страховой компании страдает хронической бессонницей и отчаянно пытается вырваться из мучительно скучной жизни. Однажды в очередной командировке он встречает некоего Тайлера Дёрдена — харизматического торговца мылом с извращенной философией. Тайлер уверен, что самосовершенствование — удел слабых, а единственное, ради чего стоит жить, — саморазрушение. Проходит немного времени, и вот уже новые друзья лупят друг друга почем зря на стоянке перед баром, и очищающий мордобой доставляет им высшее блаженство. Приобщая других мужчин к простым радостям физической жестокости, они основывают тайный Бойцовский клуб, который начинает пользоваться невероятной популярностью.",
				ReleaseYear: 1999,
				Director:    "Дэвид Финчер",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=C7-7qQ61QHU&ab_channel=%D0%A2%D1%80%D0%B5%D0%B9%D0%BB%D0%B5%D1%80%D1%8B%D0%BA%D1%84%D0%B8%D0%BB%D1%8C%D0%BC%D0%B0%D0%BC",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},
			5: {
				Id:          5,
				Title:       "Остров проклятых",
				Description: "Два американских судебных пристава отправляются на один из островов в штате Массачусетс, чтобы расследовать исчезновение пациентки клиники для умалишенных преступников. При проведении расследования им придется столкнуться с паутиной лжи, обрушившимся ураганом и смертельным бунтом обитателей клиники.",
				ReleaseYear: 2009,
				Director:    "Мартин Скорсезе",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=_l7R9Rz5URw&ab_channel=%D0%A2%D1%80%D0%B5%D0%B9%D0%BB%D0%B5%D1%80%D1%8B%D0%BA%D1%84%D0%B8%D0%BB%D1%8C%D0%BC%D0%B0%D0%BC",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},
			6: {
				Id:          6,
				Title:       "Начало",
				Description: "Кобб — талантливый вор, лучший из лучших в опасном искусстве извлечения: он крадёт ценные секреты из глубин подсознания во время сна, когда разум наиболее уязвим. Его редкий талант сделал его ценным игроком в мире корпоративного шпионажа, но он также превратил его в беглеца и лишил всего, что он когда-либо любил.",
				ReleaseYear: 2010,
				Director:    "Кристофер Нолан",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=YoHD9XEInc0",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},

			7: {
				Id:          7,
				Title:       "Зелёная миля",
				Description: "Пол Эджкомб — начальник блока смертников в тюрьме. За долгие годы работы он видел много заключённых, но Джон Коффи — особенный. Огромный чернокожий мужчина обвинён в ужасном преступлении, но он обладает сверхъестественным даром и мягким нравом.",
				ReleaseYear: 1999,
				Director:    "Фрэнк Дарабонт",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=Ki4haFrqSrw",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},

			8: {
				Id:          8,
				Title:       "Престиж",
				Description: "История двух молодых и амбициозных фокусников, соперничество которых перерастает в настоящую войну. Каждый из них готов пойти на всё, чтобы раскрыть секреты другого и создать самый впечатляющий трюк.",
				ReleaseYear: 2006,
				Director:    "Кристофер Нолан",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=o4gHCmTQDVI",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},

			9: {
				Id:          9,
				Title:       "Игра в имитацию",
				Description: "Алан Тьюринг, британский математик, логик и криптограф, возглавляет команду, которая пытается взломать код немецкой машины Энигма во время Второй мировой войны. Однако его личная тайна может разрушить всё.",
				ReleaseYear: 2014,
				Director:    "Мортен Тильдум",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=S5CjKEFb-sM",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},

			10: {
				Id:          10,
				Title:       "Семь",
				Description: "Два детектива, один из которых ветеран, а другой новичок, расследуют серию изощрённых убийств, связанных с семью смертными грехами. Чем ближе они к разгадке, тем страшнее становится правда.",
				ReleaseYear: 1995,
				Director:    "Дэвид Финчер",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=znmZoVkCjpI",
				PosterUrl:   "",
				Genres:      make([]models.Genre, 0),
			},
		},
	}
}

func (h *MovieHendlers) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid ID format"))
		return
	}

	movie, exists := h.bd[id]
	if !exists {
		c.JSON(http.StatusNotFound, models.NewApiError("Movie not found"))
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

	_, exists := h.bd[id]
	if !exists {
		c.JSON(http.StatusNotFound, models.NewApiError("Movie not found"))
		return
	}

	delete(h.bd, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Movie deleted successfully",
	})
}

func (h *MovieHendlers) FindAll(c *gin.Context) {
	movies := make([]models.Movie, 0, len(h.bd))

	for _, v := range h.bd {
		movies = append(movies, v)
	}
	c.JSON(http.StatusOK, movies)
}
func (h *MovieHendlers) Create(c *gin.Context) {
	var n models.Movie
	err := c.BindJSON(&n)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("not worked"))
		return
	}

	id := len(h.bd) + 1

	n.Id = id
	n.Genres = make([]models.Genre, 0)

	h.bd[id] = n

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

	originalMovie, ok := h.bd[id]
	if !ok {
		c.JSON(http.StatusBadRequest, models.NewApiError("not found the cinema"))
		return
	}

	var updatedMuvie models.Movie
	err = c.BindJSON(&updatedMuvie)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("not worked"))
		return
	}

	originalMovie.Title = updatedMuvie.Title
	originalMovie.Description = updatedMuvie.Description
	originalMovie.ReleaseYear = updatedMuvie.ReleaseYear
	originalMovie.Director = updatedMuvie.Director
	originalMovie.Rating = updatedMuvie.Rating
	originalMovie.IsWatched = updatedMuvie.IsWatched
	originalMovie.TrailerUrl = updatedMuvie.TrailerUrl
	originalMovie.WatchUrl = updatedMuvie.WatchUrl
	originalMovie.PosterUrl = updatedMuvie.PosterUrl
	originalMovie.Genres = updatedMuvie.Genres

	h.bd[id] = originalMovie

	c.Status(http.StatusOK)

}
