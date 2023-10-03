> 业务错误码 `BizCode`，实现了底层GF的 `gcode.Code` 接口；统一使用 `gerror.NewCode(BizCode)` 抛出业务错误，最终由中间件service.Middelware().ResponseHandler转换统一的响应格式。

## 创建与使用
在internal/biz包中提供 `WithCode`、`WithCodef` 两个方法创建业务错误码

### 使用WithCode创建错误码
```
WithCode(code gcode.Code, errorCode int, message string) gcode.Code
```

#### 参数说明

 名称 | 说明
:------:|------
code  | 基础错误码，最终决定状态码 HttpCode 类型
errorCode | 业务错误代号
message | 错误描述

#### 使用示例
```Golang
var (
	UserNotExists gcode.Code
	err error
)
// 创建错误码
UserNotExists = WithCode(CodeNotFound, 21001, "User not exists.")
// 使用错误码
err = gerror.NewCode(bizcode.UserNotExists)
```

### 使用WithCodef创建错误码
```
WithCodef(code gcode.Code, formatInfo ...interface{})
```

#### 参数说明

 名称 | 说明
:------:|------
code  | 基础错误码，最终决定状态码 HttpCode 类型
formatInfo | 格式化内容


**使用示例**
```
var (
	UserAccountExists gcode.Code
	err error
)
// 创建错误码
UserAccountExists= WithCode(CodeNil, 20003, "Account %s is already exists.")
// 使用错误码
err = gerror.NewCode(bizcode.UserNotExists, 1)
```

## 设计规范

#### 错误码命名
- 结构定义：`业务模块` + `对象/字段` + `错误类型` ，大驼峰命名
- 常用后缀：
	- ××Invalid：表示对象/字段校验不通过
	- ××Exists：表示需要写入的数据已经存在
	- ××NotExists：表示需要查询的数据不存在
	- ××Fail：表示未知具体的错误

#### 错误代号 errorCode
- 结构定义：错误代号由 `模块代号` +  `业务错误代号` ，一般前两位是以一个 `Service` 模块为单位的代号，后三位数是 `业务错误码代号`

#### 错误描述 errorMessage
错误描述不做要求，可以参考示例
