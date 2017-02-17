package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/base64"
	"github.com/psi-phi/identity-chaincode/identity"
	"github.com/golang/protobuf/proto"
	"encoding/json"
)

/* ------------------------------------------------------------------------------------------------------------------
   KEY HELPERS
   ------------------------------------------------------------------------------------------------------------------ */
func getOrgKey(domainName string) string {
	return KEY_PREFIX_ORGANIZATION + domainName
}

func getCertKey(certId string) string {
	return KEY_PREFIX_CERTIFICATE + certId
}


/* ------------------------------------------------------------------------------------------------------------------
   VERSION
   ------------------------------------------------------------------------------------------------------------------ */
func setVersion(stub shim.ChaincodeStubInterface) ([]byte, error) {
	logger.Println("Chaincode initialized. Version = " + VERSION)
	err := stub.PutState(KEY_VERSION, []byte(VERSION))
	if err != nil {
		logger.Printf("[ERROR] Failed to set version : %s\n", err)
		return nil, err
	}
	return nil, nil
}

func getVersion(stub shim.ChaincodeStubInterface) ([]byte, error) {
	val, err := stub.GetState(KEY_VERSION)
	if err != nil {
		logger.Printf("[ERROR] Failed to get version : %s\n", err)
		return nil, err
	}
	return val, nil
}


/* ------------------------------------------------------------------------------------------------------------------
   ORGANIZATION CERTIFICATES
   ------------------------------------------------------------------------------------------------------------------ */
func putOrganizationCertificate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, ERR_INVALID_ARG_COUNT
	}

	certData, err := base64.StdEncoding.DecodeString(args[0])
	if err != nil {
		return nil, err
	}

	cert := new(identity.SignedCertificate)
	err = proto.Unmarshal(certData, cert)
	if err != nil {
		return nil, err
	}

	orgKey := getOrgKey(cert.Data.DomainName)
	certKey := getCertKey(cert.Id)

	orgData, err := stub.GetState(orgKey)
	if err != nil {
		return nil, err
	}

	// TODO :  Validate Certificate

	org := new(identity.Organization)
	if orgData == nil {
		/* Create new organization */
		org.DomainName = cert.Data.DomainName
		org.ActiveCerts = []*identity.SignedCertificate{cert}

	} else {
		/* Add certificate to existing organization */
		err = proto.Unmarshal(orgData, org)
		if err != nil {
			return nil, err
		}
		org.ActiveCerts = append(org.ActiveCerts, cert)
	}

	// Serialize organization
	updatedOrgData, err := proto.Marshal(org)
	if err != nil {
		return nil, err
	}

	// Save organization data
	err = stub.PutState(orgKey, updatedOrgData)
	if err != nil {
		return nil, err
	}

	// Save certificate mapping
	err = stub.PutState(certKey, []byte(orgKey))
	if err != nil {
		if orgData == nil {
			stub.DelState(orgKey)
		} else {
			err2 := stub.PutState(orgKey, orgData)
			if err2 != nil {
				return nil, err2
			}
		}
		return nil, err
	}

	return nil, nil
}

func getOrganizationCertificate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, ERR_INVALID_ARG_COUNT
	}

	certId := args[0]
	certKey := getCertKey(certId)

	orgKeyData, err := stub.GetState(certKey)
	if err != nil {
		return nil, err
	}

	if orgKeyData == nil {
		return nil, ERR_INVALID_CERT_ID
	}
	orgKey := string(orgKeyData)

	orgData, err := stub.GetState(orgKey)
	if err != nil {
		return nil, err
	}

	org := new(identity.Organization)
	err = proto.Unmarshal(orgData, org)
	if err != nil {
		return nil, err
	}

	for _, cert := range org.ActiveCerts {
		if certId == cert.Id {
			certData, err := json.Marshal(cert)
			if err != nil {
				return nil, err
			}
			return certData, nil
		}
	}

	for _, cert := range org.BlockedCerts {
		if certId == cert.Id {
			certData, err := json.Marshal(cert)
			if err != nil {
				return nil, err
			}
			return certData, nil
		}
	}

	for _, cert := range org.BlockedCerts {
		if certId == cert.Id {
			certData, err := json.Marshal(cert)
			if err != nil {
				return nil, err
			}
			return certData, nil
		}
	}

	return nil, nil
}