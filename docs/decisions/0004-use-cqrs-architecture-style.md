---
status: accepted
date: 2024-03-15
---
# Use CQRS architectural style

## Context and Problem Statement

Our application should handle 2 types of requests - reading and writing.
We want to improve the maintanability of our application by separating concreate implementation of reading and writing from each other and evolve them independently. We haven't yet decided what implementation we will use, but we want to defer this decision for now.

## Considered Options

* Implement CQRS
* Do not implement CQRS

## Decision Outcome

Chosen option: "Implement CQRS", because we want a separate model for reading and writing for all of our modules.
Also we've decided to allow return results after command processing. In some cases there are an empty objects, but we can extendend them if needed.

### Consequences

* We can optimize/evolve independently write and read side.
* We can return resulting ID immediately.
