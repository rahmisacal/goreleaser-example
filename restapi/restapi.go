package restapi

import (
	"../domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var Books []domain.Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	_ = json.NewEncoder(w).Encode(Books)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range Books{
		if item.ID == params["id"]{
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&domain.Book{})
}

func CreateBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	var book domain.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	//book.ID = strconv.Itoa(rand.Intn(10000000))
	Books = append(Books, book)
	_ = json.NewEncoder(w).Encode(book)
}

func UpdateBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			var book domain.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			Books = append(Books, book)
			_ = json.NewEncoder(w).Encode(book)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(Books)
}

func DeleteBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}
	_ = json.NewEncoder(w).Encode(Books)
}