FROM maven:3.8.5-openjdk-17 as builder
WORKDIR /todo

COPY ./pom.xml .
COPY ./src/ ./src/

RUN ["mvn", "install"]


ENTRYPOINT ["mvn", "test"]