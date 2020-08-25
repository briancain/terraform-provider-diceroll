TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=briancain
NAME=diceroll
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}
