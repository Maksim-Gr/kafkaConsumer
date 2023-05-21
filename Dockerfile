FROM golang:1.20-alpine as builder

WORKDIR /go/src/app

ADD . .

RUN apk --no-cache add git make build-base && \
        go get -d -v ./... && \
        go install -v ./... && \
        CGO_ENABLED=0 GOOS=linux go build -o run


FROM aplpine:3

COPY --from=builder /go/src/app/run /app/run

CMD ["/app/run"]