package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	contract := new(contract)
	contract.Info.Version = "0.0.1"
	contract.Info.Description = "Smart contract"

	chaincode, err := contractapi.NewChaincode(contract)
	chaincode.Info.Title = "Sankofa Patient Contract"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from contract" + err.Error())

	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode" + err.Error())
	}

}
