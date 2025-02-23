#!/bin/bash

docker build --no-cache -t pulzo/redaccion_back:qa -f Dockerfile.prod . --platform linux/amd64

docker push pulzo/redaccion_back:qa

docker rmi pulzo/redaccion_back:qa