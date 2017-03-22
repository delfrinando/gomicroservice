package core

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/tiket-dev/tiket-microservice-configuration/config"
)

func ConnectMySQL() (db *gorm.DB) {

	log.Println("Connecting Database ....")
	configs := config.GetConfig()
	mysqlHost := configs.Mysql.Host
	mysqlPort := configs.Mysql.Port
	mysqlDatabase := configs.Mysql.Database
	mysqlUsername := configs.Mysql.Username
	mysqlPassword := configs.Mysql.Password

	db, err := gorm.Open(
		"mysql",
		mysqlUsername+":"+mysqlPassword+"@"+"tcp("+mysqlHost+":"+mysqlPort+")/"+mysqlDatabase+"?charset=utf8&parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	if err = db.DB().Ping(); err != nil {
		log.Println("Error connecting to database.")
		log.Println(err)
	} else {
		log.Println("Connected to database : ", mysqlDatabase)
	}

	return db
}
