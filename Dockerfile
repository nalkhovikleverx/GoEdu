FROM golang:1.22-alpine as build-env
LABEL authors="aliaksei.kosyrau"
WORKDIR /app
COPY . .
RUN apk --update add ca-certificates git
RUN go mod download
RUN go build -o bin/server cmd/server/main.go
RUN go build -o bin/client cmd/client/main.go


FROM scratch

WORKDIR /app

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-env /app/bin /app/bin


EXPOSE 8000
ENTRYPOINT ["bin/server"]