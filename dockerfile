FROM golang:1.18beta2-buster

LABEL Name="grpc-"
LABEL Version="1.0.0"

ENV GOPRIVATE=xxx
ENV GITHUB_USER_NAME=GITHUB-user-name
ENV GITHUB_TOKEN=GITHUB-token
ENV GITHUB_EMAIL=GITHUB-email

WORKDIR /
COPY ./init.sh init.sh

RUN mkdir -p /repo
VOLUME /go
VOLUME /repo
WORKDIR /repo

EXPOSE 8080
EXPOSE 8081
CMD ["bash", "/init.sh"]
