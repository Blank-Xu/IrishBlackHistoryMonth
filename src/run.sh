#!/bin/bash

PROJECT_NAME="webservice"

bash build.sh

echo "start service..."
echo ""

./${SERVICE_NAME}
