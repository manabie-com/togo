ARG DOCKER_POSTGRES_IMAGE_TAG="alpine"


FROM  postgres:${DOCKER_POSTGRES_IMAGE_TAG}
COPY sql/* /docker-entrypoint-initdb.d/