FROM golang:1.22-alpine as build-env
LABEL authors="aliaksei.kosyrau"
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /server cmd/server/main.go
RUN go build -o /client cmd/client/main.go

FROM alpine:latest

WORKDIR /

COPY --from=build-env /server /server
COPY --from=build-env /client /client

ENTRYPOINT ["/server"]