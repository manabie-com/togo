FROM maven:3.8.5-openjdk-17 as builder
WORKDIR /todo

COPY ./pom.xml .
COPY ./src/ ./src/

RUN --mount=type=cache,target=/root/.m2 ["mvn", "-Djar.finalName=todo", "package"]

FROM openjdk:17.0-slim-buster
WORKDIR /todo

COPY --from=builder /todo/target/*.jar ./todo.jar

EXPOSE 8080
ENTRYPOINT ["java", "-jar", "todo.jar"]