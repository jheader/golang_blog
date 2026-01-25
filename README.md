# golang_blog
一个使用 Go 语言、Gin 框架和 GORM 库开发的个人博客系统后端 API，轻量易用、结构清晰，支持用户认证、文章管理、评论交互等核心博客功能。

## 功能特性
- ✅ 用户注册与登录（账号密码校验）
- ✅ JWT 认证与授权（接口权限控制）
- ✅ 文章全量 CRUD 操作（支持创建/查询/更新/删除）
- ✅ 评论功能（关联文章与用户）
- ✅ 全局统一错误处理（标准化异常响应）
- ✅ 结构化日志记录（基于 Logrus）
- ✅ 数据库自动迁移（GORM 自动建表）
- ✅ 分页查询（文章列表支持分页）

## 技术栈
| 技术/工具 | 版本/用途 |
|----------|-----------|
| Go       | 1.24.2（核心开发语言） |
| Gin      | 高性能 Web 框架 |
| GORM     | 数据库 ORM 框架 |
| MySQL    | 关系型数据库 |
| JWT      | 无状态身份认证 |
| Logrus   | 结构化日志库 |
| bcrypt   | 密码加密算法 |

## 项目结构
golang_blog/
├── cmd/
│   └── main.go          # 程序入口
├── config/
│   └── database.go      # 数据库配置（连接、迁移）
├── controllers/
│   ├── auth.go          # 认证控制器（注册/登录/个人信息）
│   ├── post.go          # 文章控制器（CRUD）
│   └── comment.go       # 评论控制器（查询/创建）
├── middleware/
│   ├── auth.go          # JWT认证中间件（接口鉴权）
│   ├── ErrorHandle.go   # 全局异常处理中间件
│   └── logger.go        # 日志中间件（请求/响应记录）
├── models/
│   ├── user.go          # 用户模型（Users表映射）
│   ├── post.go          # 文章模型（Posts表映射）
│   └── comment.go       # 评论模型（Comments表映射）
├── routes/
│   └── routes.go        # 路由配置（接口路由注册）
├── utils/
│   ├── jwt.go           # JWT工具（生成/解析Token）
│   ├── pageresponse.go  # 分页响应格式化
│   ├── paginate.go      # 分页查询逻辑
│   └── response.go      # 统一API响应格式
├── .env                 # 环境配置（数据库、JWT密钥等）
├── go.mod               # 依赖管理
├── go.sum               # 依赖校验
└── README.md            # 项目说明文档

## 数据库设计
### Users 表（用户表）
| 字段名        | 类型         | 说明                  |
|---------------|--------------|-----------------------|
| id            | BIGINT       | 主键（自增）|
| username      | VARCHAR(50)  | 用户名（唯一）|
| email         | VARCHAR(100) | 邮箱（唯一）|
| password      | VARCHAR(100) | bcrypt加密后的密码    |
| created_at    | DATETIME     | 创建时间              |
| updated_at    | DATETIME     | 更新时间              |
| deleted_at    | DATETIME     | 软删除标记            |

### Posts 表（文章表）
| 字段名        | 类型         | 说明                  |
|---------------|--------------|-----------------------|
| id            | BIGINT       | 主键（自增）|
| title         | VARCHAR(100) | 文章标题              |
| content       | TEXT         | 文章内容              |
| user_id       | BIGINT       | 外键（关联Users.id）|
| created_at    | DATETIME     | 创建时间              |
| updated_at    | DATETIME     | 更新时间              |
| deleted_at    | DATETIME     | 软删除标记            |

### Comments 表（评论表）
| 字段名        | 类型         | 说明                  |
|---------------|--------------|-----------------------|
| id            | BIGINT       | 主键（自增）|
| content       | TEXT         | 评论内容              |
| user_id       | BIGINT       | 外键（关联Users.id）|
| post_id       | BIGINT       | 外键（关联Posts.id）|
| created_at    | DATETIME     | 创建时间              |
| updated_at    | DATETIME     | 更新时间              |
| deleted_at    | DATETIME     | 软删除标记            |

## API 接口文档
### 认证接口
| 方法 | 路径                          | 描述                 | 权限       |
|------|-------------------------------|----------------------|------------|
| POST | /api/v1/auth/register         | 用户注册             | 公开       |
| POST | /api/v1/auth/login            | 用户登录（返回JWT）| 公开       |
| GET  | /api/v1/profile?userID={id}   | 获取用户信息         | 需要认证   |

### 文章接口
| 方法   | 路径                              | 描述                          | 权限                     |
|--------|-----------------------------------|-------------------------------|--------------------------|
| GET    | /api/v1/posts?page={page}&size={size} | 获取文章列表（支持分页）| 公开                     |
| GET    | /api/v1/posts/{id}                | 获取单篇文章详情              | 公开                     |
| POST   | /api/v1/posts/saveOrUpdate        | 创建/更新文章（有id则更新）| 需要认证                 |
| DELETE | /api/v1/posts/{id}                | 删除文章                      | 需要认证（仅文章作者）|

### 评论接口
| 方法 | 路径                              | 描述               | 权限       |
|------|-----------------------------------|--------------------|------------|
| GET  | /api/v1/comments/post/{post_id}   | 获取文章评论列表   | 公开       |
| POST | /api/v1/posts/{post_id}/comment   | 创建文章评论       | 需要认证   |

## 快速运行
### 前置条件
- 已安装 Go 1.21+ 版本
- 已安装 MySQL 并创建空数据库
- 配置 `.env` 文件（数据库连接、JWT 密钥等）

### 运行步骤
1. 克隆项目到本地
   ```bash
   git clone https://github.com/你的用户名/golang_blog.git
   cd golang_blog

2. 安装项目依赖
   go mod tidy

3. 启动项目
   go run cmd/main.go

4. 验证运行：服务器默认启动在 http://localhost:8080，可访问 http://localhost:8080/ping 测试连通性

# 使用示例（curl 命令）
## 1. 用户注册
curl --location 'http://172.30.185.167:8080/api/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "testuser11331",
    "email": "test@example1331.com",
    "password": "password123"
}'

## 2. 用户登录（获取 JWT Token）

curl --location 'http://172.30.185.210:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "testuser",
    "password": "password123"
}'

## 3. 新增文章（需携带 JWT Token）

curl --location 'http://172.30.185.210:8080/api/v1/posts/saveOrUpdate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzY5MDk0MDI1LCJpYXQiOjE3NjkwODY4MjV9.nkuWb_tYETjKqYey5XtdqEFFBoirh_sbphvlbuRuOsY' \
--data '{
    "title": "第6品111",
    "content": "这是个6测试111"
}'

## 4. 更新文章（需携带 JWT Token）

curl --location 'http://172.30.185.210:8080/api/v1/posts/saveOrUpdate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzY5MDk0MDI1LCJpYXQiOjE3NjkwODY4MjV9.nkuWb_tYETjKqYey5XtdqEFFBoirh_sbphvlbuRuOsY' \
--data '{
    "postID":2,
    "title": "第一品111",
    "content": "这是个更新测试112"
}'

## 总结
1. 内容完整覆盖项目核心信息：功能、技术栈、结构、数据库、接口、运行方式、使用示例；
2. 格式严格遵循 Markdown 标准，兼容所有主流平台的渲染规则；
3. 复制后直接保存为 `README.md` 即可使用，无需额外调整格式。
