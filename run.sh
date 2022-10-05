#!/bin/sh

set -e
set -u
# set -x


SERVICE_NAME="webservice"


sh build.sh ${SERVICE_NAME}


echo "start service..."
echo ""

./${SERVICE_NAME}
