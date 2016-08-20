#!/bin/bash

rm -rf bin/
go build -o bin/hiupdate
chmod +x bin/hiupdate
cp -r static bin/

# docker build -t hiupdate .
# docker run --rm -p 8080:8080 hiupdate
