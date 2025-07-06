package	models 


import (
	"log"
)


func StartDB(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

  	db.AutoMigrate(&Idea{})
	log.Println("Database is running ! ")

}
