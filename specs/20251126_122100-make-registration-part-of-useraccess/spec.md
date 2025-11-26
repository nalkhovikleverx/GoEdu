# The Feature: make a `registration` module part of `useraccess` module

## Task Context

During initial decomposition we decided to keep these two modules separately. However, this initial decomposition can be a problem in the future due to these
modules are tightly connected with each other.
Right now we want to keep things as easy as possible, so we were decided to make registration module a part of useraccess.

## Detailed Task Description and Rules

### Documentation

Currently the C4 component diagram shows clear separation between those two modules. The adjustment is to remove "registration" module from it.
@docs/technical/C4/c3-components.puml
To regenerate images use `make plantuml' command

### OpenAPI spec

The @api/openapi/v3.0/resources/registrations needs to become a part of
@api/openapi/v3.0/resources/user-access module.

When fixes are applied we need to regenerate the openapi schema with:
`make openapi` command.

When API spec has generated the project won't build and we need to adjust the source code.

### Expected Result
- Docs adjusted
- OpenAPI spec changed
- All source code moved from registrations to useraccess modules. Project build successfuly
- Fix tests if any

## Immediate Task Description or Request
The task is to make a `registration` module part of `useraccess` module with according adjustments to OpenAPI, documentation, architecture and source code.
