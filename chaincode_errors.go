package main

import "errors"

var (
	ERR_INVALID_FUNC        = errors.New("Invalid function")
	ERR_INVALID_ARG_COUNT   = errors.New("Invalid number of input arguments")
	ERR_INVALID_CERT_ID     = errors.New("Invalid certificate id")
	ERR_INVALID_CERT_DATA   = errors.New("Invalid certificate data")
	ERR_INVALID_DOMAIN_NAME = errors.New("Invalid domain name")
	ERR_CERT_ALREADY_EXISTS = errors.New("Certificate already exists")
)
