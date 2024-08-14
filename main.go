package main

import (
	contract "example/re/contractCall"
	logic "example/re/middleware"
	"log"

	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
)

func main() {
	instance, auth := logic.Start();

	router := gin.Default()

	router.POST("/register/:address", func(c *gin.Context) {
		userAddress := c.Params.ByName("address")
		contract.RegisterUser(instance, auth, common.HexToAddress(userAddress))
		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.POST("/unregister/:address", func(c *gin.Context) {
		userAddress := c.Params.ByName("address")
		contract.UnRegister(instance, auth, common.HexToAddress(userAddress))
		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.POST("/remove/:signature", func(c *gin.Context) {	
		signature, err := hexutil.Decode(c.Params.ByName("signature"));
		if err != nil {
			log.Fatal(err);
		}

		recoveredAddress := logic.CheckSignature(signature);
	
		contract.RemoveUser(instance, auth, recoveredAddress)
		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.Run("localhost:8080")

}
