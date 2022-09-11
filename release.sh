#!/bin/bash

PROJECT_NAME="webservice"
VERSION_NO=$(date +"%Y%m%d%H%M")
RELEASE_FOLDER=${PROJECT_NAME}"_release"
RELEASE_NAME=${PROJECT_NAME}_${VERSION_NO}


echo "---- start release the project: ${PROJECT_NAME}"
echo "---- release version name: ${RELEASE_NAME}"

echo "---- mkdir ${RELEASE_FOLDER}"
mkdir ${RELEASE_FOLDER}


echo "---- copy necessary files"
cp -rf backup ${RELEASE_FOLDER}/
cp -rf configs ${RELEASE_FOLDER}/
cp -rf logs ${RELEASE_FOLDER}/


echo "---- start build project ${PROJECT_NAME}"
cd src/

echo "---- run go mod tidy"
go mod tidy

echo "---- build project"
go build -o ../${RELEASE_FOLDER}/${PROJECT_NAME} ${PROJECT_NAME}

cd ../


echo "---- package the realease version folder"
tar -zcvf ${RELEASE_NAME}.tar.gz ${RELEASE_FOLDER}

echo "---- remove the release folder ${RELEASE_FOLDER}"
rm -rf ${RELEASE_FOLDER}


echo "---- finished"
echo "${RELEASE_NAME}.tar.gz"
