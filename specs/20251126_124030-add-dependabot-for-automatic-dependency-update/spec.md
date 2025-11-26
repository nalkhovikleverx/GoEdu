# The Feature: Add GitHub dependabot for automatic dependency updates

## Task Context

Right now our project use Go modules with `go.mod` to track all dependencies. We don't want to manually update our dependencies until project is buildable.
To save some time we decided to integrate GitHub's Dependabot which will do it for us.

## Detailed Task Description and Rules

Dependabot will update our dependencies in asynchronous way so we need a mechanics to control if the change could be applied or not. For that reason
dependabot could try to build our project with "make build" command to make sure the project still buildable.

For dependabot almost all configuration need to be performed from GitHub UI. So our current change is to generate only dependabot.yml file

### Expected Result

- `dependabot.yml` file is generated
- Clear guide provided for manual steps

## Examples

As a reference, navigate to the [dependabot quick start page](https://docs.github.com/en/code-security/getting-started/dependabot-quickstart-guide)

## Immediate Task Description or Request
The task is to integrate GitHub's Dependabot for automatic dependency update with only `dependabot.yml` generation. After that provide a guidance how to proceed
for manual steps.
