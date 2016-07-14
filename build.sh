#!/bin/bash

rm -rf bin/
go build -o bin/startup-update main.go
cp -r static bin/