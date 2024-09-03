package main

import (
    "log"
    "github.com/tendermint/tendermint/abci/server"
    "github.com/tendermint/tendermint/abci/types"
    "sync"
    "encoding/json"
    v1 "example/tendermint/proto"  // Import the generated protobuf package
    "google.golang.org/protobuf/proto"
)
type MyApp struct {
    types.BaseApplication

    state map[string]int64
    mutex sync.RWMutex
}

func NewMyApp() *MyApp {
    return &MyApp{
        state: make(map[string]int64),
    }
}

func (app *MyApp) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
    // Unmarshal the transaction data using Protobuf
    var tx v1.Tx
    err := proto.Unmarshal(req.Tx, &tx)
    if err != nil {
        log.Printf("Failed to unmarshal transaction: %v", err)
        return types.ResponseDeliverTx{Code: 1} // Error code
    }

    // Update the sender's and recipient's balances
    app.mutex.Lock()
    defer app.mutex.Unlock()

    senderBalance := app.state[tx.Sender]
    if senderBalance < tx.Amount {
        return types.ResponseDeliverTx{
            Code: 2, // Insufficient funds error
        }
    }

    app.state[tx.Sender] -= tx.Amount
    app.state[tx.Recipient] += tx.Amount

    return types.ResponseDeliverTx{Code: 0} // Success code
}

func (app *MyApp) Query(req types.RequestQuery) types.ResponseQuery {
    // The key to query is expected in the path
    account := req.Path

    app.mutex.RLock()
    balance, exists := app.state[account]
    app.mutex.RUnlock()

    if !exists {
        return types.ResponseQuery{Code: 1, Log: "account not found"} // Account not found
    }

    // Create a JSON response
    response := v1.ResponseQuery{
        Value: string(balance),
    }
    resData, _ := json.Marshal(response)

    return types.ResponseQuery{
        Code:  0, // Success code
        Value: resData,
    }
}

func main() {
    // Create the ABCI application
    app := NewMyApp()

    // Create the ABCI server
    address := "tcp://127.0.0.1:26658"
    s, err := server.NewServer(address, "socket", app)
    if err != nil {
        log.Fatalf("Failed to create ABCI server: %v", err)
    }

    s.Start()

    defer s.Stop()

    select {}
}