# APIHut Project Template
[![Go](https://github.com/apihutco/apihut-layout/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/apihutco/apihut-layout/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/apihutco/apihut-layout.svg)](https://pkg.go.dev/github.com/apihutco/apihut-layout)
[![codecov](https://codecov.io/gh/apihutco/apihut-layout/branch/main/graph/badge.svg?token=MX523BC5CR)](https://codecov.io/gh/apihutco/apihut-layout)
[![Go Report Card](https://goreportcard.com/badge/github.com/apihutco/apihut-layout)](https://goreportcard.com/report/github.com/apihutco/apihut-layout)
![GitHub](https://img.shields.io/github/license/apihutco/apihut-layout)

APIHut 项目模板,使用 [Kratos](https://github.com/go-kratos/kratos) 框架

## Make

优化Makefile以适应项目需求

- 整合 `make errors` , `make validate` 为 `make api`
- 新增 `make run` 命令以编译到 ./bin 并运行
- 新增 `make crun` 命令以清理 ./bin 目录并重新编译,运行
- 新增 `make ent` 命令以生成 ent 代码
- 新增 `make wire` 命令以生成依赖注入代码

## 组件
|组件|名称|
|---|---|
|Data|Ent|
|Log|Zap|
|服务发现|Nacos|


## Use
```shell
kratos new -r https://github.com/apihutco/apihut-layout.git
```