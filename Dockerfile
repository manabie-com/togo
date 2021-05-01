# Build with $ docker build -t ubuntu/todosrv:latest .
# Start with docker compose up
# Stop with docker-compose kill

FROM ubuntu:18.04
ENV DEBIAN_FRONTEND noninteractive
USER root

CMD ["/bin/bash"]

RUN apt-get clean \ 
    && apt-get update \
    && apt-get install -y --no-install-recommends \
    sudo \
    apt-transport-https \
    curl \
#    golang \
    git \
	wget \ 
	unzip \
	gpg \
	gpg-agent \
	ca-certificates
	# redis-server

ADD ./scripts/install_postgres.sh /root
RUN chmod +x ~/install_postgres.sh
RUN ~/install_postgres.sh

ADD ./scripts/docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]
WORKDIR /root