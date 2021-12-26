#!/bin/bash

docker run --add-host host.docker.internal:host-gateway -p 8083:8080 -e BASE_URL=/ -e SWAGGER_JSON=/swagger/togo_public_api.swagger.json -v $(pwd):/swagger swaggerapi/swagger-ui
