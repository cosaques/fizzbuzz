FROM golang:1.10.2 AS build

ADD . /go/src/project
WORKDIR /go/src/project
RUN go get project
RUN go install
EXPOSE 8080
ENTRYPOINT ["/go/bin/project"]