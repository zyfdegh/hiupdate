#!/bin/bash

rm -rf bin/
go build -o bin/hiupdate
chmod +x bin/hiupdate
cp -r static bin/
cp hiupdate.conf bin/

