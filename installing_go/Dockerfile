FROM ubuntu:18.04

RUN apt-get update -y
RUN apt-get -y install curl vim git iputils-ping telnet net-tools

RUN mkdir -p /go_workspace

RUN mkdir /go_workspace/bin
RUN mkdir -p /go_workspace/src/github.com/ederavilaprado/golang/installing_go

COPY ./hello_world.go /go_workspace/src/github.com/ederavilaprado/golang/installing_go

RUN curl -O https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
RUN tar -xvf go1.10.2.linux-amd64.tar.gz
RUN mv go /usr/local
RUN rm /go1.10.2.linux-amd64.tar.gz

ENV GOPATH "$HOME/go_workspace"
ENV PATH "$PATH:/usr/local/go/bin:$GOPATH/bin"

RUN go get golang.org/x/tools/cmd/goimports

WORKDIR /go_workspace/src/github.com/ederavilaprado/golang/installing_go