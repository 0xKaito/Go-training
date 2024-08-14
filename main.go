package main

import (
	"encoding/json"
	"example/api-gin/database"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"example/api-gin/jsonDataBase"
)

var apiURL = "https://dummy.restapiexample.com/api/v1/employee/1"

func main() {

	var employeeChannel = make(chan database.Employee)
	var wg sync.WaitGroup

	wg.Add(1)

	go getData(employeeChannel)
	go jsonDatabase.Start(employeeChannel)
	go database.Start(employeeChannel, &wg)

	wg.Wait()
}

func getData(employeeChannel chan database.Employee) {
	resp, err := http.Get(apiURL)
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

	employeeChannel <- data.Data;
}
