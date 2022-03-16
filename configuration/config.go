package configuration

type Configuration struct{
	Server ServerConfiguration
	DataBase DataBaseConfiguration
	Swagger SwaggerConfig
	PATH string
	VAR string
}

type DataBaseConfiguration struct{
	DbDriver string
	User string
	Password string
	TableName string
	DbName string
}

type ServerConfiguration struct{
	Port int
}

type SwaggerConfig struct{
	//  Comapany Api:
//   version: 0.0.1
//   title: Comapany Api
//  Schemes: http, https
//  Host: localhost:5000
//  BasePath: /
//  Produces:
}