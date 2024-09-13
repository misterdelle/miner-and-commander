FROM golang:1.22.3-alpine3.20 AS builder

WORKDIR /build

COPY . .

#RUN go mod download && CGO_ENABLED=0
RUN go mod download

RUN go build -ldflags "-s -w" -o miner-and-commander

FROM alpine:3.20

WORKDIR /

RUN apk upgrade --ignore alpine-baselayout --no-check-certificate
RUN apk --no-cache add ca-certificates --no-check-certificate
RUN apk --no-cache add tzdata --no-check-certificate
RUN rm -rf /var/cache/apk/*

#RUN apk upgrade \
#    --no-cache \
#    --ignore alpine-baselayout \
#    --no-check-certificate \
#    --available && \
#    apk --no-cache add ca-certificates tzdata && \
#    rm -rf /var/cache/apk/*

RUN mkdir -p /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /build/miner-and-commander /app/miner-and-commander
COPY --from=builder /build/.env /app
RUN chmod +x /app/miner-and-commander

ENV TZ="Europe/Rome"

ENTRYPOINT ["/app/miner-and-commander"]
