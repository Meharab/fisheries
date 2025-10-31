package main

import (
    "log"
    "os"

    "github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
    "github.com/hyperledger/fabric-chaincode-go/v2/shim"

    // replace below with module path found in chaincode/go.mod:
    sc "github.com/Meharab/fisheries/tree/main/chaincode" //"github.com/hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/chaincode"
)

func main() {
    // chaincode instance
    contract := contractapi.NewChaincode(&sc.SmartContract{})

    // get CCID from env or use default
    ccid := os.Getenv("CHAINCODE_ID")
    if ccid == "" {
        ccid = "fisheries"
    }

    server := &shim.ChaincodeServer{
        CCID:    ccid,
        Address: "0.0.0.0:7052",
        CC:      contract,
    }

    log.Printf("starting chaincode server (CCID=%s) on 0.0.0.0:7052", ccid)
    if err := server.Start(); err != nil {
        log.Fatalf("chaincode server start failed: %v", err)
    }
}
