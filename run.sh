#!/bin/sh

sudo docker build -t hiupdate .
sudo docker run --rm -p 8080:8080 hiupdate

