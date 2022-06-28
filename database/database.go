package database

import (
	"log"
	"time"

	"github.com/Seyz123/lightshot/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Image{})

	go checkDeletable()
}

func checkDeletable() {
	tick := time.Tick(12 * time.Hour)

	for {
		<-tick
		db, err := GetDB()
		if err != nil {
			log.Println("Failed to check deletable images")
			continue
		}

		var images []models.Image
		db.Find(&images)

		for _, image := range images {
			diff := time.Now().Sub(image.CreatedAt)
			if diff.Hours() > 24 {
				err = db.Delete(&image).Error
				if err != nil {
					log.Println("Failed to delete image")
					continue
				}

				log.Println("Deleted image " + image.ID)
			}
		}
	}
}

func GetDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
}
