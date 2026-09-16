# 豆包空间

轻量级、单用户的个人工作台，目标运行在 macOS，也支持在 Ubuntu 上开发和交叉编译。

## 技术栈

- 前端：Vue 3、Vite、Element Plus、Vue Router
- 后端：Go、Gin
- 数据库：SQLite（纯 Go 驱动，避免 CGO 影响跨平台构建）
- 发布：前端静态资源通过 `go:embed` 嵌入 Go 二进制

## 目录结构

```text
cmd/doubao-space/       Go 程序入口
internal/config/        配置和数据目录解析
internal/database/      SQLite 初始化和迁移
internal/httpapi/       HTTP 路由和 API
web/frontend/           Vue 前端源码
web/dist/               前端构建产物（发布时嵌入二进制）
```

## 开发

需要 Go、Node.js 和 npm。启动后端：

```bash
go run ./cmd/doubao-space --dev
```

另开一个终端启动前端：

```bash
cd web/frontend
npm ci
npm run dev
```

前端开发服务器会把 `/api` 请求代理到 `http://127.0.0.1:8080`。

也可以使用 Make：

```bash
make dev-backend
make dev-frontend
```

## 构建

构建前端并生成当前平台二进制：

```bash
make build
```

从 Ubuntu 构建 Apple Silicon 版本：

```bash
make build-darwin-arm64
```

构建产物在 `build/`。正式分发时还需要在 macOS 环境完成 `.app` 打包、签名和公证。

## 数据目录

运行时数据不放在源码目录或 `.app` 内部：

- 开发模式 `--dev`：默认使用项目根目录下的 `data/`
- 生产模式 macOS：默认使用系统用户配置目录下的 `DoubaoSpace/`
- 可通过 `--data-dir` 或 `DOUBAO_DATA_DIR` 覆盖

当前数据目录包含：

```text
data/
├── doubao.db       SQLite 数据库
├── backups/        数据库备份
└── uploads/        用户上传文件
```
