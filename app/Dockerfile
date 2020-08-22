FROM golang:latest

RUN apt-get update && \
    apt-get -y install vim
RUN mkdir /go/src/project_server
WORKDIR /go/src/project_server
ADD . /go/src/project_server
CMD ["go", "run", "hello.go"]