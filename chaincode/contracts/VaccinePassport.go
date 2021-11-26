package contracts

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SimplePaymentContract contract for handling writing and reading from the world state
type VaccinePassport struct {
	contractapi.Contract
}

// Account : The asset being tracked on the chain
type Passport struct {
	// Vaccine details
	PassportID    string `json:"passportID"`
	DocNumber     string `json:"docNumber"`
	Name          string `json:"name"`
	DOB           string `json:"dob"`
	VaccineType   string `json:"vaccineType"`
	DateOfDose1   string `json:"dateOfDose1"`
	DateOfDose2   string `json:"dateOfDose2"`
	VaccineStatus string `json:"vaccineStatus"`
}

var passportIDCounter int64

// InitLedger : Init the ledger
func (spc *VaccinePassport) InitLedger(ctx contractapi.TransactionContextInterface) error {
	passportIDCounter = 0
	return nil
}

// VaccineDetails : Enter vaccine details of a user to be reviewed by approver
func (spc *VaccinePassport) VaccineDetails(ctx contractapi.TransactionContextInterface, docnumber string, name string,
	dob string, vaccinetype string, dateofdose1 string, dateofdose2 string) (*Passport, error) {

	// check if there is already a record for the document submitted
	exists, err := spc.AssetExists(ctx, docnumber)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("the asset %s already exists", docnumber)
	}

	// increment the global variable to get new passportID
	passportIDCounter += 1
	// attach the ID as prefix to the counter and get the passport number
	id := "ID" + strconv.FormatInt(passportIDCounter, 10)

	passport := Passport{
		PassportID:    id,
		DocNumber:     docnumber,
		Name:          name,
		DOB:           dob,
		VaccineType:   vaccinetype,
		DateOfDose1:   dateofdose1,
		DateOfDose2:   dateofdose2,
		VaccineStatus: "pending",
	}

	vaccineBytes, err := json.Marshal(passport)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(id, vaccineBytes)
	if err != nil {
		return nil, err
	}

	return &passport, nil
}

// AssetExists returns true when asset with given ID exists in world state
func (spc *VaccinePassport) AssetExists(ctx contractapi.TransactionContextInterface, docnumber string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(docnumber)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// UpdateStatus : Update the status of passport once review is done
func (spc *VaccinePassport) UpdateStatus(ctx contractapi.TransactionContextInterface, passportid string, status string) (*Passport, error) {

	return nil, nil
}
