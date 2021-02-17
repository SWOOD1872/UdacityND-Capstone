#!/usr/bin/env bash

dockerpath=sam72/capstone:v1

kubectl run capstone --image=$dockerpath 

kubectl get pods

sleep 3

kubectl port-forward capstone 8080:8080