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
│   ├── cmd/server/   # 主程序入口
│   ├── internal/     # 内部包
│   │   ├── config/  # 配置管理
│   │   ├── models/  # 数据模型
│   │   ├── handlers/# HTTP处理器
│   │   ├── middleware/# 中间件
│   │   ├── services/# 业务逻辑
│   │   ├── repositories/# 数据访问层
│   │   └── utils/   # 工具函数
│   └── migrations/  # 数据库迁移
├── frontend/         # Vue前端
│   └── src/
│       ├── api/     # API调用
│       ├── components/# 组件
│       ├── views/   # 页面视图
│       ├── stores/  # Pinia状态管理
│       └── router/  # 路由配置
└── docs/            # 文档
```

## 已实现功能

### 后端功能
- ✅ 完整的权限管理系统（RBAC）
- ✅ JWT认证和授权
- ✅ 用户注册和登录
- ✅ 角色和权限管理
- ✅ 工作流系统（可自定义字段）
- ✅ 走失人员数据库
- ✅ 自定义表单系统
- ✅ 文件上传功能

### 前端功能
- ✅ 用户登录和注册界面
- ✅ 路由守卫
- ✅ 走失人员列表和搜索
- ✅ 响应式设计（TailwindCSS）

## 快速开始

### 后端

1. 安装依赖：
```bash
cd backend
go mod tidy
```

2. 配置数据库：
编辑 `backend/config.yaml` 或设置环境变量：
- `DB_HOST`: 数据库主机（默认: localhost）
- `DB_PORT`: 数据库端口（默认: 5432）
- `DB_USER`: 数据库用户（默认: postgres）
- `DB_PASSWORD`: 数据库密码
- `DB_NAME`: 数据库名称（默认: cntuanyuan）
- `JWT_SECRET`: JWT密钥

3. 运行服务器：
```bash
go run cmd/server/main.go
```

服务器将在 `http://localhost:8080` 启动

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

前端将在 `http://localhost:3000` 启动

## API端点

### 认证
- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录
- `GET /api/v1/auth/profile` - 获取当前用户信息

### 权限管理
- `POST /api/v1/permissions/roles` - 创建角色
- `GET /api/v1/permissions/roles` - 获取角色列表
- `POST /api/v1/permissions/permissions` - 创建权限
- `GET /api/v1/permissions/permissions` - 获取权限列表

### 工作流
- `POST /api/v1/workflows` - 创建工作流
- `GET /api/v1/workflows` - 获取工作流列表
- `POST /api/v1/workflows/fields` - 添加字段
- `POST /api/v1/workflows/states` - 添加状态
- `POST /api/v1/workflows/instances` - 创建实例

### 走失人员
- `POST /api/v1/missing-persons` - 创建走失人员记录
- `GET /api/v1/missing-persons` - 获取走失人员列表
- `GET /api/v1/missing-persons/search` - 搜索走失人员
- `GET /api/v1/missing-persons/:id` - 获取详情
- `PUT /api/v1/missing-persons/:id` - 更新记录
- `DELETE /api/v1/missing-persons/:id` - 删除记录
- `POST /api/v1/missing-persons/upload-photo` - 上传照片

### 表单
- `POST /api/v1/forms` - 创建表单
- `GET /api/v1/forms` - 获取表单列表
- `POST /api/v1/forms/submit` - 提交表单

## 数据库设计

系统包含以下核心表：
- `users` - 用户表
- `roles` - 角色表
- `permissions` - 权限表
- `user_roles` - 用户角色关联表
- `role_permissions` - 角色权限关联表
- `workflows` - 工作流表
- `workflow_fields` - 工作流字段表
- `workflow_states` - 工作流状态表
- `workflow_instances` - 工作流实例表
- `missing_persons` - 走失人员表
- `forms` - 表单表
- `form_submissions` - 表单提交表

## 开发状态

项目核心功能已完成，包括：
- 完整的后端API
- 基础的前端界面
- 权限管理系统
- 工作流系统
- 走失人员管理
- 自定义表单系统

后续可以继续完善：
- 前端工作流配置界面
- 前端表单构建器
- 后台管理仪表盘
- UI/UX优化
- 单元测试和集成测试

## 许可证

MIT License

