package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"api-test/app"
	"api-test/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("First  Main!")
	r := mux.NewRouter()
	r.Use(app.JwtAuthentication)
	port := os.Getenv("app_port")
	if port == "" {
		port = "8083" //localhost
	}
	fmt.Println("Second  Main!")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")
	api.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
	api.HandleFunc("/students", controllers.GetAllStudents).Methods("GET")
	api.HandleFunc("/students/{doc_num}", controllers.GetStudentByDocument).Methods("GET")
	api.HandleFunc("/students/new", controllers.CreateStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, r))
	fmt.Println("Third  Main!")
}
