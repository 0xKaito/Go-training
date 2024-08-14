package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Employee struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

type APIResponse struct {
	Status  string   `json:"status"`
	Data    Employee `json:"data"`
	Message string   `json:"message"`
}

// db close connection
func Start(employeeChannel chan Employee, wg *sync.WaitGroup) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db := dbConnect()
	
	dbData := <- employeeChannel;

	// db.Create(&dbData)
	// Upsert (insert if not exist, or update if exist)
    err = db.Clauses(clause.OnConflict{
        Columns:   []clause.Column{{Name: "id"}}, // key column for uniqueness
        DoUpdates: clause.AssignmentColumns([]string{"employee_name", "employee_salary", "employee_age", "profile_image"}), // columns to update
    }).Create(&dbData).Error

	if err != nil {
        fmt.Println("Error upserting user:", err)
        return
    }

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error fetching sql db:", err)
        return
    }
	
	sqlDB.Close()

	wg.Done()
}

func dbConnect() *gorm.DB {
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
	db.AutoMigrate(&Employee{})
	return db
}
