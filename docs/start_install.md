> **提示**: 目前为止本项目仍然是一个标准的GF项目，因此在使用本项目前，你必须熟悉使用GF，特别是其重点组件、Cli工具的使用。

## 环境要求
- Go 版本>= 1.18
- Mysql版本 >=8.0
- GoFrame版本 >=v2.5.2

## 开始使用
#### 克隆项目
```Shell
git clone https://github.com/lgcgo/letga-server
```
#### 数据导入
- 创建Mysql数据库

- 修改数据库配置 `/manifest/config/config.yaml`

```Yaml
database:
  default:
    link:  "mysql:letga:××××××@tcp(127.0.0.1:3306)/letga?loc=Local&parseTime=true"
    debug: true
```

- 导入数据 `/manifest/data/letga×××.sql`

#### 启动服务
- 进入项目根目录执行
```Shell
go run .
```
或者使用GoFrame启动命令
```Shell
gf run .
```

## 相关文档
- [GoFrame 快速开始](https://goframe.org/pages/viewpage.action?pageId=1114399)
- [GoFrame 开发工具](https://goframe.org/pages/viewpage.action?pageId=1114260)
- [FoFrame 核心组件](https://goframe.org/pages/viewpage.action?pageId=1114409)