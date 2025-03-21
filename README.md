# FastFiber 🚀 

A high-performance RESTful API scaffold based on Go Fiber framework  
基于 Go Fiber 框架的高性能 RESTful API 脚手架

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue)](https://golang.org/)
[![Fiber Version](https://img.shields.io/badge/Fiber-v2.52-brightgreen)](https://gofiber.io/)

---

## Features ✨ / 功能特性

- **Blazing Fast** 🔥 - Built with Go Fiber framework  
  极致性能 - 基于 Go Fiber 框架构建
- **Full-Stack Ready** 🛠️ - Integrated with GORM, Redis, JWT, Zap logging  
  全栈就绪 - 集成 GORM, Redis, JWT, Zap 日志
- **Efficient CLI** ⚡ - Built-in database migration & user management  
  高效命令行 - 内置数据库迁移和用户管理
- **Production Ready** 🚢 - Pre-configured for release/debug modes  
  生产就绪 - 预置发布/调试模式配置
- **Secure by Default** 🔒 - Password hashing & JWT authentication  
  安全默认 - 密码哈希与 JWT 认证

---

## Tech Stack 💻 / 技术栈

**Core**
- [Go 1.23+](https://golang.org/) - Programming Language
- [Fiber](https://gofiber.io/) - Web Framework
- [GORM](https://gorm.io/) - ORM
- [Redis](https://redis.io/) - Caching

**Utilities**
- [Zap](https://github.com/uber-go/zap) - Logging
- [JWT](https://jwt.io/) - Authentication
- [Viper](https://github.com/spf13/viper) - Config Management
- [Air](https://github.com/cosmtrek/air) - Live Reload (Dev)

---

## Quick Start 🚀 / 快速开始

### Prerequisites / 环境要求
- Go 1.23+
- MySQL 8.0+ / PostgreSQL 14+
- Redis 7.0+

### Installation / 安装

```bash
Clone repository

git clone https://github.com/unclesam-LY/FastFiber.git cd FastFiber
Install dependencies
go mod tidy
```

### Configuration / 配置
1. Copy sample config: 
```bash 
cp settings.dev.yaml settings.yaml
```

2. Edit `settings.yaml`:
- yaml 
    - system: mode: "debug" # debug/release port: 8000
    - db: mode: "mysql" 
    - host: "127.0.0.1"
    - other settings.....

---
## Usage 📖 / 使用指南

### Commands / 命令

#### 数据库迁移 / Database migration
go run main.go -db

#### 查看最近用户 / Check the latest users （limit 10）
go run main.go -m user -t list

#### 启动服务（开发模式）/ start service (development mode)
air

#### 生产环境构建 / production build
go build -ldflags "-s -w" -o app

### Directory Structure / 目录结构

``` 
FastFiber/ 
    ├── api/ # 接口控制器/API Controllers 
    ├── config/ # 配置结构体/Config Structs 
    ├── core/ # 核心初始化逻辑/Core Initialization 
    ├── global/ # 全局变量/Global Variables 
    ├── models/ # 数据模型/Data Models 
    ├── routers/ # 路由配置/Router Configuration 
    ├── utils/ # 工具包/Utilities 
    └── settings.yaml # 主配置文件/Main Config File 
```

