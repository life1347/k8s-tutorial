FROM golang:onbuild

WORKDIR /go
COPY *.go /go/

RUN GOOS=linux GOARCH=386 go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -o server .

FROM alpine:3.6

WORKDIR /app

COPY --from=0 /go/server /app/server

EXPOSE 8888
ENTRYPOINT ["./server"]
