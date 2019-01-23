#!/bin/bash
#
PROJECT_NAME="pinjur-lunch"
APP_NAME="pinjur"

if [[ ! -z "$GOPATH" ]]; then

    # $1 [1st arg} is the destination project name
    if [[ ! "$1" = "" ]]; then
        echo destination project directory specified
        echo $1
        PROJECT_NAME=$1
    fi

    DEST_DIR="${GOPATH}/bin/${PROJECT_NAME}"
    go build -o ${DEST_DIR}/${APP_NAME} cmd/pinjur/main.go
    # copy required files
    cp -r web ${DEST_DIR}/web
    cp -r data ${DEST_DIR}/data
    cp -r build ${DEST_DIR}/build
    cp -r deployment ${DEST_DIR}/deployment

else
    echo no \$GOPATH in \$PATH
fi