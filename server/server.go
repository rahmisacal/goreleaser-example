package server

import (
	"../auth"
	"../domain"
	"../restapi"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {

	r := mux.NewRouter()

	restapi.Books = append(restapi.Books, domain.Book{ID: "1", Isbn: "448743", Title: "Book one", Author: &domain.Author{Firstname: "joe", Lastname: "doe"}})
	restapi.Books = append(restapi.Books, domain.Book{ID: "2", Isbn: "848743", Title: "Book two", Author: &domain.Author{Firstname: "steven", Lastname: "bremen"}})

	r.HandleFunc("/api/books", auth.BasicAuthMiddleware(restapi.GetBooks)).Methods("GET")
	r.HandleFunc("/api/books/{id}", auth.BasicAuthMiddleware(restapi.GetBook)).Methods("GET")
	r.HandleFunc("/api/books", restapi.CreateBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", restapi.UpdateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", restapi.DeleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1025", r))
}
