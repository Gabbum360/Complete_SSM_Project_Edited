package util
import "fmt"

// func Main(){
// 	fmt.Println("keep going")
// }

const DbDriver = "mysql"
const User = "root"
const Password = "gabriel1996bum"
const TableName = "students"
const DbName = "go_school_management_system"

//DataSourceName = "root:gabriel1996bum@tcp(127.0.0.1:3306)/go_school_management_system?charset=utf8"
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8", User, Password, DbName)