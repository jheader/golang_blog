# golang_blog
一个使用Go语言、Gin框架和GORM库开发的个人博客系统后端API。一个基于 Gin 框架构建的轻量级 Web 项目模板，包含用户认证、文章管理等核心功能，开箱即用。

# 功能特性
1. 用户注册和登录
2. JWT认证和授权
3. 文章的CRUD操作
4. 评论功能
5. 统一的错误处理
6. 日志记录
7. 数据库自动迁移

# 技术栈

1. Go 1.24.2
2. Gin - Web框架
3. GORM - ORM库
4. Mysql - 数据库
5. JWT - 身份认证
6. Logrus - 日志库
7. bcrypt - 密码加密

# 项目结构

golang_blog/
├── cmd/
│   └── main.go          # 程序入口
├── config/
│   └── database.go      # 数据库配置
├── controllers/
│   ├── auth.go          # 认证控制器
│   ├── post.go          # 文章控制器
│   └── comment.go       # 评论控制器
├── middleware/
│   ├── auth.go          # JWT认证中间件
│   ├── ErrorHandle.go   # 异常中间件
│   └── logger.go        # 日志中间件
├── models/
│   ├── user.go          # 用户模型
│   ├── post.go          # 文章模型
│   └── comment.go       # 评论模型
├── routes/
│   └── routes.go        # 路由配置
├── utils/
│   ├── jwt.go           # JWT工具
|   |—— pageresponse.go  # 分页插件
|   |—— paginate.go      # 分页插件
│   └── response.go      # 响应工具
|—— .env                 # 环境配置   
├── go.mod
├── go.sum
└── README.md

# 数据库设计

# Users表
id (主键)
username (用户名，唯一)
email (邮箱，唯一)
password (加密密码)
created_at, updated_at, deleted_at

# Posts表

id (主键)
title (标题)
content (内容)
user_id (外键，关联users表)
created_at, updated_at, deleted_at

# Comments表

id (主键)
content (内容)
user_id (外键，关联users表)
post_id (外键，关联posts表)
created_at, updated_at, deleted_at

# API接口

# 认证接口

POST /api/v1/auth/register - 用户注册
POST /api/v1/auth/login - 用户登录
GET /api/v1/api/v1/profile?userID=.. - 获取用户信息 (需要认证)

# 文章接口
GET /api/v1/posts?page=1&size=2 - 获取文章列表 (公开)
GET /api/v1/posts/:id - 获取文章详情 (公开)
POST /api/v1/posts/saveOrUpdate - 创建or更新文章 (需要认证,有id是更新)
DELETE /api/v1/posts/:id - 删除文章 (需要认证，仅作者)

# 评论接口
GET /api/v1/comments/post/:post_id - 获取文章评论 (公开)
POST /api/v1/posts/:post_id/comment - 创建评论 (需要认证)

# 运行项目

1. 确保已安装Go 1.21+
2. 克隆项目到本地
3. 安装依赖：
    go mod tidy
4. 运行项目：
    go run cmd/main.go
5. 服务器将在 http://localhost:8080 启动

# 使用示例

# 用户注册
curl --location 'http://172.30.185.167:8080/api/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "testuser11331",
    "email": "test@example1331.com",
    "password": "password123"
  }'

 # 用户登录

 curl --location 'http://172.30.185.210:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "testuser",
    "password": "password123"
  }'

  # 新增文章
  curl --location 'http://172.30.185.210:8080/api/v1/posts/saveOrUpdate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzY5MDk0MDI1LCJpYXQiOjE3NjkwODY4MjV9.nkuWb_tYETjKqYey5XtdqEFFBoirh_sbphvlbuRuOsY' \
--data '{
    "title": "第6品111",
    "content": "这是个6测试111"
  }'

  # 更新文章

  curl --location 'http://172.30.185.210:8080/api/v1/posts/saveOrUpdate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzY5MDk0MDI1LCJpYXQiOjE3NjkwODY4MjV9.nkuWb_tYETjKqYey5XtdqEFFBoirh_sbphvlbuRuOsY' \
--data '{
    "postID":2,
    "title": "第一品111",
    "content": "这是个更新测试112"
  }'


# 注意事项

1. JWT密钥在生产环境中应该从环境变量读取
2. 所有密码都会使用bcrypt进行加密存储
3. API返回统一的JSON格式响应
