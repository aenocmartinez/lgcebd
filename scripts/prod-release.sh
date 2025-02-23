#!/bin/bash

docker build --no-cache -t pulzo/redaccion:latest -f Dockerfile.prod . --platform linux/amd64

docker push pulzo/redaccion:latest

docker rmi pulzo/redaccion:latest