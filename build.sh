#!/bin/bash

SERVICE_NAME="webservice"


echo ""
echo "build start..."

# go get -u -v


echo " - download the necessary packages"
# Install dependencies
go mod download


echo " - verify the packages"
# Verify dependencies
go mod verify

# go mod tidy


echo " - build the service"
go build -o ${SERVICE_NAME}
echo " - build success"
echo ""
