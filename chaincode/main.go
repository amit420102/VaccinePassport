package main

import (
	"vaccine-passport-application-chaincode/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	vaccinePassport := new(contracts.VaccinePassport)

	cc, err := contractapi.NewChaincode(vaccinePassport)

	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}
