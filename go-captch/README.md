# golang生成图片验证码

## 使用的库
```shell
github.com/mojocn/base64Captcha
github.com/gin-gonic/gin
```

## 实现两个接口
### 请求验证码
#### 请求方法: get
#### 请求路径： /captcha
#### 返回内容
| 返回名称 | 返回内容          |
|------|---------------|
| data | 验证码图片的base64流 |

### 校验验证码
#### 请求方法：post
#### 请求路径: 
