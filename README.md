# FastFiber 🚀 

A high-performance RESTful API scaffold based on Go Fiber framework  
基于 Go Fiber 框架的高性能 RESTful API 脚手架

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue)](https://golang.org/)
[![Fiber Version](https://img.shields.io/badge/Fiber-v2.52-brightgreen)](https://gofiber.io/)

---

## Features ✨ / 功能特性

### Core Features / 核心功能
- **Blazing Fast** 🔥 - Built with Go Fiber framework  
  **极致性能** 🔥 - 基于 Go Fiber 框架构建
- **ORM Integration** 📦 - Full GORM support with MySQL/PostgreSQL  
  **ORM集成** 📦 - 完整支持 MySQL/PostgreSQL 的 GORM 集成
- **Redis Cache** 🗃️ - Built-in Redis client with connection pooling  
  **Redis缓存** 🗃️ - 内置连接池的 Redis 客户端

### Security & DevOps / 安全与运维
- **JWT Auth** 🔑 - Ready-to-use authentication middleware  
  **JWT认证** 🔑 - 开箱即用的身份验证中间件
- **Zap Logging** 📝 - Structured logging with rotation support  
  **Zap日志** 📝 - 支持日志轮转的结构化日志系统
- **Hot Reload** 🔄 - Live reload support via Air (dev mode)  
  **热重载** 🔄 - 通过 Air 支持实时重载（开发模式）

### CLI Tools / 命令行工具
- **DB Migration** 🚚 - `go run main.go -db` for schema migration  
  **数据库迁移** 🚚 - 使用 `go run main.go -db` 执行架构迁移
- **User Management** 👥 - CRUD operations via CLI flags  
  **用户管理** 👥 - 通过命令行参数实现增删改查

---

## Tech Stack 💻 / 技术栈

**Core Frameworks / 核心框架**
- [Go 1.23+](https://golang.org/) - Language Runtime / 语言运行时
- [Fiber v2.52](https://gofiber.io/) - HTTP Server Framework / HTTP服务框架

**Database & Cache / 数据库与缓存**
- [GORM 1.25](https://gorm.io/) - ORM with Auto-Migration / 支持自动迁移的ORM
- [Redis 7.0](https://redis.io/) - Distributed Caching / 分布式缓存

**DevOps Tools / 运维工具**
- [Zap 1.27](https://go.uber.org/zap) - Structured Logging / 结构化日志
- [Air 1.49](https://github.com/cosmtrek/air) - Live Reload (Dev) / 开发热重载
- [Viper 1.18](https://github.com/spf13/viper) - Config Management / 配置管理

**Security / 安全**
- [JWT 5.0](https://jwt.io/) - Token Authentication / 令牌认证
- [Bcrypt 4.0](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password Hashing / 密码哈希

---

## Quick Start 🚀 / 快速开始

### Prerequisites / 环境要求
- Go 1.23+
- MySQL 8.0+ / PostgreSQL 14+
- Redis 7.0+

### Installation / 安装

```bash
# Clone repository / 克隆仓库
git clone https://github.com/unclesam-LY/FastFiber.git
cd FastFiber

# Install dependencies / 安装依赖
go mod tidy
```

### Configuration / 配置
1. Copy sample config / 复制示例配置: 
```bash 
cp settings.dev.yaml settings.yaml
```

2. Edit `settings.yaml` / 编辑配置文件:
```yaml 
system:
  mode: "debug" # debug/release
  port: 8000
db:
  mode: "mysql" 
  host: "127.0.0.1"
  # other settings... / 其他设置...
```

---
## Usage 📖 / 使用指南

### Command Examples / 命令示例

#### Database Operations / 数据库操作
```bash
# 自动迁移所有模型
# Auto migrate all models
go run main.go -db

# 回滚最近迁移
# Rollback last migration
go run main.go -db -rollback
```

#### User Management / 用户管理
```bash
# 创建新用户（需root权限）
# Create new user (requires root)
go run main.go -m user -t create -name "John Doe" -email john@example.com

# 列出最近10个用户（JSON格式）
# List recent 10 users (JSON format)
go run main.go -m user -t list -format json
```

#### Service Control / 服务控制
```bash
# 开发模式（带热重载）
# Development mode (with hot reload)
air

# 生产模式（端口8080）
# Production mode (port 8080)
MODE=release PORT=8080 go run main.go
```

### Configuration / 配置说明
```yaml
# 环境变量覆盖配置示例
# Environment variables override example
export DB_HOST=cluster.pg.aws.com
export JWT_SECRET=my_super_secret
go run main.go
```

---

### Directory Structure / 目录结构

```
FastFiber/ 
    ├── api/            # API Controllers / 接口控制器
    ├── config/         # Config Structs / 配置结构体
    ├── core/           # Core Initialization / 核心初始化逻辑
    ├── global/         # Global Variables / 全局变量
    ├── models/         # Data Models / 数据模型
    ├── routers/        # Router Configuration / 路由配置
    ├── utils/          # Utilities / 工具包
    └── settings.yaml   # Main Config File / 主配置文件
```

---

## Environment Variables 🌱 / 环境变量
| 变量名          | 示例值                  | 说明                          |
|----------------|-------------------------|-----------------------------|
| DB_HOST        | cluster.pg.aws.com      | 数据库主机地址                |
| JWT_SECRET     | my_super_secret         | JWT签名密钥                   |
| REDIS_PASSWORD | redis_prod_password     | Redis连接密码                |

## Testing 🧪 / 测试方法
```bash
# 运行所有测试用例
# Run all test cases
go test -v ./...

# 生成测试覆盖率报告
# Generate coverage report
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

## License / 许可证

MIT

---

## Contributing 🤝 / 贡献指南

### Workflow / 工作流程
1. Fork 仓库并创建特性分支  
   Fork repo & create feature branch
2. 提交清晰的 commit 信息  
   Commit with clear messages
3. 添加测试用例（至少覆盖率80%）  
   Add test cases (min 80% coverage)
4. 更新相关文档  
   Update relevant documentation
5. 发起 Pull Request 到 develop 分支  
   Open PR to develop branch

### Code Style / 代码规范
- Go 代码遵循标准格式（go fmt）  
  Go code follows standard formatting
- 接口命名使用 `I` 前缀（如 `IUserService`）  
  Interface names use `I` prefix
- 错误处理使用 global/response 包  
  Error handling uses global/response pkg

---

