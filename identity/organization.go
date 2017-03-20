package identity

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	validator "github.com/asaskevich/govalidator"
	"net"
	"strings"
)

//var logger = log.New(os.Stdout, "", log.Ldate | log.Ltime | log.LUTC)

func (org *Organization) FindCert(certId string) *SignedCertificate {
	for _, cert := range org.ActiveCerts {
		if certId == cert.Id {
			return cert
		}
	}

	for _, cert := range org.BlockedCerts {
		if certId == cert.Id {
			return cert
		}
	}

	for _, cert := range org.BlockedCerts {
		if certId == cert.Id {
			return cert
		}
	}

	return nil
}

func (cert *SignedCertificate) IsValid() bool {

	// Validate domain name
	if !validator.IsHost(cert.Data.DomainName) {
		return false
	}

	// Validate org data
	if cert.Data.OrgDetails.Name == "" {
		return false
	}

	if !validator.IsEmail(cert.Data.OrgDetails.Email) {
		return false
	}

	if !validator.IsISO3166Alpha2(cert.Data.OrgDetails.Country) {
		return false
	}

	// Validate public key
	publicKey, err := x509.ParsePKIXPublicKey(cert.Data.PublicKey.Data)
	if err != nil {
		return false
	}

	// Validate self sign
	msg := cert.Id
	hashedData := sha256.Sum256([]byte(msg))
	signature := cert.SelfSign.Data

	switch cert.Data.PublicKey.Algorithm {
	case PublicKey_RSA:
		pk := publicKey.(*rsa.PublicKey)
		err = rsa.VerifyPKCS1v15(pk, crypto.SHA256, hashedData[:], signature)
		if err != nil {
			return false
		}

	case PublicKey_ECDSA:
		// TODO :  Add ECDSA support
		return false

	default:
		return false
	}

	// Auto Validate
	autoValidationMechanism := cert.AutoValidation
	if autoValidationMechanism == AutoValidationMechanism_DNS {
		dnsRecords := fetchDnsRecords(cert.Data.DomainName)
		isValid := false
		for _, dnsRecord := range dnsRecords {
			if dnsRecord == cert.Id {
				isValid = true
			}
		}

		if !isValid {
			return false
		}
	}

	return true
}

func (cert *SignedCertificate) VerifySignature(base64EncodedMsgHash string, sign *Signature) bool {
	msgHash, err := base64.StdEncoding.DecodeString(base64EncodedMsgHash)
	if err != nil {
		return false
	}
	signData := sign.Data
	publicKey, err := x509.ParsePKIXPublicKey(cert.Data.PublicKey.Data)
	if err != nil {
		return false
	}

	switch cert.Data.PublicKey.Algorithm {
	case PublicKey_RSA:
		pk := publicKey.(*rsa.PublicKey)
		err = rsa.VerifyPKCS1v15(pk, crypto.SHA256, msgHash, signData)
		if err != nil {
			return false
		}
		return true

	default:
		return false
	}
}

/* Helper Methods */
func fetchDnsRecords(domainName string) []string {
	var dnsRecords []string
	records, _ := net.LookupTXT("identity_certs." + domainName)
	for _, record := range records {
		tokens := strings.Split(record, " ")
		for _, token := range tokens {
			dnsRecords = append(dnsRecords, token)
		}
	}
	return dnsRecords
}
