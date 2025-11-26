# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

GoEdu is an educational reference project demonstrating how to build business applications in Go using:
- **Modular Monolith Architecture** with clear module boundaries
- **Clean Architecture / Hexagonal Architecture** within each module
- **Domain-Driven Design (DDD)** principles
- **CQRS (Command Query Responsibility Segregation)** pattern

**Problem Domain**: Online Consultation Marketplace - connecting consultants with clients for consultations, handling registration, authentication, session scheduling, and search.

## Essential Development Commands

```bash
# Development workflow
make run          # Generate OpenAPI, lint, and run application (complete workflow)
make build        # Build binary to bin/monolith (includes OpenAPI gen and lint)
make test         # Run all tests
make lint         # Run golangci-lint with strict rules

# OpenAPI workflow
make openapi      # Bundle OpenAPI specs, lint, and generate server code
                  # Generated: internal/api/http/server/server.gen.go

# Additional commands
make generate     # Run go generate ./...
make docker-image # Build Docker image (multi-stage, scratch-based)
make adr-lint     # Lint Architecture Decision Records
make plantuml     # Generate diagrams from PlantUML source
```

## Architecture & Design Patterns

### Modular Monolith Structure

Each business capability is a separate module with strict boundaries:

```
internal/<module-name>/
├── api/
│   └── inprocess/           # PUBLIC API - Module facade interface (contracts)
├── internal/                # PRIVATE - Implementation (Go's internal visibility)
│   ├── domain/             # Domain entities, value objects, domain events
│   ├── application/        # Command handlers, use cases, service interfaces
│   ├── infrastructure/     # Repository implementations, external adapters
│   └── interfaces/         # Facade implementations, mappers
├── module.go               # Module initialization (Root function for DI)
└── doc.go                  # Package documentation
```

**Key principle**: Modules communicate through `api/inprocess` facades (interface-based), never directly accessing another module's internals. See ADR-0001, ADR-0005.

### Clean Architecture Dependency Rules

```
Domain (innermost)
  ↑
Application (defines repository/service interfaces)
  ↑
Infrastructure (implements interfaces)
  ↑
Interfaces (adapts external APIs to internal use cases)
```

- **Domain layer**: No dependencies, pure business logic
- **Application layer**: Depends only on domain, defines interfaces for external concerns
- **Infrastructure**: Implements application/domain interfaces (repositories, external services)
- **Interfaces**: Maps between internal and external representations

Validated by `go-cleanarch` linter.

### CQRS Implementation (ADR-0004)

All write operations use the Command pattern:

```go
// Command definition
type RegisterNewUserCommand struct {
    Email    string
    Password string
    // ...
}

// Command result (pragmatic CQRS: commands can return results)
type RegisterNewUserCommandResult struct {
    UserRegistrationID string
}

// Command handler
type RegisterNewUserCommandHandler struct {
    repository Repository
    hasher     PasswordHasher
}

func (h *CommandHandler) Handle(ctx context.Context, cmd Command) (CommandResult, error) {
    // Implementation
}
```

Read operations use separate query models (not yet fully implemented in current modules).

### Module Communication Pattern

1. **Define public API** in `api/inprocess/interface.go`:
```go
type RegistrationModuleFacade interface {
    RegisterNewUser(context.Context, RegisterNewUserCommand) (RegisterNewUserCommandResult, error)
}
```

2. **Implement facade** in `internal/interfaces/inprocess/facade.go`:
```go
type Facade struct {
    commandHandler *application.CommandHandler
}

func (f *Facade) RegisterNewUser(ctx context.Context, cmd RegisterNewUserCommand) (RegisterNewUserCommandResult, error) {
    // Adapt public command to internal command and delegate
}
```

3. **Register in module.go**:
```go
func Root(ctx context.Context, dep module.Dependencies) error {
    // Wire up dependencies
    repo := memory.NewRepository()
    handler := application.NewCommandHandler(repo, ...)
    facade := inprocess.NewFacade(handler)

    // Expose via dependencies
    dep.SetRegistrationModuleAPI(facade)
    return nil
}
```

4. **Use from other modules** via `module.Dependencies`:
```go
dep.GetRegistrationModuleAPI().RegisterNewUser(ctx, cmd)
```

## Creating a New Module

1. **Create module structure** following the pattern above
2. **Implement Module interface** in `module.go`:
```go
type Module struct{}

func (m *Module) Init(ctx context.Context, dep module.Dependencies) error {
    return Root(ctx, dep)
}
```

3. **Register in `cmd/monolith/main.go`**:
```go
modules: []module.Module{
    &registration.Module{},
    &yournewmodule.Module{},  // Add here
    &httpapi.Module{},        // HTTP API must be last
}
```

4. **Add methods to `module.Dependencies`** interface for inter-module communication

## Domain Layer Patterns

### Value Objects
Immutable, self-validating domain primitives:

```go
// Constructor validates and returns value object or error
func NewUserEmail(email string) (UserEmail, error) {
    // Validation logic
    return UserEmail{value: email}, nil
}

// Or panics for must-succeed scenarios
func MustNewUserEmail(email string) UserEmail {
    // Used when email is already validated
}
```

Examples: `UserEmail`, `UserPassword`, `UserName`, `HashedPassword`

### Entities
Identity-based objects with mutable state:

