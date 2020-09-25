package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"github.com/nurrizkyimani/golang-mux-backend/model"
)

var (
	//DBConn is xxx
	DBConn *gorm.DB
)

//InitDatabase xxxxs
func InitDatabase() {
	var err error
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbURL := os.Getenv("db_url")
	fmt.Println(dbURL)

	DBConn, err = gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Print(err)
		panic("failed to conenct ")
	}

	

	fmt.Println("Connection opened database")

	DBConn.AutoMigrate(&model.Blog{})
	DBConn.AutoMigrate(&model.BlogComment{})

	DBConn.AutoMigrate(&model.Article{})

	fmt.Println("database migrated")

	// db = conn
	// db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}
