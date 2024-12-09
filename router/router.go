package router

import (
	contract "example/re/contractCall"
	processor "example/re/middleware"
	"example/re/types"

	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
)


func Start(config types.Config) {
	router := gin.Default()

	router.POST("/register/:address", func(c *gin.Context) {
		userAddress := c.Params.ByName("address")
		err := contract.RegisterUser(config.Instance,config.Auth, common.HexToAddress(userAddress))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		c.IndentedJSON(http.StatusAccepted, 1)
	})

	// to test signature validation

	// router.GET("/createSignature/:privateKey", func(c *gin.Context) {
	// 	userPrivateKey := c.Params.ByName("privateKey")
	// 	sig, userAddress := processor.CreateDataAndSign(config, userPrivateKey);
	// 	fmt.Println("signature", hexutil.Encode(sig));
	// 	fmt.Println("user address", userAddress);
	// 	c.IndentedJSON(http.StatusAccepted, sig)
	// })

	router.POST("/removeAddress/:address/:signature", func(c *gin.Context) {
		userAddress := c.Params.ByName("address");
		signature, err := hexutil.Decode(c.Params.ByName("signature"));
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		err = processor.VerifySignature(config, signature, userAddress);
		if err != nil {
			c.IndentedJSON(http.StatusForbidden, err)
		}

		err = contract.RemoveUser(config.Instance,config.Auth, common.HexToAddress(userAddress));
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.POST("/unregister/:address", func(c *gin.Context) {
		userAddress := c.Params.ByName("address")
		err := contract.UnRegister(config.Instance,config.Auth, common.HexToAddress(userAddress))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.POST("/remove/:signature", func(c *gin.Context) {
		signature, err := hexutil.Decode(c.Params.ByName("signature"));
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		recoveredAddress, err := processor.CheckSignature(signature);
		if err != nil {
			c.IndentedJSON(http.StatusForbidden, err)
		}
	
		err = contract.RemoveUser(config.Instance,config.Auth, recoveredAddress)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		c.IndentedJSON(http.StatusAccepted, 1)
	})

	router.Run("localhost:8080")
}