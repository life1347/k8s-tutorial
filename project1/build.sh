#!/bin/sh


#THIS_DIR=$(realpath $(dirname $0))
#docker run -v $THIS_DIR:/proj microsoft/dotnet:2.0.0-sdk sh -c "cd /proj && ./project-build.sh"

FROM golang:onbuild

WORKDIR /go
COPY *.go /go/

#RUN GOOS=linux GOARCH=386 go build -o server .
GOOS=linux GOARCH=linux go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -o server .

FROM alpine:3.6

WORKDIR /app

#RUN apk update
#RUN apk add coreutils binutils findutils grep

COPY --from=0 /go/server /app/server

EXPOSE 8888
ENTRYPOINT ["./server"]
