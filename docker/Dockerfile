FROM golang:1.22-alpine as build-env

WORKDIR /app

COPY go.* ./
RUN go mod download

# Build the binary.
COPY .. ./
RUN go build -ldflags="-s -w" -v -o /app/bin/monolith ./cmd/monolith

FROM scratch

WORKDIR /app

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-env /app/bin /app/bin

ENTRYPOINT ["/app/bin/monolith"]
