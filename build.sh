#!/bin/bash

rm -rf bin/
go build -o bin/startup-update
chmod +x bin/startup-update
cp -r static bin/

