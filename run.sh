#!/bin/sh

image="zyfdedh/hiupdate:dev"

sudo docker build -t ${image} .
sudo docker run --rm -p 8080:8080 -v /tmp/db:/usr/local/bin/db ${image}
