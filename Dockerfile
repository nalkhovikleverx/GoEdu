FROM golang:1.22-alpine as build-env
LABEL authors="aliaksei.kosyrau"
WORKDIR /app
ENV CGO_ENABLED=0
COPY . .
RUN apk --update add ca-certificates git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /server cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /client cmd/client/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-env /app .

EXPOSE 8000
ENTRYPOINT ["/server"]