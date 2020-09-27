#!/bin/sh

# Get root path
root=$(dirname "${0}")/..
# Compile frontend
if test "${1}" = "-a"
then
    sh "${root}"/scripts/front2back.sh || exit 1
fi
# Run the server
go run "${root}"/main.go