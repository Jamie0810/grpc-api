FROM golang:1.14 as builder


WORKDIR /kbc2
COPY . /kbc2

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go build -ldflags="-X config.versionTag=$(git describe --tags)" -mod=vendor -v -o captain_marvel main.go
# RUN go build -mod=vendor -v -o captain_marvel 

FROM alpine:3

RUN apk add --no-cache ca-certificates

RUN mkdir -p /kbc2
WORKDIR /kbc2

COPY --from=builder /kbc2/captain_marvel /kbc2/captain_marvel
COPY --from=builder /kbc2/internal/pkg/config /kbc2/internal/pkg/config

# ENTRYPOINT ./captain_marvel
CMD ./captain_marvel grpc