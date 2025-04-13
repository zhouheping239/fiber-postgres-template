FROM golang:1.23.0-alpine

RUN apk add --no-cache --upgrade bash ca-certificates openjdk21 && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ENV JAVA_HOME /usr/lib/jvm/java-21-openjdk
ENV PATH $PATH:$JAVA_HOME/bin

WORKDIR /backend_project

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

ARG PROJECT_NAME="backend_project"
ENV PROJECT_NAME ${PROJECT_NAME}

RUN go build -o ${PROJECT_NAME} ./cmd

RUN sed -i 's/\r$//' /backend_project/run.sh
RUN chmod +x /backend_project/run.sh

EXPOSE 8080
