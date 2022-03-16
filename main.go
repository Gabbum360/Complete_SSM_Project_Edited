package main

import (
	//"database/sql"
	//	"encoding/json"
	// "strings"
	//"errors"
	"fmt"
	//	"io/ioutil"
	"mySSM/controllers"
//	"mySSM/util"
	"net/http"
	//"regexp"
	//	validation "github.com/go-ozzo/ozzo-validation"
	//	"github.com/go-playground/validator"
	//	"github.com/go-ozzo/ozzo-validation/is"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	c "mySSM/configuration"
	"github.com/go-openapi/runtime/middleware"
)

//var db *sql.DB
//var err error

// ValidateUpdateReq the given struct.
//func ValidateUpdateReq(i interface{}) (bool, map[string]string) {
//     errors := make(map[string]string)
//     v := validator.New()
//     v.SetTagName("updatereq")

//     if err := v.Struct(i)
// 		err != nil {
//         for _, err := range err.(validator.ValidationErrors) {
//             if err.Tag() == "email" {
//                 errors[strings.ToLower(err.Field())] = "Invalid E-mail format."
//                 continue
//             }
//             errors[strings.ToLower(err.Field())] = fmt.Sprintf("%s is %s %s", err.Field(), err.Tag(), err.Param())
//         }
//         return false, errors
//     }
//     return true, nil
// }


func main(){	
		
	   // documentation for developers
opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
sh := middleware.SwaggerUI(opts, nil)
//set the name of the configuration file
	viper.SetConfigName("config")
//set the path to look for the configuration files
	viper.AddConfigPath("configuration")
//enable viper to environment variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	//viper.SetConfigFile(".env")
	var configuration c.Configuration

	if err := viper.ReadInConfig()
	err !=nil {
		fmt.Printf("error reading config file, %s", err)
	}
//set undefined variables
	viper.SetDefault("database.dbname", "test_db")
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.DataBase)
	fmt.Println("Port is\t\t", configuration.Server.Port)
	fmt.Println("PATH is\t", configuration.PATH)
	fmt.Println("VAR is\t", configuration.VAR)
// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Database is\t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))
	fmt.Println("PATH is\t", viper.GetString("PATH"))
	fmt.Println("VAR is\t", viper.GetString("VAR"))

//opening the database for this API
/* 	db, err = sql.Open(util.DbDriver, util.DataSourceName)
	if err !=nil{
	panic(err.Error())
	}else{
 	fmt.Println("connected successfully!")
	}
 defer db.Close()  //ensure to always close DB*/

r := mux.NewRouter()
r.Handle("/docs", sh)
//Students Controller 
r.HandleFunc("/getAll", controllers.GetAllStudents).Methods("GET")
r.HandleFunc("/justOne/{id}", controllers.GetOneStudent).Methods("GET")//by Student_ID
//get student by unique number...
r. HandleFunc("/unique/{id}", controllers.GetStudentByNo) //by Student_No
//get student by Class ...
r.HandleFunc("/class/{id}", controllers.GetStudentByClass) //by ClassArm_ID
//get student by subject ...
r.HandleFunc("/subjects/{id}", controllers.GetStudentBySubject) // by SS_ID
r.HandleFunc("/subjectAndClass/{subjectid}/{classarmid}", controllers.GetStudentBySubjectAndClass)
r.HandleFunc("/create", controllers.CreateStudent).Methods("POST")
r.HandleFunc("/updateS/{id}", controllers.UpdateRecord) // by Student_ID
r.HandleFunc("delete/{id}", controllers.DeleteStudent) // By Student_ID

//Teacher Controller
r.HandleFunc("/allTeachers", controllers.GetAllTeachers).Methods("GET")
r.HandleFunc("/one/{id}", controllers.GetOneTeacher).Methods("GET")//by Teacher_ID
r.HandleFunc("/staffNo/{id}", controllers.GetTeacherByStaffNo).Methods("GET") //by Staff_No
r.HandleFunc("/updateT/{id}", controllers.UpdateTeacherRecord).Methods("PUT")
r.HandleFunc("/delete/{id}", controllers.DeleteRow).Methods("DELETE")

//StudentClass Controller
r.HandleFunc("/allbyClass/{id}", controllers.GetStudentsByClass) //by ST_ID
r.HandleFunc("/assign", controllers.AssignStudentToClass).Methods("POST")
r.HandleFunc("/remove/{id}", controllers.RemoveStudent) //by ST_ID
//r.HandleFunc("/edit/{id}", editRecord)

//StudentSubject Controller
r.HandleFunc("/assignS", controllers.AssignSubjectToStudent)

//Classes Controller
r.HandleFunc("/allClasses", controllers.GetClasses)
http.ListenAndServe(":8001", r)
fmt.Println("Gabbum")
}