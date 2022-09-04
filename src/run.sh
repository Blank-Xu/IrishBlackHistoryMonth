#!/bin/bash

PROJECT_NAME="webservice"


echo "---- update package"
go get -u -v


echo "---- run go mod tidy"
go mod tidy


echo "---- run test"
go test -v .


echo "---- build application"
go build -o ${PROJECT_NAME}


echo "---- start application"
./${PROJECT_NAME}
