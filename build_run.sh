#!/usr/bin/env bash

export dockerImgTag=cascade_proxy_mubeng
export dockerCtnName=cascade_proxy_mubeng

docker build --tag=${dockerImgTag} .
docker rm -f ${dockerCtnName}
docker run -dit --name=${dockerCtnName} --network=host ${dockerImgTag}
