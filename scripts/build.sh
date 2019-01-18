#!/bin/bash

PROJECT_NAME="pinjur-lunch"
PROJECT_DIRECTORY="src/github.com/djukela17/$PROJECT_NAME"

APP_NAME="pinjur-lunch"

cd $GOPATH/$PROJECT_DIRECTORY

go build -o cmd/$APP_NAME main.go
