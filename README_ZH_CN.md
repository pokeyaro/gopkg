# Go 工具包

[English](README.md) | 中文

一些常用的 `GO` 基础库的封装，统一汇集到 `gopkg` 仓库中。

## 安装

```bash
go get -u github.com/pokeyaro/gopkg/<子包名>
```


## 目录

| 模块      | 简介  | 文档                                         | 日期       |
|---------|-----|--------------------------------------------|----------|
| logger  | 日志包 | [logger 使用文档](./go-logger/README_ZH_CN.md) | 23/08/07 |


## 子包管理

> `gopkg` 采用 `go module` 子包模式进行依赖和版本控制，相关介绍如下

目录结构：

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

子包名（示例）：

- `package_a/go.mod` 中位于首行的 `module` 名称为：`github.com/pokeyaro/gopkg/package_a`
- `package_b/go.mod` 中位于首行的 `module` 名称为：`github.com/pokeyaro/gopkg/package_b`


## 开发规范

待完善...
