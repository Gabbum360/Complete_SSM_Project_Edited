package model

type Teachers struct{
	Teacher_ID int		`json:"teacher_id"`
	First_Name string	`json:"first_name" validate:"required"`
	Age int 			`json:"age" validate:"required"`
	Sex string			`json:"sex" validate:"required"`
	Email string		`json:"email" validate:"required"`
	Phone_Number string	`json:"phone_number" validate:"required"`
	Staff_No int		`json:"staff_no" validate:"required"`
	Country string		`json:"country" validate:"required"`
}

