package contracts

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SimplePaymentContract contract for handling writing and reading from the world state
type VaccinePassport struct {
	contractapi.Contract
}

// Account : The asset being tracked on the chain
type Passport struct {
	// Account details
}

// InitLedger : Init the ledger
func (spc *VaccinePassport) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}

// RegisterUserAccount : User registers his account
func (spc *VaccinePassport) RegisterUserAccount(ctx contractapi.TransactionContextInterface, name string, bank string) (*Passport, error) {
	// your Register logic
	return nil, nil
}

// Balance : to check the senders balance
func (spc *VaccinePassport) Balance(ctx contractapi.TransactionContextInterface) (int64, error) {
	// your balance logic

	return 0, nil
}

// Transfer : to transfer amount and update balances
func (spc *VaccinePassport) Transfer(ctx contractapi.TransactionContextInterface, beneficiary string, amount int64) (string, error) {
	// your Tranfer logic
	return "", nil
}
