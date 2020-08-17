FROM golang:latest
RUN mkdir /go/src/project_server
WORKDIR /go/src/project_server
ADD . /go/src/project_server
CMD ["go", "run", "hello.go"]