```go
type UserRegistration struct {
    id             UserRegistrationID
    email          UserEmail
    status         RegistrationStatus
    // ...
}

// Use snapshot pattern for persistence
func (u *UserRegistration) Snapshot() UserRegistrationSnapshot {
    return UserRegistrationSnapshot{ /* fields */ }
}
```

### Domain Events
Track state changes, defined in `events.go`:

```go
type UserRegisteredEvent struct {
    UserRegistrationID string
    Email              string
    OccurredAt         time.Time
}
```

## Testing Approach

- **Location**: Test files alongside production code (`*_test.go`)
- **Pattern**: Table-driven tests with descriptive names
- **Isolation**: Use test doubles (mocks, spies, stubs) for dependencies
- **Library**: `github.com/stretchr/testify` for assertions

Example structure:
```go
func TestPositiveScenario(t *testing.T) {
    tests := map[string]struct {
        dependency1 *MockDependency
        dependency2 *SpyDependency
        command     Command
        expected    Result
    }{
        "descriptive test case name": { /* setup */ },
    }

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            // Arrange, Act, Assert
        })
    }
}
```

## Code Quality & Conventions

### Linter Configuration (`.golangci.yml`)
Strict rules enforced via `make lint`:
- **Function complexity**: Max 60 lines, 40 statements
- **Cognitive complexity**: Max 20
- **Cyclomatic complexity**: Max 10
- **No global variables** (except approved cases)
- **No init functions**
- **No naked returns**
- **Strict error checking** (errcheck, goerr113)

### Naming Conventions
- Commands: `<Verb><Noun>Command` (e.g., `RegisterNewUserCommand`)
- Results: `<Command>Result`
- Handlers: `<Command>Handler`
- Repositories: `<Entity>Repository` interface (in application layer)
- Facades: `<Module>ModuleFacade`

### Package Documentation
Every package must have `doc.go` with package-level documentation.

### Error Handling
- Define domain-specific errors in domain package
- Use sentinel errors: `var ErrNotFound = errors.New("...")`
- Wrap errors with context when propagating
- Application layer defines error types for use cases

## Configuration & Environment

- **Config management**: `github.com/kelseyhightower/envconfig`
- **Local development**: `.env.local` file (loaded via `ENVIRONMENT` variable)
- **Key settings**:
  - `SERVICE_NAME` (default: "noname")
  - `ENVIRONMENT` (default: "dev")
  - `LOG_LEVEL` (default: "DEBUG")
  - `HTTP_HOST` / `HTTP_PORT` (default: "0.0.0.0:8080")
  - `SHUTDOWN_TIMEOUT` (default: 30s)

Configuration struct location: `internal/pkg/config/config.go`

## HTTP API Development

- **OpenAPI-first**: Define API in `api/openapi/v3.0/spec.yaml`
- **Code generation**: Server interfaces generated via `oapi-codegen`
  - Generated file: `internal/api/http/server/server.gen.go`
  - Implement `ServerInterface` in `internal/api/http/handlers/`
- **Bundled spec**: `api/openapi/dist/monolith.openapi.yaml` (for distribution)

After modifying OpenAPI spec:
1. Run `make openapi` to regenerate server code
2. Implement new handlers
3. Register in HTTP API module

## Docker

Multi-stage Dockerfile at `docker/Dockerfile`:
- **Build stage**: `golang:1.22-alpine` with dependency caching
- **Runtime stage**: `scratch` (minimal image, ~10MB)
- **Binary**: Built with `-ldflags="-s -w"` (stripped)
- **Exposed port**: 8000 (note: app defaults to 8080, configure via env)

Build: `make docker-image` → tags as `goedu`

## Important Documentation

- **ADRs**: `docs/decisions/` - Read these to understand architectural decisions
  - ADR-0001: Modular Monolith rationale
  - ADR-0004: CQRS implementation
  - ADR-0005: Domain API module pattern
- **C4 Diagrams**: `docs/C4/` - Visual architecture (PlantUML)
- **README.md**: Problem domain and learning objectives

## Key Dependencies

- `github.com/google/uuid` - UUID generation
- `github.com/kelseyhightower/envconfig` - Environment-based config
- `github.com/deepmap/oapi-codegen/v2` - OpenAPI server generation
- `github.com/stretchr/testify` - Testing assertions
- `go.opentelemetry.io/otel` - Observability (tracing)

## Current Modules

1. **Registration** (`internal/registration/`)
   - User registration workflow
   - Email uniqueness verification
   - Registration confirmation

2. **User Access** (`internal/useraccess/`)
   - Authentication and login
   - Password verification
   - User session management

3. **HTTP API** (`internal/api/http/`)
   - REST API layer
   - OpenAPI-generated server
   - Request/response mapping

## Graceful Shutdown

Uses `internal/pkg/waiter` package for coordinated shutdown:
- Catches OS signals (SIGINT, SIGTERM)
- Respects `SHUTDOWN_TIMEOUT` configuration
- Allows services to clean up resources

Example in modules:
```go
waiter.Add(func(ctx context.Context) error {
    // Cleanup logic
    return service.Shutdown(ctx)
})
```

## Technology Stack Summary

- **Language**: Go 1.22+
- **Architecture**: Modular Monolith + Clean Architecture + DDD + CQRS
- **HTTP**: Standard library + OpenAPI code generation
- **Testing**: `go test` + `testify`
- **Linting**: `golangci-lint` (50+ enabled linters)
- **Observability**: OpenTelemetry (traces)
- **Logging**: `log/slog` (structured logging)
