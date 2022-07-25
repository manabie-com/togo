FROM openjdk:11
VOLUME /tmp
ADD target/manabie-togo.jar manabie-togo.jar
EXPOSE 8008
ENTRYPOINT ["java","-jar","manabie-togo.jar"]