package main

// Version
const VERSION = "0.9"

// Function Names
const (
	FUNC_GET_VERSION      = "get_version"
	FUNC_GET_CERT         = "get_cert"
	FUNC_GET_PRIMARY_CERT = "get_primary_cert"
	FUNC_VERIFY_SIGN      = "verify_sign"

	FUNC_PUT_CERT       = "put_cert"
	FUNC_ADD_VALIDATION = "add_validation"
	FUNC_EXPIRE_CERT    = "expire_cert"
	FUNC_BLOCK_CERT     = "block_cert"
)

// Keys
const (
	KEY_VERSION             = "version"
	KEY_PREFIX_ORGANIZATION = "org::"
	KEY_PREFIX_CERTIFICATE  = "cert::"
)
