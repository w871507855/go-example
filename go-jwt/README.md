# golang使用jwt

## 使用的库
```shell
github.com/golang-jwt/jwt/v5
github.com/gin-gonic/gin
```

## 实现两个接口
### 用户登录
#### 请求方法 post
#### 请求路径 /login
#### 请求内容
| 参数名      | 参数内容 |
|----------|------|
| username | 用户名  |
| password | 密码   |
#### 返回内容
| 参数名  | 参数内容    |
|------|---------|
| data | token内容 |

### 根据token获取用户详情
#### 请求方法 get
#### 请求路径 /userinfo
#### 请求内容
header中带Authorization

