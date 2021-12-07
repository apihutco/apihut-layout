# APIHut Project Template

APIHut 项目模板

## Make

优化Makefile以适应项目需求

- 整合 `make errors` , `make validate` 为 `make api`
- 新增 `make run` 命令以编译到 ./bin 并运行
- 新增 `make crun` 命令以清理 ./bin 目录并重新编译,运行
- 新增 `make ent` 命令以生成 ent 代码
- 新增 `make wire` 命令以生成依赖注入代码

## Data

内置 Ent ORM
