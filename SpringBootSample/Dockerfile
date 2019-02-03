FROM maven:3.6.0-jdk-8 as builder

USER root
RUN mkdir /app
WORKDIR /app

COPY pom.xml .
RUN mvn verify clean --fail-never

COPY src ./src
RUN mvn package

##############################################33
FROM fabric8/java-centos-openjdk8-jre:latest as runner

USER root

RUN ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

ENV JAVA_APP_JAR app.jar
ENV JAVA_APP_DIR /opt/app

COPY --from=builder /app/target/*.jar /opt/app/app.jar

EXPOSE 8080
