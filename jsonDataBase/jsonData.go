package jsonDatabase

import (
	"encoding/json"
	"example/api-gin/database"
	"example/api-gin/types"
	"fmt"
	"log"
	"os"
	"time"
	"sync"
)

func Start(apiConfig *types.ApiConfig, wg *sync.WaitGroup) {
	defer wg.Done();

	for {
		select {
		case <- apiConfig.Ctx.Done():
			return
		default:
			data := <- apiConfig.EmployeeChannel
			insertData(data);
			
			apiConfig.EmployeeChannel <- data
		}

		time.Sleep(apiConfig.SyncInterval)
	}
}

func insertData(data database.Employee) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error: %v", fmt.Errorf("error marshaling JSON data: %v", err))
	}
	
	// Write the JSON data to a file
	err = os.WriteFile("./employee.json", jsonData, 0644);
	if err != nil {
		log.Fatalf("Error: %v", fmt.Errorf("error writing data to file: %v", err))
	}	
}
