package types

import (
	"context"
	"database/sql"
	"example/api-gin/database"
	"example/api-gin/constants"
	"fmt"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type ApiConfig struct {
	Db *gorm.DB
	SyncInterval time.Duration
	SqlDB *sql.DB
	EmployeeChannel chan database.Employee
	Ctx context.Context
	Wg *sync.WaitGroup
}

func (config *ApiConfig) Initialize() {
	ctx, _ := context.WithCancel(context.Background())
	config.Ctx = ctx;
	config.EmployeeChannel = make(chan database.Employee);

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	config.Db = database.DbConnect();
	sqlDB, err := config.Db.DB()
	if err != nil {
		fmt.Println("Error fetching sql db:", err)
    	return
    }
	config.SqlDB = sqlDB
	config.SyncInterval = time.Duration(constants.SLEEP_INTERVAL)*time.Second;
	var wg sync.WaitGroup;
	config.Wg = &wg;
}