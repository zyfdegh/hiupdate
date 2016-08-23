#!/bin/sh

sudo docker build -t hiupdate .
sudo docker run --rm -p 8080:8080 -v /tmp/db:/usr/local/bin/db hiupdate
