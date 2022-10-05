#!/bin/sh

set -e
set -u
# set -x


SERVICE_NAME=$1
if [ X"${SERVICE_NAME}" = X"" ]
then
  SERVICE_NAME="webservice"
fi


echo ""
echo "service name: ${SERVICE_NAME}" 
echo "build start..."


# go get -u -v
# go mod tidy


echo " - download the necessary packages"
# Install dependencies
go mod download


echo " - verify the packages"
# Verify dependencies
go mod verify


echo " - build the service"
CGO_ENABLED=0 go build -o ${SERVICE_NAME}
echo " - build success"
echo ""
