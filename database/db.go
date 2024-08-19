package database

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Registration struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	NewUser   common.Address `json:"new_user"`
}

// db close connection
func Start() *gorm.DB {
	// Get configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open(postgres.Open(
		"host="+dbHost+
			" port="+dbPort+
			" user="+dbUser+
			" password="+dbPassword+
			" dbname="+dbName+
			" sslmode=disable"),
		&gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto-migrate your model to create the corresponding table
	db.AutoMigrate(&Registration{})
	return db
}

func AddUser(db *gorm.DB, address common.Address) {
    registration := Registration{
        NewUser: address,
    }

    if err := db.Create(&registration).Error; err != nil {
        panic("Failed to add user")
    }
}
func RemoveUser(db *gorm.DB, address common.Address) {
    if err := db.Where("new_user = ?", address.Hex()).Delete(&Registration{}).Error; err != nil {
        panic("Failed to remove user")
    }
}
