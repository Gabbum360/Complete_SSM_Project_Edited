package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
		"github.com/spf13/viper"
	c "mySSM/configuration"
)

type DataBaseConfig struct{
  Db *sql.DB
}

var db *sql.DB

//create a dataBase func where your operations calls from. using viper
func Init() DataBaseConfig {
	var configuration c.Configuration
	err := viper.Unmarshal(&configuration)
	if err != nil {
	fmt.Printf("Unable to decode into struct, %v", err)
	}
	//DataSourceName = "root:gabriel1996bum@tcp(127.0.0.1:3306)/go_school_management_system?charset=utf8"
	var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8", configuration.DataBase.User, configuration.DataBase.Password, configuration.DataBase.DbName)
	db, err = sql.Open(configuration.DataBase.DbDriver, DataSourceName)
	if err !=nil{
		panic(err.Error())
	}else{
		fmt.Println("Access Granted!")
	}
  //  defer db.Close()  //ensure to always close DB
    return DataBaseConfig {
		Db: db,
	}
}