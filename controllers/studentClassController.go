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

/*func (s Students) validate() error{
return validation.ValidateStruct(&s, 
	validation.Field(&s.Name, validation.Required, validation.Match(regexp.MustCompile("^[A_Z]{2}$"))),
	validation.Field(&s.Country, validation.Required, validation.NotNil.Error("please check that the Country is not blank")),
	validation.Field(&s.Sex, validation.Required, validation.Length(3,10)))
}
// s := Students{
// 	Name: "Moses",
// 	Country: "Unknown",
// 	Sex: "Male" 
// }	
// err := s.validate()
// fmt.Println(err)*/

func GetStudentsByClass(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// var allStudents []StudentClass
	// allStudents = make([]StudentClass, 0)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Student_ID, ClassArm_ID from studentclass where ST_ID = ?", params["id"])
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()

	for result.Next(){
	var info model.StudentClass
	err = result.Scan(&info.Student_ID, &info.ClassArm_ID)
	if err !=nil{
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(info)
	}
}
//assign student to StudentClass
func AssignStudentToClass(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	newStudent := model.StudentClass{}
	result, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(result, &newStudent)
	if err !=nil{
		panic(err.Error())
	}
	validate = validator.New()
	err = validate.Struct(&newStudent)
	var errorResult string
	if err !=nil{
		for _, err := range err.(validator.ValidationErrors){
			errorResult += err.StructField() +" "+ "is" + " " +  err.ActualTag() + ","
		}
		json.NewEncoder(w).Encode(errorResult)
		return
	}
	// errors := make(map[string]string)
	// //var errorResult string
	// v := validator.New()
	// v.SetTagName("updatereg")
	// if err := v.Struct(s)
	// err != nil{
	// 	for errorResult, err:= range err.(validator.ValidationErrors){
	// 		if err.Tag() == "Student_ID"{
	// 			errors[strings.ToLower(err.Field())] = "invalid Student id"
	// 			continue
	// 		}
	// 		errors[strings.ToLower(err.Field())] = fmt.Sprintf("%s is %s", err.Field(), err.ActualTag())
	// 		json.NewEncoder(w).Encode(errorResult)
	// 	}
	// }
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	_, err = dataBaseConfig.Db.Exec("insert into studentclass (Student_ID, ClassArm_ID) values ( ?, ?)", newStudent.Student_ID, newStudent.ClassArm_ID)
	if err !=nil{
		panic(err.Error())
	}
	fmt.Fprintf(w,"new Record inserted successfully")
}

func RemoveStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	stmt, err := dataBaseConfig.Db.Prepare("delete from studentclass where ST_ID =?")
	if err !=nil{
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err !=nil{
		panic(err.Error())
	}
	fmt.Fprintf(w, "row with ID = %s has been deleted", params["id"])
}