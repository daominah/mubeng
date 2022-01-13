#!/usr/bin/env bash

export dockerImgTag=cascade_proxy_mubeng
export dockerCtnName=cascade_proxy_mubeng
export dockerCtnName2=cascade_proxy_mubeng_ipv6

docker build --tag=${dockerImgTag} .
docker rm -f ${dockerCtnName}
docker run -dit --name=${dockerCtnName} --restart=always \
    --network=host \
    ${dockerImgTag}

docker rm -f ${dockerCtnName2}
docker run -dit --name=${dockerCtnName2} --restart=always \
    --network=host \
    ${dockerImgTag} \
    /bin/mubeng -a 127.0.0.1:23529 -f /proxies_ipv6.txt
