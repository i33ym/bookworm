package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"api.bookworm.cc/internal/format"
	"github.com/julienschmidt/httprouter"
)

type Book struct {
	ID        int       `json:"book_id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Authors   []string  `json:"authors"`
	Genres    []string  `json:"genres"`
	Pages     int       `json:"pages"`
	Rating    float64   `json:"rating"`
	Publisher string    `json:"press"`
	Year      int       `json:"year_of_publish,omitempty"`
}

func (handlers *Handlers) CreateBook(response http.ResponseWriter, request *http.Request) {
	type input struct {
		Title     string   `json:"title"`
		Authors   []string `json:"authors"`
		Genres    []string `json:"genres"`
		Pages     int      `json:"pages"`
		Publisher string   `json:"press"`
		Year      int      `json:"year_of_publish"`
	}

	temp := &input{}

	if err := format.Read(request, temp); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Bad Request\n"))
		return
	}

	book := &Book{
		ID:        123,
		CreatedAt: time.Now(),
		Title:     temp.Title,
		Authors:   temp.Authors,
		Genres:    temp.Genres,
		Pages:     temp.Pages,
		Publisher: temp.Publisher,
		Year:      temp.Year,
	}

	headers := make(http.Header)
	headers.Add("Location", fmt.Sprintf("/v1/books/%d", book.ID))

	if err := format.Respond(response, http.StatusCreated, book, headers); err != nil {
		handlers.logger.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Internal Server Error\n"))

		return
	}

	// curl -i -X POST -d '{"title": "Hello World", "authors": ["Isroil Muhitdinov", "Feruz Durmamatov"], "genres": ["Computer Science", "Literature"], "pages": 540, "press": "INHA University in Tashkent", "year_of_publish": 2018}' localhost:5000/v1/books
}

func (handlers *Handlers) ViewBook(response http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("The requested resource could not be found\n"))

		return
	}

	book := &Book{
		ID:        id,
		Title:     "The Structure and Interpretation of Computer Programs",
		Authors:   []string{"Hall Abelson", "Jay Sussman"},
		Genres:    []string{"Computer Science", "Programming", "Non-Fiction"},
		Pages:     864,
		Rating:    4.8,
		Publisher: "The MIT Press",
		Year:      1995,
	}

	if err := format.Respond(response, http.StatusOK, book, nil); err != nil {
		handlers.logger.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Internal Server Error\n"))

		return
	}
}
