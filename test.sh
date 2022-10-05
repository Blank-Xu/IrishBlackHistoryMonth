#!/bin/sh

set -e
set -u
# set -x


echo "test start..."
go test -v  -cover ./...
