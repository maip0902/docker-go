FROM golang:latest

RUN apt-get update && \
    apt-get -y install vim
RUN mkdir /go/src/project_server
WORKDIR /go/src/project_server
RUN go mod init
RUN go get github.com/99designs/gqlgen
RUN go run github.com/99designs/gqlgen init
ADD . /go/src/project_server
CMD ["go", "run", "hello.go"]