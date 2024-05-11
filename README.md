# GoEdu - educational reference project

This is educational and reference project created to demonstrate how can we implement **business applications** in Go.

## Table of contents

[1. Introduction](#1-introduction)

&nbsp;&nbsp;[1.1 Purpose of this Repository](#11-purpose-of-this-repository)

[2. How to Run](#2-how-to-run)

&nbsp;&nbsp;[2.1 Tools Used](#21-tools-used)

[3. Inspirations](#2-inspirations)

## 1. Introduction

![Hexagonal Architecture](https://github.com/nalkhovikleverx/GoEdu/blob/master/docs/images/hexagon.jpg?raw=true)

### 1.1 Purpose of this Repository

This is a list of the main goals of this repository:

- Get experience of implementing a monolith in a modular way in Go
- Applying software engineering *best practices* (Clean Architecture, Design Patterns, Testable Design etc)
- Keeping design artifacts alongside with the source code: **C4 Model**, **diagram as code**, **ADR**

## 2. How to Run

For now only you can do is run the tests:

```bash
go test -v ./...
```

### 2.1 Tools Used

- [golangci-lint](https://golangci-lint.run/)
- [go-cleanarch](https://github.com/roblaszczak/go-cleanarch)
- [plantUML](https://github.com/plantuml/plantuml)
- [markdownlint (for ADR)](https://github.com/DavidAnson/markdownlint)

## 3. Inspirations and Recommendations

- [Modular Monolith Reference by Kamil Grzybek](https://github.com/kgrzybek/modular-monolith-with-ddd)
- [EventStore reference app](https://github.com/EventStore/training-advanced-go)
- [Wild-Workouts by ThreeDots](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example)
- [VaughnVernon Implementing DDD reference](https://github.com/VaughnVernon/IDDD_Samples)
- [Event-driven Architecture in Go book reference](https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang)
- [eShop reference application by .NET team](https://github.com/dotnet/eShop)
- ["How modular can your monolith go" series by Chris Richardson](https://microservices.io/post/architecture/2023/07/31/how-modular-can-your-monolith-go-part-1.html)
