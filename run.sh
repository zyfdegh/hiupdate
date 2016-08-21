#!/bin/sh

docker build -t hiupdate .
docker run --rm -p 8080:8080 -v /tmp/hiupdate:/usr/local/bin hiupdate

