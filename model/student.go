package model


// swagger: model student
type Students struct{
	// student_id
	// in: int64
Student_ID int		`json:"student_id"`
	// Name of the student
	// in: string
Name string			`json:"name" validate:"required"`
	// Age of student
	// in: int64
Age int				`json:"age" validate:"required"`
	// Sex of student
	// in: string
Sex string			`json:"sex"`
	// Student_no of student
	// in: int64
Student_No int64		`json:"student_no"`
	// Classarm_id of student
	// in: int64
ClassArm_ID string	`json:"classarm_id" validate:"required"`
	// Country of student
	// in: string
Country string		`json:"country" validate:"required"`
}

/*type ReqAddCompany struct {
   // Name of the company
   // in: string
   Name string `json:"name" validate:"required,min=2,max=100,alpha_space"`
   // Status of the company
   // in: int64
   Status int64 `json:"status" validate:"required"`
}*/

/*// func (s Students) validate() error{
// return validation.ValidateStruct(&s, 
// 	validation.Field(&s.Name, validation.Required, validation.Match(regexp.MustCompile("^[A_Z]{2}$"))),
// 	validation.Field(&s.Country, validation.Required, validation.NotNil.Error("please check that the Country is not blank")),
// 	validation.Field(&s.Sex, validation.Required, validation.Length(3,10)),
// validation.Field(&s.Student_ID, validation.Required, validation.Match(regexp.MustCompile("^[A_Z]{3}$"))))
// }
// s := Students{
// 	Name: "Moses",
// 	Country: "Unknown",
// 	Sex: "Male",
// 	Student_ID: "506",
// }	
// err := s.validate()
// fmt.Println(err)*/


// swagger:model StudentSubjects
type StudentSubjects struct{
	// SS_ID of studentsubjects
	// in: int64
	SS_ID int 			`json:"ss_id"`
	// Subject_ID of studentsubjects
	// in: int64
	Subject_ID int		`json:"subject_id"`
	// Student_ID of studentsubjects
	// in: int64
	Student_ID int		`json:"student_id"`
}