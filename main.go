package main

import (
	"fmt"
	"net/http"
	"os"

	"api-test/app"
	"api-test/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(app.JwtAuthentication)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083" //localhost
	}
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")
	api.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
	api.HandleFunc("/students", controllers.GetAllStudents).Methods("GET")
	api.HandleFunc("/students/{doc_num}", controllers.GetStudentByDocument).Methods("GET")
	api.HandleFunc("/students/new", controllers.CreateStudent).Methods("POST")
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Print(err)
	}
}
