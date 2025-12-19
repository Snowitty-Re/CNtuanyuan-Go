# 志愿者OA系统 - 走失人员寻亲

一个完整的志愿者OA系统，用于社会走失人员寻亲管理。

## 技术栈

### 后端
- Go 1.25
- Gin Web框架
- GORM ORM
- PostgreSQL 数据库
- JWT 认证

### 前端
- Vue 3 (Composition API)
- Vite
- TailwindCSS
- Vue Router
- Pinia

## 项目结构

```
CNtuanyuan-Go/
├── backend/          # Go后端
├── frontend/         # Vue前端
└── docs/            # 文档
```

## 快速开始

### 后端

1. 安装依赖：
```bash
cd backend
go mod tidy
```

2. 配置数据库：
编辑 `backend/config.yaml` 或设置环境变量

3. 运行服务器：
```bash
go run cmd/server/main.go
```

### 前端

1. 安装依赖：
```bash
cd frontend
npm install
```

2. 运行开发服务器：
```bash
npm run dev
```

## 环境变量

- `PORT`: 服务器端口（默认: 8080）
- `DB_HOST`: 数据库主机
- `DB_USER`: 数据库用户
- `DB_PASSWORD`: 数据库密码
- `DB_NAME`: 数据库名称
- `JWT_SECRET`: JWT密钥

## 开发计划

详见项目计划文档。

