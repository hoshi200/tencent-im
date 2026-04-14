# 项目知识库

**生成时间:** 2026-04-13
**提交:** 5e06582
**分支:** master

## 概述

腾讯云 IM REST API 的 Go SDK 封装库。模块 `github.com/hoshi200/tencent-im`，Go 1.22.3，单一外部依赖 `github.com/dobyte/http`。

## 结构

```
tencent-im/
├── im.go              # 门面入口：IM 接口 + NewIM() 构造器 + sync.Once 懒加载模块
├── im_test.go         # 集成测试（1511行，硬编码凭证，调用真实API）
├── account/           # 账号管理：导入/删除/查询/踢人/在线状态
├── callback/          # 回调事件：HTTP 监听器，22种腾讯IM回调事件
├── group/             # 群组管理：最大模块（30+方法），6个文件
├── sns/               # 关系链：好友/黑名单/好友分组
├── private/           # 单聊消息：发送/导入/查询/撤回/已读
├── push/              # 全员推送：广播/用户属性/标签
├── profile/           # 用户资料：设置/拉取
├── mute/              # 全局禁言：设置/查询
├── operation/         # 运营：数据/历史记录/IP列表
├── recentcontact/     # 最近联系人：会话列表/删除
├── example/           # 使用示例（注意：存在编译bug，缺少 im 包导入）
└── internal/          # 内部包（不可外部导入）
    ├── core/          #   HTTP 客户端 + UserSig 鉴权 + URL 构建
    ├── sign/          #   HMAC-SHA256 UserSig 签名生成
    ├── entity/        #   公共实体结构体（消息/离线推送/用户）
    ├── enum/          #   公共枚举常量
    ├── types/         #   响应接口契约
    ├── conv/          #   类型转换工具
    └── random/        #   随机数工具
```

## 查找指南

| 需求 | 位置 | 备注 |
|------|------|------|
| 新增 API 模块 | 顶层目录（如 `push/`） | 参照任意现有模块的 `api.go` + `types.go` + `enum.go` 三文件结构 |
| 修改请求/响应结构 | 各模块的 `types.go` | 结构体字段与腾讯云 API 文档对应 |
| 修改鉴权逻辑 | `internal/core/client.go` | UserSig 生成、URL 拼装、请求分发 |
| 添加回调事件 | `callback/types.go` + `callback/callback.go` | 定义事件结构体 → 注册事件常量 → 实现解析 |
| 修改签名算法 | `internal/sign/sign.go` | HMAC-SHA256 |
| 公共实体/枚举 | `internal/entity/`、`internal/enum/` | 跨模块复用的结构体和常量 |
| API 使用示例 | `example/main.go`、`README.md` | README 含完整 SDK 列表和方法对照表 |

## 代码地图

| 符号 | 类型 | 位置 | 作用 |
|------|------|------|------|
| `IM` | interface | `im.go:31` | SDK 主入口，所有模块的访问门面 |
| `NewIM` | func | `im.go:114` | 构造器，接收 Options 返回 IM |
| `Options` | struct | `im.go:56` | 配置：AppId/AppSecret/UserId/Expiration |
| `core.Client` | interface | `internal/core/client.go` | HTTP 客户端，负责鉴权和请求 |
| `core.NewClient` | func | `internal/core/client.go` | 客户端构造 |
| `sign.GenUserSig` | func | `internal/sign/sign.go` | UserSig 签名生成 |
| `callback.Callback` | interface | `callback/callback.go` | 回调事件管理器 |
| `sns.API` | interface | `sns/api.go` | 关系链管理（好友/黑名单/分组） |
| `group.API` | interface | `group/api.go` | 群组管理（30+方法） |
| `account.API` | interface | `account/api.go` | 账号管理 |
| `private.API` | interface | `private/api.go` | 单聊消息 |
| `push.API` | interface | `push/api.go` | 全员推送 |
| `profile.API` | interface | `profile/api.go` | 用户资料 |
| `mute.API` | interface | `mute/api.go` | 全局禁言 |
| `operation.API` | interface | `operation/api.go` | 运营管理 |
| `recentcontact.API` | interface | `recentcontact/api.go` | 最近联系人 |

## 约定

- **模块文件结构**：每个公开模块固定包含 `api.go`（接口+实现）、`types.go`（请求/响应结构体）、`enum.go`（常量）。复杂模块可额外拆分（如 `group/` 的 `group.go`、`member.go`、`message.go`、`filter.go`）
- **构造函数模式**：`NewAPI(core.Client) API` — 每个模块的构造函数签名统一
- **懒加载**：`im.go` 中所有模块通过 `sync.Once` 延迟初始化
- **错误处理**：`core.Error` 类型（含 `Code()` + `Message()`），通过 `err.(im.Error)` 类型断言区分 API 错误和普通错误
- **方法命名**：单数方法调用复数方法（如 `AddFriend` 内部调用 `AddFriends`），避免代码重复
- **续拉模式**：`Pull*` 方法基于 `Fetch*` 封装，自动处理分页续拉逻辑
- **注释语言**：中文

## 反模式（本项目）

- **不要** 将 `internal/` 包的内容暴露给外部使用者
- **不要** 修改 `go.sum` — 它被 `.gitignore` 忽略（虽然这不符合标准实践）
- **不要** 直接运行 `im_test.go` — 它包含硬编码凭证并调用真实 API

## 已知问题

- `example/main.go` 缺少 `im` 包导入，无法编译
- `internal/entity/mesage.go` 文件名拼写错误（应为 `message.go`）
- `internal/core/client.go` 使用了 Go 1.20 已弃用的 `rand.Seed()`
- 无 CI/CD 配置、无 Makefile、无 linter 配置
- `go.sum` 被 `.gitignore` 排除，不符合 Go 模块最佳实践

## 命令

```bash
# 安装
go get github.com/hoshi200/tencent-im

# 运行测试（需要有效的腾讯云IM凭证）
go test ./...

# 构建
go build ./...
```
