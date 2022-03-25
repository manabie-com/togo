#!/bin/bash
set -e

./tools/volume.sh &
cd src
go build -o master
./master