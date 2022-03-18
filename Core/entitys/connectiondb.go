package entitys

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// setup database connection is creating a new connection to our database
func DatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	db.AutoMigrate(
		&Administrador{},
	)

	return db

}

// Close database connection method is closin a connection between your app db
func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()

}

//Closedb
func Closedb() {
	var db *gorm.DB = DatabaseConnection()
	CloseConnection(db)

}
