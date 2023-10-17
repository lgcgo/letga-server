## 权限认证

Letga 的权限认证系统，建立在Casbin [域内基于角色的访问控制](https://casbin.org/zh/docs/rbac-with-domains) 的基础之上。在此系统中，分别有 `User` 、`Role`、`Route` 这四个数据实体 `Entity` 参与其中，这些数据实体中的某些字段，将与Casbin 模型中的令牌名称形成映射关系。

### RBAC 模型

这里假设你已经了解Casbin的工作原理，Letga使用了以下 **RBAC模型**
```
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && r.act == p.act || r.sub == "role::root"
```

我们把模型中定义的 sub（访问实体Subject）, obj（访问资源Object）, act（访问动作Action） 称之为令牌名称 `Token`，他们与系统的数据实体 `Entity` 中指定字段存在映射关系，最终实现的角色关系政策、路由授权政策。

### 角色关系政策
- 示例：
```
g, role::roleName, role::subRoleName
```
- 说明：表示访问角色 `roleName` 拥有角色 `subRoleName` 的所有权限；此时 `r.sub` 和 `p.sub` 分别对应 `role::roleName`和 `role::subRoleName`，作用于匹配规则中的 `g(r.sub, p.sub)`

### 路由授权政策
- 示例：
```
p, role::roleName, /account/info, GET
```
说明：表示角色default，拥有 /account/info 接口的 GET 请求权限；此时sub，obj，act 分别对应`role::roleName`， `/account/info` 和 `GET`

## 授权 Authorization

### 基于JWT的权限授权

我们到 [https://jwt.io](https://jwt.io)，把签发的Token字符串解密，其中PAYLOAD是长这样的
```
{
  "ist": "grant",
  "isr": "Manager",
  "dom": "manager",
  "iss": "mongoapi.com",
  "sub": "3y4eyz0e440crq24017elbk300unzx8d",
  "exp": 1681541407,
  "nbf": 1681455007,
  "iat": 1681455007
}
```

### 签发字段

字段 | 名称 | 说明 
:------:|:------:|----
ist | 签发类型| grant=签发；refresh=刷新
isr | 签发角色| 用户登录时申请的登录角色
iss| 授权方 | jwt标准字段
sub | 主题 | jwt标准字段，这里是用户的UUID 

## 中间件
待补充...

### 扩展资料
- 域内基于角色的访问控制 https://casbin.org/zh/docs/rbac-with-domains