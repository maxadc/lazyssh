# Interface Engineer — 项目注入

版本 1.0 · 关联 constitution.md, standards.md

---

## 项目概况

**lazyssh** — Terminal-based SSH manager。通信方式为**同进程函数调用**（Go 包内直接调用接口方法），无跨进程通信。SSH 连接通过 `os/exec` 调用系统 `ssh` 二进制。

## 任务单存放

任务单按 ADR 分目录，避免扁平堆积：

```
docs/tasks/
├── ADR-0001/                    ← 对应 ADR-0001
│   ├── TASK-01-{子模块}.md
│   └── TASK-02-{子模块}.md
└── ADR-0002/                    ← 对应 ADR-0002
    └── TASK-01-{子模块}.md
```

| 类型 | 目录 | 格式 |
|------|------|------|
| 开发任务单 | `docs/tasks/{ADR编号}/` | TASK-{2位序号}-{子模块名}.md |

## 任务单编号规则

1. 从架构师产出中确认本次 ADR 编号（如 ADR-0001）
2. 创建子目录 `docs/tasks/ADR-0001/`
3. 查看该子目录下已有任务单，取最大序号 +1，补零至 2 位
4. 大模块拆分子任务单时，共享同序号后缀：`TASK-01-用户服务.md`、`TASK-01-用户服务-仓储.md`

## 通信约定

| 约定 | 内容 |
|------|------|
| 内部通信 | Go 包导入 + 接口方法直调 |
| 外部通信 | `os/exec.Command("ssh", args...)` 调用系统 SSH 客户端 |
| 错误响应 | Go `error` 返回值 + 错误链包装 (`%w`) |
| 并发 | goroutine + `sync.Mutex`，错误通过 log 记录 |
| 事件 | 回调函数注入 (`OnSearch`, `OnSelectionChange` 等)，fire-and-forget 模式 |

## 核心模块接口

| 模块 | 路径 | 核心接口 |
|------|------|----------|
| TUI 主控 | `internal/adapters/ui/tui.go` | `App` interface (`Run() error`) |
| 服务器列表 | `internal/adapters/ui/server_list.go` | `ServerList` struct |
| 服务器表单 | `internal/adapters/ui/server_form.go` | `ServerForm` struct（5 选项卡） |
| 搜索栏 | `internal/adapters/ui/search_bar.go` | `SearchBar` struct |
| 业务服务 | `internal/core/ports/services.go` | `ServerService` interface（9 方法） |
| 数据仓储 | `internal/core/ports/repositories.go` | `ServerRepository` interface（6 方法） |
| SSH 仓储实现 | `internal/adapters/data/ssh_config_file/ssh_config_file_repo.go` | `Repository` struct |
| 日志 | `internal/logger/logger.go` | `New(service string) (*zap.SugaredLogger, error)` |

## 数据模型

| 模型 | 位置 | 说明 |
|------|------|------|
| `domain.Server` | `internal/core/domain/server.go` | 核心领域模型，120+ 字段映射 SSH 配置 |
| `ServerMetadata` | `internal/adapters/data/ssh_config_file/metadata_manager.go` | 运行时元数据（Pin, LastSeen, SSHCount） |
| 各 UI 组件 struct | `internal/adapters/ui/*.go` | TUI 组件状态持有者 |

## 设计模式指引

| 模式 | 适用场景 |
|------|---------|
| Repository | 所有 `~/.ssh/config` 读写操作通过 `ServerRepository` 接口 |
| 构建器链 | TUI 组件装配（`Initialize → Build → Layout → Events → Data`） |
| 回调注入 | UI 事件处理（搜索、选择、按键） |
| 策略模式 | 排序逻辑（按别名/最近连接，`SortMode` 枚举） |

## 项目约束

- **接口优先**：新功能必须先定义在 `core/ports/` 接口中，再实现
- **单向依赖**：`adapters → services → ports → domain`，不可反向
- **文件系统抽象**：所有文件 I/O 通过 `FileSystem` 接口，便于测试 mock
- **零值安全**：struct 字段不使用指针类型，零值即有效默认
- **TUI 组件隔离**：各 UI 组件不直接引用其他组件，通过回调通信

## 历史教训

暂无
