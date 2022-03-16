package controllers

import(
	"encoding/json"
	"mySSM/model"
	"net/http"
	"mySSM/repository"

)

func GetClasses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var allClass []model.Classes
	allClass =make([]model.Classes,0)
	dataBaseConfig := util.Init()	//calling the func that connects to the dataBase.
	result, err := dataBaseConfig.Db.Query("select Class_ID, Name, Arm_ID, Teacher_ID from classes")
	if err !=nil{
		panic(err.Error())
	}
	defer result.Close()

	for result.Next(){
		var info model.Classes
		err := result.Scan(&info.Class_ID, &info.Name, &info.Arm_ID, &info.Teacher_ID)
		allClass = append(allClass, info)
		if err !=nil{
			panic(err.Error())
		}
		json.NewEncoder(w).Encode(allClass)
	}
}

// func editStudentAndClass(w http.ResponseWriter, r *http.Request){
// 	var editRecord StudentClass
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err !=nil{
// 		panic(err.Error())
// 	}
// 	json.Unmarshal(body,&editRecord)
// 	result, err := db.Exec("update studentclass set Student_ID =?, ClassArm_ID =? where ST_ID =?", editRecord.Student_ID, editRecord.ClassArm_ID, editRecord.ST_ID)
// 	if err !=nil{
// 		panic(err.Error())
// 	}
// 	updatedRecord, err := result.RowsAffected()
// 	if err !=nil{
// 		panic(err.Error())
// 	}
// 	json.NewEncoder(w).Encode(updatedRecord)
// }