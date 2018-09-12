FROM golang:1.11-alpine3.7

RUN apk add --no-cache curl

COPY server.go /go

RUN go build server.go

ENTRYPOINT [ "/go/server" ]
