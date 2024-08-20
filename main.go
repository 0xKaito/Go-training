package main

import (
	// "database/sql"
	"encoding/json"
	"example/api-gin/database"
	"example/api-gin/types"
	"example/api-gin/constants"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"example/api-gin/jsonDataBase"
)


func main() {

	var apiConfig = types.ApiConfig{}
	apiConfig.Initialize();

	
	apiConfig.Wg.Add(3);
	
	go getData(apiConfig, apiConfig.Wg)
	go jsonDatabase.Start(&apiConfig, apiConfig.Wg)
	go database.Start(apiConfig.EmployeeChannel, apiConfig.Wg, apiConfig.Ctx, apiConfig.SyncInterval, apiConfig.Db)

	apiConfig.Wg.Wait()
}

func getData(apiConfig types.ApiConfig, wg *sync.WaitGroup) {
	resp, err := http.Get(constants.API_URL)
	if err != nil {
		log.Fatalf("Error: %v", fmt.Errorf("error fetching API data: %v", err))
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error: %v", fmt.Errorf("error reading API response: %v", err))
	}

	var data database.APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Error: %v", fmt.Errorf("error unmarshaling JSON: %v", err))
	}

	quit := make(chan error);

	defer wg.Done();

	for {
		select {
		case apiConfig.EmployeeChannel <- data.Data:
			continue;
		case quit <- err:
			return;
		}
	}
}
