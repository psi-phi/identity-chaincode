package main

import (
	"bitbucket.org/psi-phi/two-factor-auth-chaincode/lib"
	"encoding/base64"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/psi-phi/identity-chaincode/identity"
	"strconv"
)

/* ------------------------------------------------------------------------------------------------------------------
   HELPERS
   ------------------------------------------------------------------------------------------------------------------ */
func getOrgKey(domainName string) string {
	return KEY_PREFIX_ORGANIZATION + domainName
}

func getCertKey(certId string) string {
	return KEY_PREFIX_CERTIFICATE + certId
}

func getOrganizationFromKey(stub shim.ChaincodeStubInterface, orgKey string) (*identity.Organization, error) {
	orgData, err := stub.GetState(orgKey)
	if err != nil {
		return nil, err
	}

	if orgData == nil {
		return nil, nil
	} else {
		org := new(identity.Organization)
		err = proto.Unmarshal(orgData, org)
		if err != nil {
			return nil, err
		}
		return org, nil
	}
}

func getOrganizationFromDomain(stub shim.ChaincodeStubInterface, domainName string) (*identity.Organization, error) {
	orgKey := getOrgKey(domainName)
	return getOrganizationFromKey(stub, orgKey)
}

func getCertificateFromId(stub shim.ChaincodeStubInterface, certId string) (*identity.SignedCertificate, error) {
	certKey := getCertKey(certId)

	orgKeyData, err := stub.GetState(certKey)
	if err != nil {
		return nil, err
	}

	if orgKeyData == nil {
		return nil, ERR_INVALID_CERT_ID
	}
	orgKey := string(orgKeyData)
	org, err := getOrganizationFromKey(stub, orgKey)
	if err != nil {
		return nil, err
	}
	if org == nil {
		return nil, ERR_INVALID_CERT_ID
	}

	return org.FindCert(certId), nil
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
   CERTIFICATES
   ------------------------------------------------------------------------------------------------------------------ */
func putCertificate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// Parse input
	if len(args) < 1 {
		return nil, ERR_INVALID_ARG_COUNT
	}

	certData, err := base64.StdEncoding.DecodeString(args[0])
	if err != nil {
		return nil, err
	}

	setAsPrimary := false
	if len(args) >= 2 {
		setAsPrimary, err = strconv.ParseBool(args[1])
		if err != nil {
			setAsPrimary = false
		}
	}

	// Create certificate from serialized input
	cert := new(identity.SignedCertificate)
	err = proto.Unmarshal(certData, cert)
	if err != nil {
		return nil, err
	}

	// Validate certificate
	if !cert.IsValid() {
		return nil, ERR_INVALID_CERT_DATA
	}

	// Fetch organization
	org, err := getOrganizationFromDomain(stub, cert.Data.DomainName)
	if err != nil {
		return nil, err
	}

	// Add certificate in organization's active certificate list
	if org == nil {
		/* Create new organization */
		org = new(identity.Organization)
		org.DomainName = cert.Data.DomainName
		org.PrimaryCertId = cert.Id
		org.ActiveCerts = []*identity.SignedCertificate{cert}

	} else {
		/* Add certificate to existing organization */
		alreadyExists := false
		for _, c := range org.ActiveCerts {
			if c.Id == cert.Id {
				alreadyExists = true
				break
			}
		}
		if alreadyExists {
			return nil, ERR_CERT_ALREADY_EXISTS
		}

		org.ActiveCerts = append(org.ActiveCerts, cert)

		orgJson, _ := json.Marshal(org)
		logger.Println(string(orgJson))

		// Set as Primary certificate
		if setAsPrimary || org.PrimaryCertId == "" {
			org.PrimaryCertId = cert.Id
		}
	}

	// Serialize organization
	updatedOrgData, err := proto.Marshal(org)
	if err != nil {
		return nil, err
	}

	// Save organization data
	orgKey := getOrgKey(org.DomainName)
	err = stub.PutState(orgKey, updatedOrgData)
	if err != nil {
		return nil, err
	}

	// Save certificate mapping
	certKey := getCertKey(cert.Id)
	err = stub.PutState(certKey, []byte(orgKey))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func getCertificate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, ERR_INVALID_ARG_COUNT
	}

	certId := args[0]
	cert, err := getCertificateFromId(stub, certId)
	if err != nil {
		return nil, err
	}

	if cert != nil {
		certData, err := proto.Marshal(cert)
		if err != nil {
			return nil, err
		}
		encodedCertData := base64.StdEncoding.EncodeToString(certData)
		return []byte(encodedCertData), nil
	}

	return nil, nil
}

func getCeritificateStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, ERR_INVALID_ARG_COUNT
	}

	certId := args[0]
	cert, err := getCertificateFromId(stub, certId)
	if err != nil {
		return nil, err
	}

	if cert != nil {
		status := cert.Status.String()
		return []byte(status), nil
	}

	return nil, nil
}

func getRole(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, ERR_INVALID_ARG_COUNT
	}

	certId := args[0]
	cert, err := getCertificateFromId(stub, certId)
	if err != nil {
		return nil, err
	}

	if cert != nil {
		role := cert.Data.Role.String()
		return []byte(role), nil
	}

	return nil, nil
}

func addValidation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// TODO : Implement this
	return nil, nil
}

func expireCertificate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// TODO : Implement this
	return nil, nil
}

func blockCertificate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// TODO : Implement this
	return nil, nil
}

func getPrimaryCertificate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, ERR_INVALID_ARG_COUNT
	}

	domainName := args[0]
	org, err := getOrganizationFromDomain(stub, domainName)
	if err != nil {
		return nil, err
	}
	if org == nil {
		return nil, ERR_INVALID_DOMAIN_NAME
	}

	cert := org.FindCert(org.PrimaryCertId)
	if cert != nil {
		certData, err := proto.Marshal(cert)
		if err != nil {
			return nil, err
		}
		encodedCertData := base64.StdEncoding.EncodeToString(certData)
		return []byte(encodedCertData), nil
	}

	return nil, ERR_INVALID_CERT_ID
}

func verifySignature(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// Parse Input
	if len(args) != 2 {
		return nil, lib.ERR_INCORRECT_ARGS
	}

	encodedMsgHash := args[0]

	encodedSignData := args[1]
	signData, err := base64.StdEncoding.DecodeString(encodedSignData)
	if err != nil {
		return nil, err
	}

	// Create signature from serialized input
	sign := new(identity.Signature)
	err = proto.Unmarshal(signData, sign)
	if err != nil {
		return nil, err
	}

	// Get certificate corresponding to this signature
	signerCert, err := getCertificateFromId(stub, sign.SignerCertId)
	if err != nil {
		return nil, err
	}

	// Verify signature
	isValid := signerCert.VerifySignature(encodedMsgHash, sign)
	if isValid {
		return []byte("true"), nil
	} else {
		return []byte("false"), nil
	}
}
