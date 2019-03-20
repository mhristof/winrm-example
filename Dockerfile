FROM golang
MAINTAINER Mike Christof <mhristof@gmail.com>

COPY foobar.go /go/src/foobar/foobar.go
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN cd /go/src/foobar &&\
        dep init &&\
        go build
