package model

type Class_Teachers struct{
	CT_ID int			`json:"ct_id"`
	ClassArm_ID int		`json:"classarm_id"`
	Teacher_ID int		`json:"teacher_id" validate:"omitempty"`
}

type Classes struct{
	Class_ID int		`json:"class_id" validate:"required"`
	Name string			`json:"name" validate:"required"`
	Arm_ID int			`json:"arm_id"`
	Teacher_ID int		`json:"teacher_id"`
}

type ClassArm struct{
	ClassArm_ID int		`json:"classarm_id"`
	Department string	`json:"department" validate:"required"`
	Class_ID int		`json:"class_id"`
	LABooks string		`json:"labooks"`
	Arm_ID int			`json:"arm_id"`
	Subject_ID int		`json:"subject_id"`
	Student_ID int		`json:"student_id"`
}

type Arms struct{
	Arm_ID int			`json:"arm_id"`
	Name string			`json:"name"`
}

type ClassSubjects struct{
	CS_ID int			`json:"cs_id"`
	Student_ID int 		`json:"student_id"`
	ClassArm_ID int		`json:"classarm_id"`
	Subject_ID int		`json:"subject_id"`
}

// type Library struct{
// 	ID int
// 	Title string
// 	Author string
// 	Date_Published string
// }