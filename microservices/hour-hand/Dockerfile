FROM golang:1.15 as builder

WORKDIR /go/src/github.com/kzmake/dapr-clock

ENV GO111MODULE=on

RUN groupadd -g 10001 docker && \
    useradd -u 10001 -g docker docker

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY constants constants
COPY microservices/hour-hand microservices/hour-hand

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app github.com/kzmake/dapr-clock/microservices/hour-hand

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 3000

USER docker

CMD ["/go/bin/app"]
