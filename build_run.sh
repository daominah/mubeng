#!/usr/bin/env bash

export dockerImgTag=rotator_mubeng
export dockerCtnName=rotator_mubeng

docker build --tag=${dockerImgTag} .
docker rm -f ${dockerCtnName}
docker run -dit --name=${dockerCtnName} --network=host ${dockerImgTag}
