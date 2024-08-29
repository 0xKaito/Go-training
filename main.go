package main

import (
    "log"
    "github.com/tendermint/tendermint/abci/server"
)

func main() {
    // Create the ABCI application
    app := NewMyApp()

    // Create the ABCI server
    address := "tcp://127.0.0.1:26658"
    s, err := server.NewServer(address, "socket", app)
    if err != nil {
        log.Fatalf("Failed to create ABCI server: %v", err)
    }

    // Start the server
    s.Start()

    // Wait for the server to stop
    defer s.Stop()

    // Keep the application running
    select {}
}