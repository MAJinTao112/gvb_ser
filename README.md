# GVB博客系统后端服务

基于 Go + Gin + Vue 开发的个人博客系统后端服务

## 📋 项目简介

GVB（Go-Vue-Blog）是一个功能完善的个人博客系统后端服务，提供文章管理、用户系统、评论互动、图片管理等核心功能。系统采用前后端分离架构，本项目为后端API服务。

## ✨ 主要特性

- 🔐 **用户认证**：支持邮箱登录、QQ第三方登录，JWT Token认证
- 📝 **文章系统**：文章发布、编辑、标签分类、全文检索
- 💬 **互动功能**：评论、点赞、收藏
- 🖼️ **图片管理**：支持本地存储和七牛云CDN
- 📢 **广告管理**：首页广告位管理
- 🏷️ **标签系统**：文章标签分类管理
- 💌 **消息系统**：站内消息通知
- 📊 **数据统计**：文章浏览量、点赞数统计
- 📖 **API文档**：集成Swagger在线文档

## 🛠️ 技术栈

### 核心框架
- **Go 1.25.1**：编程语言
- **Gin**：Web框架
- **GORM**：ORM框架

### 数据存储
- **MySQL**：关系型数据库
- **Redis**：缓存和会话管理
- **Elasticsearch**：全文搜索引擎

### 第三方服务
- **七牛云**：对象存储CDN
- **QQ互联**：第三方登录
- **邮件服务**：SMTP邮件发送

### 工具库
- **JWT**：用户认证
- **Swagger**：API文档生成
- **Logrus**：日志管理
- **Validator**：参数验证
- **YAML**：配置文件解析

## 📁 项目结构

```
gvb_server/
├── api/                    # API接口层
│   ├── advert_api/        # 广告管理
│   ├── article_api/       # 文章管理
│   ├── digg_api/          # 点赞功能
│   ├── image_api/         # 图片管理
│   ├── menu_api/          # 菜单管理
│   ├── message_api/       # 消息管理
│   ├── settings_api/      # 系统设置
│   ├── tag_api/           # 标签管理
│   └── user_api/          # 用户管理
├── config/                # 配置文件结构体
├── core/                  # 核心初始化（数据库、日志等）
├── docs/                  # Swagger文档
├── flag/                  # 命令行工具
├── global/                # 全局变量
├── middleware/            # 中间件（JWT认证等）
├── models/                # 数据模型
│   ├── ctype/            # 自定义类型
│   └── res/              # 响应结构
├── plugins/               # 插件
│   ├── email/            # 邮件插件
│   ├── qiniu/            # 七牛云插件
│   └── qq/               # QQ登录插件
├── routers/               # 路由配置
├── service/               # 业务逻辑层
│   ├── common/           # 通用服务
│   ├── es_ser/           # Elasticsearch服务
│   ├── image_ser/        # 图片服务
│   ├── redis_ser/        # Redis服务
│   └── user_ser/         # 用户服务
├── utils/                 # 工具函数
├── uploads/               # 上传文件存储
├── settings.yaml          # 配置文件
├── go.mod                # Go模块依赖
└── main.go               # 主入口文件
```

## 🚀 快速开始

### 环境要求

- Go 1.25+
- MySQL 5.7+
- Redis 5.0+
- Elasticsearch 7.0+

### 安装步骤

1. **克隆项目**
```bash
git clone <repository-url>
cd gvb_server
```

2. **安装依赖**
```bash
go mod download
```

3. **配置文件**

复制 `settings.yaml.example` 为 `settings.yaml`（如果不存在则创建），并修改配置：

```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  user: root
  password: your_password
  db: gvb_db

redis:
  ip: 127.0.0.1
  port: 6379
  password: 

es:
  host: "127.0.0.1"
  port: 9200
  
system:
  host: 0.0.0.0
  port: 8080
  env: dev

jwt:
  secret: your_jwt_secret
  expires: 2
  issuer: gvb_blog
```

4. **初始化数据库**

```bash
# 创建数据库
mysql -u root -p
CREATE DATABASE gvb_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 执行数据迁移
go run main.go -db
```

5. **创建管理员用户**

```bash
go run main.go -u admin -p admin123
```

6. **启动服务**

```bash
go run main.go
# 或使用 air 热加载
air
```

7. **访问服务**

- API服务：http://localhost:8080
- Swagger文档：http://localhost:8080/swagger/index.html

## 📚 API文档

### 访问方式

启动服务后访问：http://localhost:8080/swagger/index.html

### 主要API模块

#### 用户模块 `/api/users`
- `POST /api/email_login` - 邮箱登录
- `POST /api/qq_login` - QQ登录
- `GET /api/users` - 用户列表
- `POST /api/users` - 创建用户
- `DELETE /api/users` - 删除用户
- `PUT /api/user_role` - 更新用户角色
- `PUT /api/user_password` - 修改密码
- `POST /api/logout` - 退出登录

#### 文章模块 `/api/articles`
- `GET /api/articles` - 文章列表
- `POST /api/articles` - 创建文章
- `GET /api/articles/:id` - 文章详情
- `PUT /api/articles/:id` - 更新文章
- `DELETE /api/articles` - 删除文章
- `GET /api/articles/calendar` - 文章日历
- `GET /api/articles/tags` - 文章标签统计
- `POST /api/articles/collects` - 收藏文章
- `GET /api/articles/collects` - 收藏列表

