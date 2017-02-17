#!/usr/bin/env bash

rm common.proto
rm organization.proto
rm gen_go.sh
wget https://raw.githubusercontent.com/psi-phi/identity-protobuf/master/common.proto
wget https://raw.githubusercontent.com/psi-phi/identity-protobuf/master/organization.proto
wget https://raw.githubusercontent.com/psi-phi/identity-protobuf/master/gen_go.sh
chmod +x gen_go.sh
PROTO_GO_OUT_DIR=../identity ./gen_go.sh