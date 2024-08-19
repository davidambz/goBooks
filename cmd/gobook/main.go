package main

import (
	"database/sql"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)
	bookHandlers := web.NewBookHandlers(bookService)

	router := mux.NewRouter()
	router.HandleFunc("/books", bookHandlers.GetBooks).Methods("GET")
	router.HandleFunc("/books", bookHandlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", bookHandlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", bookHandlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", bookHandlers.DeleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
