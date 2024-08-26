package processor

import (
	"context"
	"crypto/ecdsa"
	"example/re/types"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func CheckSignature(signature []byte) common.Address {
	hash := crypto.Keccak256Hash(types.SignatureData)
	fmt.Println(hash.Hex())

	// Recover the public key
	pubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatalf("Failed to recover public key: %v", err)
	}

	// Get the signer's address
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Printf("Recovered address: %s\n", recoveredAddr.Hex())

	return recoveredAddr;
}

func CreateDataAndSign(config types.Config, privateKEY string) ([]byte, common.Address) {
	privateKey, err := crypto.HexToECDSA(privateKEY)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	user := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := config.Client.PendingNonceAt(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	dataType := CreatTypedData(config.ContractAddress, config.ChainId, user.Hex(), int64(nonce));
	sig, err := signTypedData(dataType, privateKey);
	if err != nil {
		log.Fatalf("Failed to recover public key: %v", err)
	}

	return sig, user;
}


func CreatTypedData(verifyingContract string, chainId int64, user string, nonce int64) (apitypes.TypedData) {
	typeddata := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Registration": []apitypes.Type{
				{Name: "user", Type: "address"},
				{Name: "nonce", Type: "uint256"},
			},
		},
		PrimaryType: "Registration",
		Domain: apitypes.TypedDataDomain{
			Name:              "Registration",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(chainId),
			VerifyingContract: verifyingContract,
		},
		Message: apitypes.TypedDataMessage{
			"user":            user,
			"nonce":		   math.NewHexOrDecimal256(nonce),
		},
	}
	return typeddata
}

func signTypedData(typedData apitypes.TypedData, privateKey *ecdsa.PrivateKey) (sig []byte , err error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return sig, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return sig, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := common.BytesToHash(crypto.Keccak256(rawData))
	sig, err = crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return sig, err
	}

	sig[64] += 27
	return
}

func VerifySignature(config types.Config, signature string, userPrivateKey string) (common.Address) {
	sig, userAddress := CreateDataAndSign(config, userPrivateKey);

	if hexutil.Encode(sig) != signature {
		log.Fatal("Signature verification failed");
	}

	return userAddress;
}
