#!/usr/bin/env bash

go clean
go build
CORE_CHAINCODE_ID_NAME=identity CORE_PEER_ADDRESS=0.0.0.0:7051 ./identity-chaincode