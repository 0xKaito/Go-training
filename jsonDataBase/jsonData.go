package jsonDatabase

import (
	"example/api-gin/database"
	"fmt"
	"log"
	"os"
	"encoding/json"
)

func Start(employeeChannel chan database.Employee) {
	// Marshal the data back to JSON format (for saving to file)
	data := <- employeeChannel;
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error: %v", fmt.Errorf("error marshaling JSON data: %v", err))
	}
	
	// Write the JSON data to a file
	err = os.WriteFile("./employee.json", jsonData, 0644);
	if err != nil {
        log.Fatalf("Error: %v", fmt.Errorf("error writing data to file: %v", err))
	}

	employeeChannel <- data;
}
