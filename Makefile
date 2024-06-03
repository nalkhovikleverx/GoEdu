.PHONY: lint adr-lint plantuml generate test openapi build run docker-image

lint:
	golangci-lint run

adr-lint:
	markdownlint-cli2 --config ./docs/decisions/.markdownlint.yml docs/decisions/*.md

plantuml:
	find ./docs/C4/ -name "*.puml" -type f -exec plantuml {} +

generate:
	go generate ./...

test:
	go test ./...

openapi:
	redocly lint --config ./api/openapi/redocly.yaml monolith@v1
	redocly bundle --output ./api/openapi/dist/monolith.openapi.yaml --config ./api/openapi/redocly.yaml monolith@v1 --ext yaml
	@$(MAKE) generate

build: openapi lint
	go build -o bin/monolith cmd/monolith/main.go

run: openapi lint
	go run cmd/monolith/main.go

docker-image: openapi
	docker build -t goedu -f docker/Dockerfile .
