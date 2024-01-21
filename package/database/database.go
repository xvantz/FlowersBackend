package database

import (
	"log"
	"milena/package/models"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func Connect() *gorm.DB {
    // connect URL database
   databaseURL := os.Getenv("DB_URL")
    // Open PostgreSQL connection
   db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
    NowFunc: func() time.Time {
		ti, _ := time.LoadLocation("Asia/Shanghai")
		return time.Now().In(ti)
	},
   })
    // Error
   if err != nil {
    log.Fatalln(databaseURL)
   }
   // Migrate models
   db.AutoMigrate(&models.Work{}, &models.Post{})

   return db
}