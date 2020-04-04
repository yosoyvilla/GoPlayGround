package students

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Grade string `json:"grade"`
}

type allStudents []student

var students = allStudents{}

func createStudent(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql",
		"rector:admin123@tcp(localhost:8889)/colegio")
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		var newStudent student
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "The student structure is and id, name and grade!"}`))
		}
		json.Unmarshal(reqBody, &newStudent)
		q := fmt.Sprintf("INSERT INTO estudiantes(name, email, grade) values ('%s', '%s', '%s')", newStudent.Name, newStudent.Email, newStudent.Grade)
		fmt.Println(q)
		if _, err := db.Exec(q); err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"message": "There was an issue trying to insert the new student!"}`))
		}
	}
	defer db.Close()
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	studentId := 0
	if s, err := strconv.Atoi(mux.Vars(r)["id"]); err == nil {
		studentId = s
	}
	var updatedStudent student

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(`{"message": "The student structure is and id, name, email and grade!"}`))
	}
	json.Unmarshal(reqBody, &updatedStudent)

	for i, sts := range students {
		if sts.Id == studentId {
			sts.Name = updatedStudent.Name
			sts.Email = updatedStudent.Email
			sts.Grade = updatedStudent.Grade
			students = append(students[:i], sts)
			json.NewEncoder(w).Encode(sts)
		}
	}
}

func checkIfExists(sts student) bool {
	stsExists := false
	for _, v := range students {
		stsExists = v.Id == sts.Id && v.Email == sts.Email
	}
	return stsExists
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql",
		"rector:admin123@tcp(localhost:8889)/colegio")
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		stsquery, err := db.Query("SELECT * from estudiantes;")
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(err.Error())
		} else {
			for stsquery.Next() {
				var sts student
				err := stsquery.Scan(&sts.Id, &sts.Name, &sts.Email, &sts.Grade)
				if err != nil {
					fmt.Println(err)
					json.NewEncoder(w).Encode(err.Error())
				} else {
					students := append(students, sts)
					json.NewEncoder(w).Encode(students)
				}
			}
		}
	}
	defer db.Close()
}