#### 图片模块 `/api/images`
- `GET /api/images` - 图片列表
- `POST /api/images` - 上传图片
- `PUT /api/images` - 更新图片
- `DELETE /api/images` - 删除图片
- `GET /api/images/names` - 图片名称列表

#### 广告模块 `/api/adverts`
- `GET /api/adverts` - 广告列表
- `POST /api/adverts` - 创建广告
- `PUT /api/adverts/:id` - 更新广告
- `DELETE /api/adverts` - 删除广告

#### 标签模块 `/api/tags`
- `GET /api/tags` - 标签列表
- `POST /api/tags` - 创建标签
- `PUT /api/tags/:id` - 更新标签
- `DELETE /api/tags` - 删除标签

#### 菜单模块 `/api/menus`
- `GET /api/menus` - 菜单列表
- `POST /api/menus` - 创建菜单
- `GET /api/menus/:id` - 菜单详情
- `PUT /api/menus/:id` - 更新菜单
- `DELETE /api/menus` - 删除菜单

#### 消息模块 `/api/messages`
- `GET /api/messages` - 消息列表
- `POST /api/messages` - 发送消息
- `GET /api/messages_all` - 所有消息
- `GET /api/messages/record` - 聊天记录

#### 点赞模块 `/api/digg`
- `POST /api/digg/article` - 点赞文章

#### 系统设置 `/api/settings`
- `GET /api/settings/:name` - 获取设置
- `PUT /api/settings/:name` - 更新设置

## 🔑 认证说明

### JWT Token

大部分API需要在请求头中携带Token：

```
Authorization: Bearer <your_jwt_token>
```

### 权限等级

- **游客**：只能查看公开内容
- **普通用户**：可以发布文章、评论、点赞
- **管理员**：拥有所有权限

## 📝 命令行工具

项目提供了一些命令行工具：

```bash
# 迁移数据库表结构
go run main.go -db

# 创建用户
go run main.go -u username -p password

# 同步ES索引
go run main.go -es
```

## ⚙️ 配置说明

### MySQL配置
```yaml
mysql:
  host: 127.0.0.1      # 数据库地址
  port: 3306           # 端口
  user: root           # 用户名
  password: root       # 密码
  db: gvb_db          # 数据库名
  log_level: dev       # 日志级别
```

### Redis配置
```yaml
redis:
  ip: 127.0.0.1        # Redis地址
  port: 6379           # 端口
  password:            # 密码
  pool_size: 100       # 连接池大小
```

### 七牛云配置
```yaml
qi_niu:
  enable: true         # 是否启用
  access_key: xxx      # AccessKey
  secret_key: xxx      # SecretKey
  bucket: xxx          # 存储空间名
  cdn: xxx             # CDN域名
  zone: z2             # 存储区域
  size: 5              # 文件大小限制(MB)
```

### JWT配置
```yaml
jwt:
  secret: your_secret  # 密钥
  expires: 2           # 过期时间(小时)
  issuer: gvb_blog    # 签发者
```

### 邮件配置
```yaml
email:
  host: smtp.qq.com    # SMTP服务器
  port: 465            # 端口
  user: xxx@qq.com     # 邮箱账号
  password: xxx        # 授权码
  use_ssl: true        # 使用SSL
```

## 📊 数据库模型

### 主要数据表

- `user_models` - 用户表
- `article_models` - 文章表
- `tag_models` - 标签表
- `banner_models` - 图片表
- `advert_models` - 广告表
- `menu_models` - 菜单表
- `menu_image_models` - 菜单图片关联表
- `message_models` - 消息表
- `comment_models` - 评论表
- `user_collect_models` - 用户收藏表
- `login_data_models` - 登录日志表

## 🔧 开发说明

### 添加新的API

1. 在 `api/` 目录下创建对应模块
2. 在 `routers/` 中注册路由
3. 在 `models/` 中定义数据模型
4. 在 `service/` 中实现业务逻辑
5. 添加Swagger注释

### 日志管理

项目使用 Logrus 进行日志管理：

```go
global.Log.Info("信息日志")
global.Log.Error("错误日志")
global.Log.Warn("警告日志")
```

### 错误处理

使用统一的响应格式：

```go
res.OkWithMessage("操作成功", c)
res.FailWithMessage("操作失败", c)
res.OkWithData(data, c)
res.FailWithCode(res.ArgumentError, c)
```

## 🐛 常见问题

### 1. 无法连接数据库
- 检查MySQL是否启动
- 确认配置文件中的数据库信息是否正确
- 检查数据库是否存在

### 2. Swagger文档无法访问
- 运行 `swag init` 重新生成文档
- 确认服务正常启动

### 3. 上传图片失败
- 检查 `uploads/` 目录权限
- 如果使用七牛云，检查配置是否正确

### 4. JWT Token失效
- 检查JWT配置中的过期时间
- 确认系统时间是否正确

## 📄 开源协议

本项目采用 MIT 协议开源

## 👨‍💻 作者信息

- 作者：海涛
- 邮箱：398886755@qq.com
- 博客：https://www.fengfengzhidao.com

## 🙏 致谢

感谢所有为这个项目做出贡献的开发者！

## 📮 联系方式

如有问题或建议，欢迎通过以下方式联系：

- 提交 Issue
- 发送邮件至：398886755@qq.com

---

⭐ 如果这个项目对你有帮助，欢迎 Star！
