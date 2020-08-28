TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=briancain
NAME=diceroll
BINARY=terraform-provider-${NAME}
VERSION=0.1.1
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}

install: build
		mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
		mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
