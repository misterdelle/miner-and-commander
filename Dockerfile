FROM golang:1.22.3-alpine3.19 as builder

WORKDIR /build

COPY . .

RUN go mod download && CGO_ENABLED=0

RUN go build -ldflags "-s -w" -o miner-and-commander

FROM alpine:3.17

WORKDIR /

RUN apk upgrade --no-cache --ignore alpine-baselayout --available && \
    apk --no-cache add ca-certificates tzdata && \
    rm -rf /var/cache/apk/* && \
    mkdir -p /app && \
    mkdir -p /app/templates

COPY --from=builder /build/miner-and-commander /app/miner-and-commander
COPY --from=builder /build/templates/* /app/templates
COPY --from=builder /build/.env /app
RUN chmod +x /app/miner-and-commander

ENTRYPOINT ["/app/miner-and-commander"]
