package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC)

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		logger.Printf("[ERROR] Failed to start chaincode : %s\n", err)
	}
}

type Chaincode struct {
}

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return setVersion(stub)
}

func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case FUNC_PUT_CERT:
		return putOrganizationCertificate(stub, args)
	default:
		return nil, ERR_INVALID_FUNC
	}
}

func (t *Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case FUNC_GET_VERSION:
		return getVersion(stub)
	case FUNC_GET_CERT:
		return getOrganizationCertificate(stub, args)
	default:
		return nil, ERR_INVALID_FUNC
	}
}
