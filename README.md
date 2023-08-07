# gopkg

English | [中文](README_ZH_CN.md)

This repository collects some commonly used Go basic libraries and wraps them into packages, which are centralized in the `gopkg` repository.

## Installation

```bash
go get -u github.com/pokeyaro/gopkg/<sub-package>
```


## Catalogs

| Module   | Description | Documentation                               | Date |
|---------|--------------|---------------------------------------------|----|
| logger  | Logging package | [logger Usage Guide](./go-logger/README.md) | Date |


## Subpackage Management

> `gopkg` uses `go module` subpackage mode for dependency and version control. Here's an overview

Directory structure:

```textmate
├── .gitignore.go
├── LICENSE
├── README.md
├── package_a
│   ├── go.mod
│   ├── go.sum
│   ├── README.md
│   └── entry_a.go
├── package_b
│   ├── go.mod
│   ├── go.sum
│   ├── README.md
│   └── entry_b.go
...
```

Subpackage names (example):

- The `module` name in the first line of `package_a/go.mod` is: `github.com/xxx/gopkg/package_a`
- The `module` name in the first line of `package_b/go.mod` is: `github.com/xxx/gopkg/package_b`


## Development Guidelines

TBD...
