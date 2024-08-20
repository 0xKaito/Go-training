package router

import (
	contract "example/re/contractCall"
	processor "example/re/middleware"
	"example/re/store"
	"log"

	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
)


func Start(instance *store.Store, auth *bind.TransactOpts) {
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

		recoveredAddress := processor.CheckSignature(signature);
	
		contract.RemoveUser(instance, auth, recoveredAddress)
		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.Run("localhost:8080")
}