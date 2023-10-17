## 介绍

Letga 是一个基于 GoFrame 和 AntDesign 的中后台管理系统。Letga 集成了通用的中后台基础功能组件，是一款 **规范化**、**易扩展**、**体验佳**的企业级开源系统。

### 规范化、易扩展

GoFrame 是 LetgaServer 的底座，具备基础组件丰富、文档全面、通用性强等特性。在此基础上，Letga 针对应用提出更细致的**规范与约束**方案。

### 体验佳

得益于AntDesign 的设计价值观，Letga 在整体设计上践行 **简洁**、**一致**、**易用** 的设计原则，并对通用的功能组件做了进一步封装，开发者可轻易复用或者扩展 **体验良好** 的功能组件。

Letga 是前后端分离项目，本项目是后端部分，前端请前往：[LetgaFrontend](https://github.com/lgcgo/letga-frontend/)

## 特性

- 遵循 `OpenAPIv3` 规范，自动构建`Swagger`文档
- 基于 `Casbin` + `JWT` 高效的 RBAC 权限认证设计
- 数据 ID 索引加密，接口层隐藏自增 ID

## 文档

- [安装使用](docs/start_install.md)
- [权限系统](docs/sys_authority.md)
- [错误处理](docs/sys_error.md)

## 截图

| ![](https://github.com/lgcgo/letga-server/assets/42335782/d4a310fb-59a6-41fb-971c-13a02ea35c43) | ![](https://github.com/lgcgo/letga-server/assets/42335782/7723dde4-3e37-4a40-8c96-b1c52e2a253e) |
|:-----------------------------------------------------------------------------------------------:|:----|
| ![](https://github.com/lgcgo/letga-server/assets/42335782/e4b30c98-969c-48ca-91f3-bb48eff698bc) | ![](https://github.com/lgcgo/letga-server/assets/42335782/50efa16a-a908-4cce-b517-f66ceee0f4a8) |
| ![](https://github.com/lgcgo/letga-server/assets/42335782/48b3da26-b16c-40bc-a61c-a2dcb318e748) | ![](https://github.com/lgcgo/letga-server/assets/42335782/4dfa2d8c-6951-4b59-bdfc-88c8958b7f35) |

## 致谢

- [github.com/gogf/gf](https://github.com/gogf/gf)
- [github.com/ant-design/ant-design-pro](https://github.com/ant-design/ant-design-pro)
- [github.com/casbin/casbin](https://github.com/casbin/casbin)
- [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)

## 版权

遵循[Apache-2.0 License](https://github.com/lgcgo/letga-server/blob/main/LICENSE)，保留系统版权，可免费商用。
