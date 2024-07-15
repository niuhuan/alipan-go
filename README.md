alipan-go
=========

阿里云盘 SDK for Golang

## 🚀 实现功能

- [x] OAUTH
    - [x] 登录
    - [x] AccessToken自动管理
- [x] 用户
    - [x] 获取用户信息
    - [x] 获取用户云盘信息
    - [x] 获取用户空间信息
- [x] 文件
    - [x] 获取文件列表
    - [x] 获取文件信息 (单独、批量)
    - [x] 创建文件夹
    - [x] 上传文件
    - [x] 文件更名、收藏、取消收藏
    - [x] 文件移动、复制
    - [x] 文件下载（获取链接）
    - [x] 文件删除、移动到回收站
    - [x] 异步任务状态查询

## 📖 使用方法

### 📦 安装

```go
import "github.com/niuhuan/alipan-go/oauth_client"
import "github.com/niuhuan/alipan-go/adrive_client"
```

### 📃 调用

参照
- [oauth_client_test.go](oauth_client/oauth_client_test.go)
- [adrive_client_test.go](adrive_client/adrive_client_test.go)

## 📕 协议

参考 `LICENSE` 文件
