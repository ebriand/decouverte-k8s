FROM golang:1.11-alpine3.7 as builder

COPY server.go /go

RUN CGO_ENABLED=0 GOOS=linux go build server.go
RUN mkdir /go/static
COPY demo1.png /go/static/demo1.png

FROM scratch

EXPOSE 80

COPY --from=builder /go/server /server
COPY --from=builder /go/static /static

ENTRYPOINT [ "/server" ]
