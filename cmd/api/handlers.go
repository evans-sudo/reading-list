package main

import (
	"encoding/json"
	"evansgopher/internal/data"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
    

  data := map[string]string {
		"status": "available",
		"environment": app.config.env,
		"version": version,
  }

  js, err := json.Marshal(data)
  if err != nil {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
  }

  js = append(js, '\n')

  w.Header().Set("content-Type", "application/json")

  w.Write(js)

}



func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		books := []data.Book {
			{
				ID: 1,
				CreatedAt: time.Now(),
				Title: "The Darkening of Tristram",
				Published: 1998,
				Pages: 300,
				Genres: []string{"Fiction", "Thriller"},
				Rating: 4.5,
				Version: 1,
			},

			{
				ID: 2,
				CreatedAt: time.Now(),
				Published: 2007,
				Pages: 432,
				Genres: []string{"Fiction", "Adventure"},
				Rating: 4.9,
				Version: 1,
			},
		}
		if err := app.WriteJSON(w, http.StatusOK, envelope{"books":books}); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		
	}

	if r.Method == http.MethodPost {
		var input struct {
			Title string `json:"title"`
			Published int `json:"published"`
			Pages int `json:"pages`
			Genres []string `json:"genres"`
			Rating float64 `json:"rating"`
		} 
 

		err := app.readJSON(w,r, &input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%v\n", input)
	}
}


func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request)  {

	switch r.Method {
	case http.MethodGet: 
		app.getBook(w,r)
	case http.MethodPut:
		app.updateBook(w,r)
	case http.MethodDelete:
		app.deleteBook(w,r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 3)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} 

	book := data.Book {
		ID: idInt,
		CreatedAt: time.Now(),
		Title: "Echoes in the darkness",
		Published: 2019,
		Pages: 300,
		Genres: []string{"Fiction", "Thriller"},
		Rating: 4.5,
		Version: 1,
	}

	

	if err := app.WriteJSON(w, http.StatusOK, envelope{"book":book}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Display the details of book with ID:%d", idInt)
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 12, 1)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
     
	var input struct {
		Title *string `json:"title"`
		Published *int `json:"published"`
		Pages *int `json:"pages"`
		Genres []string `json:"genres"`
		Rating *float32 `json:"rating"`
	} 


	book := data.Book{
		ID: idInt,
		CreatedAt:time.Now(),
		Title: "Echoes in the Darkness",
		Published: 2019,
		Pages: 300,
		Genres: []string{"Fiction", "Thriller"},
		Rating: 4.5,
		Version: 1,
	}



	err = app.readJSON(w,r,&input)
		if err != nil {
			http.Error(w,http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}


		if input.Title != nil {
			book.Title = *input.Title
		}
	

		if input.Published != nil {
			book.Published = *input.Published
		}           

		if input.Pages != nil {
			book.Pages = *input.Pages 
		}

		if len(input.Genres ) > 0 {
			book.Genres = input.Genres
		}

		if input.Rating != nil {
			book.Rating = *input.Rating
		}

		fmt.Fprintf(w, "%v\n", book)


}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id:= r.URL.Path[len("/v1/books/"):]
	idint, err := strconv.ParseInt(id, 10, 4)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "delete the details of the book with id %d", idint)
}