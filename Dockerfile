FROM golang:1.21-alpine3.17 as builder

WORKDIR /build

COPY . .

RUN go mod download && CGO_ENABLED=0

RUN go build -ldflags "-s -w" -o miner-and-commander

FROM alpine:3.17

WORKDIR /

RUN apk upgrade --no-cache --ignore alpine-baselayout --available && \
    apk --no-cache add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

COPY --from=builder /build/miner-and-commander .
RUN chmod +x miner-and-commander

ENTRYPOINT ["/miner-and-commander"]
