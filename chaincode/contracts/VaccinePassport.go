package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SimplePaymentContract contract for handling writing and reading from the world state
type VaccinePassport struct {
	contractapi.Contract
}

// Account : The asset being tracked on the chain
type Passport struct {
	// Vaccine details

	DocNumber     string `json:"docNumber"`
	Name          string `json:"name"`
	DOB           string `json:"dob"`
	VaccineType   string `json:"vaccineType"`
	DateOfDose1   string `json:"dateOfDose1"`
	DateOfDose2   string `json:"dateOfDose2"`
	Comments      string `json:"comments"`
	VaccineStatus string `json:"vaccineStatus"`
}

// InitLedger : Init the ledger
func (spc *VaccinePassport) InitLedger(ctx contractapi.TransactionContextInterface) error {

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

	passport := Passport{

		DocNumber:     docnumber,
		Name:          name,
		DOB:           dob,
		VaccineType:   vaccinetype,
		DateOfDose1:   dateofdose1,
		DateOfDose2:   dateofdose2,
		Comments:      "documents uploaded",
		VaccineStatus: "pending",
	}

	vaccineBytes, err := json.Marshal(passport)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(docnumber, vaccineBytes)
	if err != nil {
		return nil, err
	}

	return &passport, nil
}

// AssetExists returns true when asset with given ID exists in world state
func (spc *VaccinePassport) AssetExists(ctx contractapi.TransactionContextInterface, docnumber string) (bool, error) {
	fmt.Println("Inside AsseExists")
	assetJSON, err := ctx.GetStub().GetState(docnumber)
	fmt.Println(assetJSON)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// fetch details of the given passport-id
func (spc *VaccinePassport) PassportDetails(ctx contractapi.TransactionContextInterface, docnumber string) (*Passport, error) {

	passportBytes, err := ctx.GetStub().GetState(docnumber)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	if passportBytes == nil {
		return nil, fmt.Errorf("no details exists for the passport id: %v", docnumber)
	}

	var passport Passport
	err = json.Unmarshal(passportBytes, &passport)
	if err != nil {
		return nil, err
	}

	return &passport, nil
}

// UpdateStatus : Update the status of passport once review is done
func (spc *VaccinePassport) UpdateStatus(ctx contractapi.TransactionContextInterface, docnumber string, status string,
	comments string) (*Passport, error) {

	passportBytes, err := ctx.GetStub().GetState(docnumber)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	if passportBytes == nil {
		return nil, fmt.Errorf("no details exists for the passport id: %v", docnumber)
	}

	var passport Passport
	err = json.Unmarshal(passportBytes, &passport)
	if err != nil {
		return nil, err
	}

	// update the status and comments to the record fetched for the passport id
	passport.VaccineStatus = status
	passport.Comments = passport.Comments + "\n" + comments

	// Marshal the passport variable and make it ready to put on ledger
	vaccineBytes, err := json.Marshal(passport)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(docnumber, vaccineBytes)
	if err != nil {
		return nil, err
	}
	return &passport, nil
}
