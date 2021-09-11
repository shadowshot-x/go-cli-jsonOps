#!/bin/sh
docker login -u $DOCKER_USER -p $DOCKER_PASS
docker push shadowshotx/jsoncliops:latest
