FROM golang:1.11-alpine3.8 as builder

RUN apk add --no-cache make git musl-dev gcc

WORKDIR /github.com/envoyproxy/envoy/examples/grpc-bridge/client
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
RUN make mod-download

COPY . .
RUN make && make install

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /github.com/envoyproxy/envoy/examples/grpc-bridge/client/bin/grpc-bridge-client /app/bin/grpc-bridge-client