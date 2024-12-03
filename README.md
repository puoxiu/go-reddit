# 项目简介
该项目是一个基于 Go 语言的 Web 应用，使用了 Gin 框架，并通过 MySQL 和 Redis 来处理数据存储和缓存功能。项目包含用户认证、社区管理、帖子发布与投票等核心功能。


# 目录结构
```bash
.
├── config.yaml              # 项目配置文件
├── controllers/             # 控制器层，处理各个路由的请求逻辑
├── dao/                     # 数据访问层，包含 MySQL 和 Redis 相关操作
├── docker-compose.yaml      # Docker 配置文件，用于启动 MySQL 和 Redis 容器
├── logger/                  # 日志功能
├── logic/                   # 业务逻辑层
├── middlewares/             # 中间件，包含认证等功能
├── models/                  # 数据模型定义
├── pkg/                     # 工具包，例如 JWT 和分布式 ID 生成器
├── routes/                  # 路由定义
├── script/                  # 初始化脚本(包含数据库sql文件)
├── settings/                # 配置信息读取
├── main.go                  # 项目入口
└── README.md                # 项目文档
```


# 主要功能
* 用户注册与登录 (支持 JWT 和短信验证码)
* 社区管理
* 帖子发布与投票
* 分布式ID生成 (基于 Snowflake 算法)
* 日志记录
* 令牌桶算法限流

# 环境依赖
* Go 1.18+
* MySQL 8.0+
* Redis 6.0+
* Docker



# 快速开始
1. 克隆项目

2. 启动 MySQL 和 Redis 容器

```bash
    docker-compose up -d
```
3. 修改 config.yaml 配置文件，根据你的环境配置数据库连接和其他参数。

4. 初始化数据库


5. 启动应用

``` bash
go run main.go
```