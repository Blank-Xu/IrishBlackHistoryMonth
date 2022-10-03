#!/bin/bash

PROJECT_NAME="webservice"
VERSION_NO=$(date +"%Y%m%d%H%M")
RELEASE_FOLDER="release"
RELEASE_NAME=${PROJECT_NAME}_${VERSION_NO}


echo "-- start release the project: ${PROJECT_NAME}"
echo "   release version name: ${RELEASE_NAME}"
# mkdir ${RELEASE_FOLDER}


echo "-- start build project ${PROJECT_NAME}"
bash build.sh
mv -f ${PROJECT_NAME} ${RELEASE_FOLDER}/


echo "-- package the realease version folder"
tar -zcvf ${RELEASE_NAME}.tar.gz ${RELEASE_FOLDER}

echo ""
echo "-- release package finished"
echo "${RELEASE_NAME}.tar.gz"
