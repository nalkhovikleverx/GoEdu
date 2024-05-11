---
status: accepted
date: 2024-05-11
---
# Use domain API modules

[Discussion](https://github.com/nalkhovikleverx/GoEdu/discussions/10)

## Context and Problem Statement

When designing modular monolith there are several problems we may face:

1. Circular dependencies
2. Build time
3. Slow test runs within a module

Typical application may contain circular dependencies between modules.
For example, we may have a module `order` that depends on `customer` module and `customer` module that depends on `order` module - case when
the order need to have information about a customer and customer need to have information about his orders.

`Go modules` prohibit circular dependencies. Another downside is build time. Modifications in one module might necessitate rebuilding all modules. Additionally, running tests for a single module might be slower because all dependent modules could need recompilation.

## Considered Options

* Split modules into two parts: domain API and implementation
* Do not split modules

## Decision Outcome

Chosen option: "Split modules to domain API and implementation", because we cannot avoid circular dependencies in our application. Also we want to improve build time and test runs.

## Pros and Cons of the Options

### Splitting modules into domain API and implementation

* Good, because we avoid circular dependencies
* Good, because we can improve build time and test runs time
* Neutral, because we need to maintain two parts of the module which leads to extra code and code duplication
* Bad, because it may be hard to understand the codebase

### Do not split modules

* Bad, because we cannot avoid circular dependencies

## More Information

Proof of concept is described in the following article:
[Physical design principles for faster builds](https://microservices.io/post/architecture/2023/09/12/how-modular-can-your-monolith-go-part-4-physical-design.html)
