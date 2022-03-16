package model
import (

)
type StudentClass struct{
	ST_ID int 			`json:"sc_id"`
	Student_ID int		`json:"student_id" validate:"required,gte=3"`
	ClassArm_ID int		`json:"classarm_id"`
	//Class_ID int		`json:"class_id"`
}