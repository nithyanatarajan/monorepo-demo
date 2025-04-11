# Monorepo Demo

## 1. Repo Structure

> Walk through folders: `services/`, `pkg/`, `Taskfile.yml`

- `services/`: Microservices with CLI/API/core/mocks
- `pkg/`: Shared modules like `module-a`, `module-b`
- `Taskfile.yml`: One task runner to rule them all

## 2. Unified Dev Workflow with Taskfile

> Single Taskfile, Shared Dev Tools

```shell
# Show all tasks
task
```

2a. Targeted Builds Per Service

```shell
task build-service SERVICE=service-a
```

2b. Unified Testing and Coverage

```shell
task test-coverage
open coverage.html
```

2c. Dependency graph

```shell
task deps-graph
```

## 3. Dependency-Aware CI in Action

> Make a change in module-a
> Commit + push → only service-a is rebuilt + tested
> Highlight affected logic from logs

```shell
# a change in module-b
open https://github.com/nithyanatarajan/monorepo-demo/actions/runs/14411010863/job/40418687058
```

## 4. What if we add a new service?

```shell
mkdir -p services/service-c/{cmd,api,core}
touch services/service-c/core/core.go
```
