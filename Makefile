TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=briancain
NAME=diceroll
BINARY=terraform-provider-${NAME}
VERSION_FILE=version.txt
VERSION:=$(shell cat $(VERSION_FILE))
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}

install: build
		mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
		mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
