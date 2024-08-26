package router

import (
	contract "example/re/contractCall"
	processor "example/re/middleware"
	"example/re/types"
	"fmt"
	"log"

	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
)


func Start(config types.Config) {
	router := gin.Default()

	router.POST("/register/:address", func(c *gin.Context) {
		userAddress := c.Params.ByName("address")
		contract.RegisterUser(config.Instance,config.Auth, common.HexToAddress(userAddress))
		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.GET("/createSignature/:privateKey", func(c *gin.Context) {
		userPrivateKey := c.Params.ByName("privateKey")
		sig := processor.CreateDataAndSign(config, userPrivateKey);
		fmt.Println("signature", sig);
		c.IndentedJSON(http.StatusAccepted, sig)
	})

	router.POST("/unregister/:address", func(c *gin.Context) {
		userAddress := c.Params.ByName("address")
		contract.UnRegister(config.Instance,config.Auth, common.HexToAddress(userAddress))
		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.POST("/remove/:signature", func(c *gin.Context) {
		signature, err := hexutil.Decode(c.Params.ByName("signature"));
		if err != nil {
			log.Fatal(err);
		}

		recoveredAddress := processor.CheckSignature(signature);
	
		contract.RemoveUser(config.Instance,config.Auth, recoveredAddress)
		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.Run("localhost:8080")
}