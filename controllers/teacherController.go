package controllers

import (
	"fmt"
	"encoding/json"
	"mySSM/model"
	"net/http"
	"io/ioutil"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"mySSM/repository"
)

func GetAllTeachers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var allTeachers []model.Teachers
	allTeachers = make([]model.Teachers, 0)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Teacher_ID, First_Name, Age, Staff_No, Country from teachers")
	if err !=nil{
		panic(err.Error())
	}
	for result.Next(){
		var info model.Teachers
		err := result.Scan(&info.Teacher_ID, &info.First_Name, &info.Age, &info.Staff_No, &info.Country)
		allTeachers = append(allTeachers, info)
		if err !=nil{
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(allTeachers)
	}
}
//get Teacher by ID ...
func GetOneTeacher(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select First_Name, Age, Sex, Country from teachers where Teacher_ID = ?", params["id"])
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()
	var teacher model.Teachers
	for result.Next(){
		err = result.Scan(&teacher.First_Name, &teacher.Age, &teacher.Sex, &teacher.Country)
		if err !=nil{
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(teacher)
	}
}
//get Teacher by Staff_No ...
func GetTeacherByStaffNo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Teacher_ID, First_Name, Age, Sex, Country from teachers where Staff_No = ?", params["id"])
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()
	var teacher model.Teachers
	for result.Next(){
		err = result.Scan(&teacher.Teacher_ID, &teacher.First_Name, &teacher.Age, &teacher.Sex, &teacher.Country)
		if err !=nil{
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(teacher)
	}
}

func UpdateTeacherRecord(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	var teacher model.Teachers
		body, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		panic(err.Error())
	}
	json.Unmarshal(body, &teacher)
	validate = validator.New()
	///use "structExcept()" to specify which field not to validate.
	err = validate.StructExcept(&teacher,"Teacher_ID", "Sex", "Email", "Phone_Number", "Staff_No", "Country")
	//err = validate.Struct(teacher)
	var errorResult string
	if err !=nil{
		for _, err := range err.(validator.ValidationErrors){
			errorResult += err.StructField() +" "+ "is" +" " + err.ActualTag() + ";"	//this line is passes the "required error message" to postMan.
			//json.NewEncoder(w).Encode(errorResult)
		}
		json.NewEncoder(w).Encode(errorResult)
	}
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Exec("UPDATE teachers SET First_Name=?, Age =? WHERE Teacher_ID = ?", teacher.First_Name, teacher.Age, teacher.Teacher_ID )
	if err !=nil{
		panic(err.Error())
	}
	rowsUpdated, err := result.RowsAffected()
	if err !=nil{
		panic(err.Error())
	}
	fmt.Println(rowsUpdated)
	json.NewEncoder(w).Encode(rowsUpdated)
}

func DeleteRow(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	stmt, err := dataBaseConfig.Db.Prepare("delete from teachers where Teacher_ID = ?")
	if err !=nil{
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(params["id"])
	if err !=nil{
		panic(err.Error())
	}
	json.NewEncoder(w).Encode("deleted")
	fmt.Fprintf(w, "row with ID = %s has been deleted", params["id"])
}