package controllers

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mySSM/model"
	"net/http"
	"mySSM/repository"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var validate *validator.Validate
//var db *sql.DB

//var err error

/*type menuListImpl struct{}

func NewMenuCategoryListHandler() menu.CategoryListHandler {
	return &menuListImpl{}
}

func (impl *menuListImpl) Handle(params menu.CategoryListParams) middleware.Responder {
	responseVal := &models.Categories{
		&models.Category{
			BcID:       2001,
			BcName:     "Fruits",
			BcIsActive: true,
			BcImageURL: "",
			SubCategories: []*models.SubCategory{
				{
					ScID:       1002,
					ScName:     "Apple",
					ScImageURL: "",
					ScIsActive: true,
				},
			},
		},
	}
	return menu.NewCategoryListOK().WithPayload(*responseVal)
}*/

func GetAllStudents(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var allStudents []model.Students
	allStudents = make([]model.Students, 0)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Student_ID, Name, ClassArm_ID, Country from students")
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()

	for result.Next(){
		var info model.Students
		err = result.Scan(&info.Student_ID, &info.Name, &info.ClassArm_ID, &info.Country)
		allStudents = append(allStudents, info)
		if err !=nil{
			panic(err)
		}
		//fmt.Println(students) commented bcos i only want it from postman.
		json.NewEncoder(w).Encode(allStudents)
	}
}

func GetOneStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Student_ID, Name, Age, ClassArm_ID, Country from students where Student_ID = ?", params["id"])
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()
	
	for result.Next(){
		var student model.Students
		err := result.Scan(&student.Student_ID, &student.Name, &student.Age, &student.ClassArm_ID, &student.Country)
		if err !=nil{
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(student)
	}
}

//get student by Unique no "Student_No"...
func GetStudentByNo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Student_ID, Name, Age, ClassArm_ID, Country from students where Student_No = ?", params["id"])
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()
	var student model.Students
	for result.Next(){
		err := result.Scan(&student.Student_ID, &student.Name, &student.Age, &student.ClassArm_ID, &student.Country)
		if err !=nil{
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(student)
	}
}
//get student by Class ...
func GetStudentByClass(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Student_ID, Name, Age, ClassArm_ID, Country from students where ClassArm_ID = ?", params["id"])
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()
	var student model.Students
	for result.Next(){
		err = result.Scan(&student.Student_ID, &student.Name, &student.Age, &student.ClassArm_ID, &student.Country)
		if err !=nil{
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(student)
	}
}
// get student by subjects ...
func GetStudentBySubject(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Student_ID, Subject_ID from studentsubjects where SS_ID = ?", params["id"])
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()
	
	for result.Next(){
		var subject model.StudentSubjects
		err = result.Scan(&subject.Student_ID, &subject.Subject_ID)
		if err !=nil{
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(subject)
	}
}
//get by subject and class ...
func GetStudentBySubjectAndClass(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Student_ID, ClassArm_ID, Subject_ID from classsubjects WHERE Subject_ID =? AND ClassArm_ID =?", params["subjectid"], params["classarmid"])
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()

	for result.Next(){
		var subjectClass model.ClassSubjects
		err = result.Scan(&subjectClass.Student_ID, &subjectClass.ClassArm_ID, &subjectClass.Subject_ID)
		if err !=nil{
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(subjectClass)
	}
}
//receive request, check validation and send back error message if available...
func CreateStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	student:= model.Students{}
	getBodyFromPostMan, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(getBodyFromPostMan, &student)

	if err !=nil{
		panic(err.Error())
	}
	//check if theres an error in the validation...
	validate = validator.New()
	err = validate.Struct(&student)
	var errorResult string	//create a variable of string...
    if err !=nil {
			for _, err := range err.(validator.ValidationErrors) {	//use for loop to loop through the fields for validation...
				errorResult += err.StructField() + " " + "is" + " " + err.ActualTag() + ","	//...use concatenation to arrange the errors on a single line
			/*fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println("---------------")
	        return*/
	        }
			json.NewEncoder(w).Encode(errorResult)
			return
}
dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	_, err = dataBaseConfig.Db.Exec(`insert into students (Name, Age, Sex, Student_No, ClassArm_ID, Country) values ( ?, ?, ?, ?, ?, ?)`,
student.Name, student.Age, student.Sex, student.Student_No, student.ClassArm_ID, student.Country)
if err !=nil{
	panic(err.Error())
}
fmt.Fprintf(w, "new student created")
}

func UpdateRecord(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type", "application/json")
	//sparams := mux.Vars(r)
	var student model.Students
	body, err := ioutil.ReadAll(r.Body)
	if err !=nil{
	panic(err.Error())
	}
	json.Unmarshal(body, &student)
	validate = validator.New()
	err = validate.Struct(&student)
	var errorResult string
	if err !=nil{
		for _, err:= range err.(validator.ValidationErrors){
			if err.StructField() == "Age"{	//...the "if" statement on this line is to validate a particular Field 
	        errorResult += err.StructField() +" "+ "is" + " " + err.ActualTag() + "; "	//ActualTag is the validation rule "required..."
			}
        }
		json.NewEncoder(w).Encode(errorResult)
		return
    }
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Exec("update students set Age =? WHERE Student_ID = ?",student.Age, student.Student_ID )
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
func DeleteStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	stmt, err := dataBaseConfig.Db.Prepare("delete from students where Student_ID = ?")
	if err !=nil{
	panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err !=nil{
		panic(err.Error())
	}
	fmt.Fprintf(w, "studentRow with ID = %s was deleted", params["id"])
}

func AssignSubjectToStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newSubjectToStudent model.StudentSubjects
	body, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		panic(err.Error())
	}
	json.Unmarshal(body, &newSubjectToStudent)
	validate = validator.New()
	err = validate.Struct(&newSubjectToStudent)
	var errorResult string
	if err != nil { 
		for _, err := range err.(validator.ValidationErrors){
	    errorResult += err.StructField() +" "+ "is" + err.ActualTag() + ";"
    	}
	    json.NewEncoder(w).Encode(errorResult)
    }
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	_, err = dataBaseConfig.Db.Exec("insert into studentsubjects (Subject_ID, Student_ID) values (?, ?)", newSubjectToStudent.Subject_ID, newSubjectToStudent.Student_ID)
	if err !=nil{
		panic(err.Error())
	}
	fmt.Fprintf(w, "new data inserted")
}