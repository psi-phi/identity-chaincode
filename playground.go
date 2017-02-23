package main

//
//import (
//	"encoding/base64"
//	"crypto/x509"
//	"fmt"
//	"crypto"
//	"crypto/sha256"
//	"crypto/rsa"
//	"github.com/psi-phi/identity-chaincode/identity"
//	"time"
//	"github.com/golang/protobuf/proto"
//)
//
//func main() {
//	pkStr := "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAKWU3SFMIuxiEG4HJ04goxOM6Y1xp6RJTET2SRuRYX4d6xgr2enohzIfmfNCU5TS3AdV1nYGa7up/yuFnkmgiGlwmldjzk3RJkOnLcq+cX0OalfWaD4W8SESNMD2LGeuTlfrJG6neV7/rHS07MPpNxrgOYYIjTAXL4eCg4ODzY4tAgMBAAECgYAQ8gorZTJRxLtvtzWzji2CS7J/Mjl427N9f0L+GkPC92be42X6xTxIyekkdw48tdOkwIkhLQkLfJtLpfIMEOzEy2bQv6447Z31pGJFuLjXq+ISFKYdzOGeKMaFw5WYHooEt4HC1MaIYnAMFiMmBIOcFCyfd3S7IstHWgA6SBO3QQJBAN69pOPMdPm2Wu6+vUrMDYB20VKbPwfqqeJbtPXqFLyBqW/hdvyR3o2H0waiIMxZ5MFI2yehEl3K22zaduEeqd0CQQC+Tkhe/ff+X7cp86+h4rhkPghqOE+E6eAZQ45h64nw6piqucZe7xdeFn0vJDHqiIeevdcJhtS5kXomFtAWcziRAkAFfgei1lfMEIMNgAaK4Z0znbprnwhe2Zp2ymwb5Dm+rDPRXm3grHggZUj+0OCeKVlKqtE8mOwrA+WFOZ3UzzzBAkBH8QIM3weMIxT0CApCMZoxUv4NYaI2Bc/Q3SgLPmaUK6txBu/F3a7Aw9GpK46vMdPLH8sV7+GlESjTE1aw1ffxAkEAhESDpHka2rFFpPIIhOnXnxNMM36LnEhcqy3ln4v3Fz6P/vv0uBtrp7+EYYuPbajb7j2l5L9be9IKswFyEWfjjQ=="
//	pkData, _ := base64.StdEncoding.DecodeString(pkStr)
//	pk, _ := x509.ParsePKCS8PrivateKey(pkData)
//	privK := pk.(*rsa.PrivateKey)
//
//	pubKStr := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCllN0hTCLsYhBuBydOIKMTjOmNcaekSUxE9kkbkWF+HesYK9np6IcyH5nzQlOU0twHVdZ2Bmu7qf8rhZ5JoIhpcJpXY85N0SZDpy3KvnF9DmpX1mg+FvEhEjTA9ixnrk5X6yRup3le/6x0tOzD6Tca4DmGCI0wFy+HgoODg82OLQIDAQAB"
//	pubKData, _ := base64.StdEncoding.DecodeString(pubKStr)
//	pubK0,_ := x509.ParsePKIXPublicKey(pubKData)
//	_ = pubK0.(*rsa.PublicKey)
//
//	msg := "hello"
//	hash := sha256.Sum256([]byte(msg))
//	encodedHash := base64.StdEncoding.EncodeToString(hash[:])
//	fmt.Println(encodedHash)
//
//	s, _ := rsa.SignPKCS1v15(nil, privK, crypto.SHA256, hash[:])
//
//	t := time.Now().Unix()
//
//	signature := identity.Signature{
//		SignerCertId:"1xpA25jlvK920O5dCQVS4zlzBwx861sxuwpYWfOPuTU",
//		Data:s,
//		Timestamp:uint64(t),
//	}
//
//	data, _ := proto.Marshal(&signature)
//	sign := base64.StdEncoding.EncodeToString(data)
//	fmt.Println(sign)
//
//
//	//s1,_ := base64.StdEncoding.DecodeString("Anp51+M1fZ5HVU+WnD6jKSJNdOnr4WQRg/DXlMVouLW0dF1f9hxYYuBympZSzRj9El+oNwJ8hnC8dr/Cf/i7d35kvSoV0Ez1389oU3/H47xYyLNr+asaZ4MN/jfrKa0sc2pXXx4m32AJdrtR5AH0ddb/umlYGMgt67XElA00uHU=")
//	//fmt.Println(rsa.VerifyPKCS1v15(pubK, crypto.SHA256, hash[:], s1))
//}
