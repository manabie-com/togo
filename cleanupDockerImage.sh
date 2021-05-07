#!/bin/bash

name=$1
if [ -z "${name}" ]
then
name="none"
fi

docker images | grep -ie ${name} | awk '{print $3}' | xargs docker rmi -f
