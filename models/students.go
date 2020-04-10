package models

import (
	u "api-test/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//a struct to rep students
type Student struct {
	gorm.Model
	Name    string `json:"name"`
	DocType string `json:"doc_type"`
	DocNum  uint   `json:"doc_num"`
	Grade   string `json:"grade"`
	UserId  uint   `json:"user_id"`
}

func (student *Student) Validate() (map[string]interface{}, bool) {

	if student.Name == "" {
		return u.Message(false, "Student name should be on the payload"), false
	}

	if student.DocType == "" {
		return u.Message(false, "Document Type should be on the payload"), false
	}

	if student.DocNum <= 0 {
		return u.Message(false, "Document Number should be on the payload"), false
	}

	if student.Grade == "" {
		return u.Message(false, "Grade should be on the payload"), false
	}

	if student.UserId <= 0 {
		return u.Message(false, "The user was not recognize"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (student *Student) Create() map[string]interface{} {

	if resp, ok := student.Validate(); !ok {
		return resp
	}

	GetDB().Create(student)

	resp := u.Message(true, "success")
	resp["student"] = student
	return resp
}

func GetStudentByDocument(documentNum uint) *Student {

	student := &Student{}
	err := GetDB().Table("students").Where("doc_num = ?", documentNum).First(student).Error
	if err != nil {
		return nil
	}
	return student
}

func GetStudents() []*Student {

	students := make([]*Student, 0)
	err := GetDB().Table("students").Find(&students).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return students
}
