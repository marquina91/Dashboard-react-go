package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Post struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	PublishedAt string `json:"publishedAt"`
}

type Users struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/posts", AllPost)
	r.Get("/users", AllUsers)
	http.ListenAndServe(":5000", r)
}

func AllPost(w http.ResponseWriter, r *http.Request) {
	post := [...]Post{
		{
			Id:          "1",
			Title:       "Post One",
			Body:        "This is the Body 1",
			PublishedAt: "2020-12-25",
		},
		{
			Id:          "2",
			Title:       "Post Two",
			Body:        "This is the Body 2",
			PublishedAt: "2020-12-25",
		},
		{
			Id:          "3",
			Title:       "Post Three",
			Body:        "This is the Body 3",
			PublishedAt: "2020-12-25",
		},
		{
			Title:       "Post Four",
			Body:        "This is the Body 4",
			PublishedAt: "12-25-2020",
			Id:          "DdCICG7",
		},
	}
	respondWithJSON(w, http.StatusOK, post)
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	user := [...]Users{
		{
			Id:    "1",
			Name:  "Pablito cerveza",
			Email: "Pablitocerv@gmail.com",
		},
		{
			Id:    "2",
			Name:  "Manuel Mazias",
			Email: "manuelpajares@gmail.com",
		},
		{
			Id:    "3",
			Name:  "Tom Pajares",
			Email: "tomcpaje@gmail.com",
		},
	}
	respondWithJSON(w, http.StatusOK, user)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	//***cors***
	//w.Header().Set("Access-Control-Allow-Origin", "localhost:3000")
	//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//w.Header().Set("Access-Control-Expose-Headers", "Content-Range")
	w.Header().Set("Content-Range", "*")
	//w.Header().Set("X-Total-Count", "30")
	//w.Header().Set("Access-Control-Expose-Headers", "Content-Range, X-Total-Count")
	w.WriteHeader(code)
	w.Write(response)
}
