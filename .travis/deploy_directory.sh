#!/bin/bash
docker login -u $DOCKER_USER -p $DOCKER_PASS
docker build -f Dockerfile -t shadowshotx/jsoncliops .
docker push shadowshotx/jsoncliops:latest
