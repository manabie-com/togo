# Start with a base image containing Java runtime https://hub.docker.com/r/adoptopenjdk/openjdk11
FROM gradle:6.8.3-jdk11 as builder
#Possibility to set JVM options (https://www.oracle.com/technetwork/java/javase/tech/vmoptions-jsp-140102.html)
#Set app home folder
ENV APP_HOME=/home/root/build/
WORKDIR $APP_HOME
COPY . .
RUN ./gradlew clean build -x test

# Start with a base image containing Java runtime
FROM openjdk:11-slim as runtime
WORKDIR /home/root/todo/
COPY --from=builder /home/root/build/build/libs/todo_api-0.1.0.jar $WORKDIR
EXPOSE 8080
ENTRYPOINT [ "sh", "-c", "java -jar $JAVA_OPTS -Djava.security.egd=file:/dev/./urandom todo_api-0.1.0.jar" ]
