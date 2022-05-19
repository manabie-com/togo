#!/bin/bash
rm migrations/migrate.test
touch migrations/migrate.test
docker-compose -f docker-compose.test.yaml up