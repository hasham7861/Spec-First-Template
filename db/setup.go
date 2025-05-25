package db

import (
	"log"
	"sample-spec-first-api/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Single global instance to DB
var db *gorm.DB

var logPrefix = "[DB]"

func InitDB(dbPath string) {
	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if(err != nil) {
		log.Fatalf(logPrefix + "Failed to connect to database: %v", err)
	}

	migrateModels()
	log.Println(logPrefix +  "Database connection established")
}

// Add your models here as they're created
var migrations = []interface{} {
	&models.User{},
}

// migrateModels automatically migrates all models
func migrateModels() {
	for _, model := range migrations {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf(logPrefix + "Failed to migrate model %T: %v", model, err)
		}
		log.Printf(logPrefix + "Model %T migrated successfully", model)
	}
}

func GetDBInstance() *gorm.DB {
	return db
}