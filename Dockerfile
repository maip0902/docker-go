FROM golang:latest

RUN apt-get update && \
    apt-get -y install vim
RUN mkdir /go/src/project_server
WORKDIR /go/src/project_server
RUN go mod init
RUN go get gopkg.in/mgo.v2
RUN go get github.com/99designs/gqlgen
RUN go run github.com/99designs/gqlgen init
ADD ./app /go/src/project_server
WORKDIR /go/src/project_server/project
RUN go build -o ../build hello.go
CMD ["/go/src/project_server/build/hello"]