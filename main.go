package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", students.createStudent.Methods(http.MethodPost)
	api.HandleFunc("", students.getAllStudents).Methods(http.MethodGet)
	api.HandleFunc("", students.updateStudent).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8083", r))
}
