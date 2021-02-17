#!/usr/bin/env bash

docker build -t capstone:v1 .

docker image ls

docker run -p 8080:8080 -d --name "capstone_webserver" capstone:v1