#!/bin/bash
#
#PROJECT_NAME="pinjur-lunch"
#PROJECT_DIRECTORY="src/github.com/djukela17/$PROJECT_NAME"
#
#APP_NAME="pinjur-lunch"
#
#cd $GOPATH/$PROJECT_DIRECTORY
#
#if [[ ! -z "$GOPATH" ]]; then
#
#    cd ${GOPATH}/${PROJECT_DIRECTORY}
#
#    if [[ "$1" = "" ]]; then
#        echo No version specified. Building with default name.
#        go build -o cmd/${APP_NAME} main.go
#    else
#        echo version name specified
#        go build -o cmd/${APP_NAME}-$1 main.go
#    fi
#
#else
#    echo no \$GOPATH in \$PATH
#fi