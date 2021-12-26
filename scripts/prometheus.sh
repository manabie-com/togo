#!/bin/bash

docker run --add-host host.docker.internal:host-gateway -p 9094:9090 -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus 
