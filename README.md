# GoEdu - educational reference project

This is educational and reference project created to demonstrate how can we implement **business applications** in Go.

## Table of contents

[1. Introduction](#1-introduction)

&nbsp;&nbsp;[1.1 Purpose of this Repository](#11-purpose-of-this-repository)

&nbsp;&nbsp;[1.2 Introducing Problem Space](#12-introducing-problem-space)

[2. How to Run](#2-how-to-run)

&nbsp;&nbsp;[2.1 Running the application from scratch](#21-running-the-application-from-scratch)

&nbsp;&nbsp;[2.2 Make Binary](#22-make-binary)

&nbsp;&nbsp;[2.3 Build Docker image](#23-build-docker-image)

&nbsp;&nbsp;[2.4 Run Tests](#24-run-tests)

&nbsp;&nbsp;[2.5 Tools Used](#25-tools-used)

[3. Inspirations](#2-inspirations)

## 1. Introduction

![Hexagonal Architecture](https://github.com/nalkhovikleverx/GoEdu/blob/master/docs/technical/images/hexagon.jpg?raw=true)

### 1.1 Purpose of this Repository

This is a list of the main goals of this repository:

- Get experience of implementing a monolith in a modular way in Go
- Applying software engineering *best practices* (Clean Architecture, Design Patterns, Testable Design etc)
- Keeping design artifacts alongside with the source code: **C4 Model**, **OpenAPI spec**, **ADR**

### 1.2 Introducing Problem Space

Problem domain is disscussed [here](https://github.com/nalkhovikleverx/GoEdu/discussions/18)

## 2. How to Run

### 2.1 Running the application from scratch:

```bash
make run
```

### 2.2 Make binary

```bash
make build
```

### 2.3 Build Docker image

```bash
make docker-image
```

### 2.4 Run tests

```bash
make test
```

### 2.5 Tools Used

- [golangci-lint](https://golangci-lint.run/)
- [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) to check the dependency rule
- [plantUML](https://github.com/plantuml/plantuml)
- [markdownlint](https://github.com/DavidAnson/markdownlint) to lint ADR
- [Redocly](https://redocly.com/docs/cli/installation/#install-redocly-cli) to bundle and lint OpenAPI
- [oapi-codegen](https://github.com/deepmap/oapi-codegen) to generate HTTP server from API spec
- Docker

## 3. Inspirations and Recommendations

- [Modular Monolith Reference by Kamil Grzybek](https://github.com/kgrzybek/modular-monolith-with-ddd)
- [EventStore reference app](https://github.com/EventStore/training-advanced-go)
- [Wild-Workouts by ThreeDots](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example)
- [VaughnVernon Implementing DDD reference](https://github.com/VaughnVernon/IDDD_Samples)
- [Event-driven Architecture in Go book reference](https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang)
- [eShop reference application by .NET team](https://github.com/dotnet/eShop)
- ["How modular can your monolith go" series by Chris Richardson](https://microservices.io/post/architecture/2023/07/31/how-modular-can-your-monolith-go-part-1.html)
