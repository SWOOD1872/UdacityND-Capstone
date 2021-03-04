[![SWOOD1872](https://circleci.com/gh/SWOOD1872/UdacityND-Capstone.svg?style=svg)](https://app.circleci.com/pipelines/github/SWOOD1872/UdacityND-Capstone)


# UdacityND-Capstone
Udacity AWS Cloud DevOps Engineer Nanodegree - Capstone Project

## Overview

My Capstone project is a simple web server written in Go, which serves a basic web page with some styling. I used the [gorilla/mux router](https://github.com/gorilla/mux) for routing rather than just the standard library. I've included a few 'extra' features as well, namely graceful shutdown, custom filesystem to prevent directory listings and usage of the 'go:embed' directive to embed static assets in to the binary.

The Dockerfile is a multi-stage docker build, in order to reduce the size of the image.

## Testing

You can run the unit tests for this project as follows

```bash
go clean -testcache ./... && \
go test -v -race ./...
```

## Running Locally

You can run the web server locally by compiling the code and running the binary/executable

```bash
go build -o server .

./server
```
## Running in Docker

You can run the web server inside Docker

```bash
./run_docker.sh
```

## Running in Kubernetes

You can also run inside Kubernetes on your own machine, using Minikube for example

```bash
./run_kubernetes.sh
```

## Pipeline

The CICD tool of choice is CircleCI. As such, any commits will trigger the pipeline.

Any commits on the main branch, will also trigger a deployment to the AWS Cloud. This will perform a rolling deployment to a Kubernetes cluster running in Amazon EKS.

### Main Pipeline Steps:

1. Lint the Dockerfile and the Go code
2. Vet the Go code
3. Compile the Go code
4. Run the unit tests
5. Build the Docker image
6. Tag the Docker image and push to Docker Hub
7. Create an Amazon EKS cluster with Kuberneres version 1.18
8. Apply the Kubernetes deployment and service to the cluster
9. Perform a rolling deployment/restart

Build artifacts including test results, load balancer url and output from 'kubectl describe all' are all stored in the CircleCI artifacts tab of each build.