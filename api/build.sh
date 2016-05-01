#!/bin/bash

echo building binary
env GOOS=linux GOARCH=amd64 go build -v

VERSION=`git rev-parse --short HEAD`
echo building docker image waieez/fbforbots-api:$VERSION
docker build -t waieez/fbforbots-api:$VERSION .
docker tag waieez/fbforbots-api:$VERSION waieez/fbforbots-api:latest

mv api ../bin