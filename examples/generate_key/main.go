package main

import (
	"fmt"

	"github.com/CaioLuColaco/hedera-sdk-local"
)

func main() {
	// Generating key
	privateKey, err := hedera.GeneratePrivateKey()
	if err != nil {
		panic(fmt.Sprintf("%v : error generating PrivateKey", err))
	}

	// Retrieve the public key
	publicKey := privateKey.PublicKey()

	fmt.Printf("private = %v\n", privateKey)
	fmt.Printf("public = %v\n", publicKey)
}
