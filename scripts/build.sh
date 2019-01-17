#!/bin/bash

PROJECT_NAME="pinjur-lunch"
PROJECT_DIRECTORY="src/github.com/djukela17/$PROJECT_NAME"

APP_NAME="pinjur-lunch"

echo $GOPATH/$PROJECT_DIRECTORY
cd $GOPATH/$PROJECT_DIRECTORY

# create cmd directory if it is not created
if [ ! -d cmd ] ; then
    mkdir cmd
fi

go build -o cmd/$APP_NAME main.go

if [[ ! -d cmd/web ]]; then
    mkdir cmd/web
fi

if [[ ! -d cmd/web/templates ]]; then
    echo cmd/templates does not exist. Creating ...
    mkdir cmd/web/templates
fi


if [[ ! -d cmd/data ]]; then
    echo cmd/data does not exist. Creating ...
    mkdir cmd/data
fi

cp -r web/templates cmd/web/templates
cp -r data cmd/data


