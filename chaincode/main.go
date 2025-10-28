/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	// "github.com/hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/chaincode"
	"github.com/hyperledger/fabric-chaincode-go/shim"
    sc "github.com/meharab/fisheries/chaincode/chaincode"
)

func main() {
	// assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	// if err != nil {
	// 	log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	// }

	// if err := assetChaincode.Start(); err != nil {
	// 	log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	// }

	contract := contractapi.NewChaincode(&sc.SmartContract{})

    server := &shim.ChaincodeServer{
        CCID:   "fisheries",          // logical id; match --name later
        Address: "0.0.0.0:7052",      // bind inside container
        CC:      contract,
    }

    log.Println("Starting Fisheries chaincode server on 0.0.0.0:7052")
    if err := server.Start(); err != nil {
        log.Fatalf("Failed to start chaincode server: %v", err)
    }
}
