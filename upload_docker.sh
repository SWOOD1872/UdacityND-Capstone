#!/usr/bin/env bash

dockerpath=sam72/capstone:v1

echo "Docker ID and Image: $dockerpath"
docker login
docker tag capstone:v1 $dockerpath

docker push $dockerpath