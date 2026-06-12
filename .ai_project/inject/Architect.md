# Architect — 项目注入

版本 1.0 · 关联 constitution.md, standards.md

---

## 技术栈

| 项 | 选型 |
|----|------|
| 语言 | Go 1.24.4 |
| 包管理器 | Go Modules |
| 测试 | `go test` (std) + `-race` + `-coverprofile` |
| 格式化 | gofmt + goimports + gofumpt v0.7.0 |
| 类型检查 | go vet + staticcheck + golangci-lint (30+ linters) |
| Lint | golangci-lint v1.64.2 |
| 构建 | make (makefile) + goreleaser + GitHub Actions |

## 架构分层

```
cmd/main.go                  ← 入口，组装依赖图
    │
    ├── adapters/ui/         ← TUI 适配器层 (tview/tcell)
    │       tui.go           ← 主 TUI 结构体，构建器链式装配
    │       server_list.go   ← 服务器列表组件
    │       server_form.go   ← 添加/编辑表单（5 选项卡）
    │       server_details.go← 服务器详情面板
    │       search_bar.go    ← 模糊搜索栏
    │       status_bar.go    ← 状态栏
    │       header.go        ← 应用头部
    │       handlers.go      ← 键盘事件处理
    │       validation.go    ← 输入校验（UI 层）
    │       sort.go          ← 排序逻辑
    │       ...
    │
    ├── adapters/data/       ← 数据适配器层
    │       ssh_config_file/ ← SSH Config 持久化
    │           ssh_config_file_repo.go  ← Repository 实现
    │           config_io.go            ← 读写 + 原子写入
    │           crud.go                 ← 增删改逻辑
    │           mapper.go               ← domain ↔ ssh_config 映射
    │           backup.go               ← 备份管理
    │           metadata_manager.go     ← 元数据管理（Pin, LastSeen, SSHCount）
    │           file_system.go          ← 文件系统接口（可 mock）
    │
    ├── core/                ← 核心层
    │   ├── domain/          ← 领域模型
    │   │       server.go    ← Server struct（117 行，120+ SSH 字段）
    │   ├── ports/           ← 接口定义
    │   │       repositories.go  ← ServerRepository interface
    │   │       services.go      ← ServerService interface
    │   └── services/        ← 业务服务实现
    │           server_service.go ← SSH 执行、端口转发、Ping、校验
    │
    └── logger/              ← 基础设施
            logger.go        ← zap 封装（输出至 ~/.lazyssh/lazyssh.log）
```

**依赖方向（单向）：**
```
adapters → core/services → core/ports → core/domain
     ↑
adapters ───→ core/ports → core/domain
```

## 架构约束

| 编号 | 约束 | 理由 |
|------|------|------|
| C-001 | UI 层不直接依赖 `ssh_config` 库 | 保持 TUI 与数据源解耦，通过 `ports.ServerService` 接口 |
| C-002 | 数据适配层仅实现 `core/ports` 接口 | 保证仓储可替换（如未来支持其他 SSH 管理方式） |
| C-003 | domain.Server 为纯数据结构，不含业务逻辑 | 领域模型只承载数据，业务逻辑在 service 层 |
| C-004 | 所有文件操作必须通过 `FileSystem` 接口 | 便于单元测试 mock 文件系统 |
| C-005 | 端口转发进程管理在 service 层 | goroutine 生命周期管理归属业务逻辑，不泄漏到 UI |
| C-006 | 新增 SSH 配置字段需同步 domain、mapper、form | 三层一致性：领域模型 ↔ 持久化映射 ↔ UI 表单 |

## 设计模式补充

constitution.md D2 已定义事件驱动、Repository 模式、构建器模式。本项目额外使用：

#### 构建器链
TUI 组件通过链式调用装配：`initializeTheme().buildComponents().buildLayout().bindEvents().loadInitialData()`。每步返回 `*tui` 自身，任一阶段失败通过 error 传播。

#### 回调注入
组件通过 `OnXxx(func)` 方法注入回调：
- `searchBar.OnSearch().OnEscape().OnNavigate()`
- `serverList.OnSelectionChange().OnReturnToSearch()`

回调内不持有组件引用，避免循环依赖。

## 关键文件

| 文件 | 用途 |
|------|------|
| `cmd/main.go` | 依赖组装入口 |
| `internal/core/domain/server.go` | Server 领域模型 |
| `internal/core/ports/repositories.go` | ServerRepository 接口 |
| `internal/core/ports/services.go` | ServerService 接口 |
| `internal/core/services/server_service.go` | 业务逻辑实现 |
| `internal/adapters/ui/tui.go` | TUI 主控制器 |
| `internal/adapters/data/ssh_config_file/ssh_config_file_repo.go` | SSH 配置持久化 |
| `internal/adapters/data/ssh_config_file/mapper.go` | domain ↔ ssh_config 映射 |
| `internal/logger/logger.go` | 日志基础设施 |

## 文档约定

| 类型 | 目录 | 格式 |
|------|------|------|
| 架构文档 | `docs/architecture/` | {模块名}.md |
| 决策记录 | `docs/decisions/` | ADR-{4位序号}-{标题}.md |
| 工作笔记 | `docs/worknotes/architect/` | YYYY-MM-DD.md |

## ADR 格式

状态(Draft/Active/Deprecated) + 关联 PRD + 决策清单(D{序号}:决策/备选/理由) + 修订记录

## ADR 编号规则

1. 查看 `docs/decisions/` 目录下已有 ADR 文件
2. 取最大序号 +1，补零至 4 位
3. 示例：已有 ADR-0003-xxx.md → 新文件为 ADR-0004-yyy.md

## 质量门禁

- 架构变更必有 ADR
- 无代码提交
- 审查点：依赖方向、接口一致性、constitution.md 合规

## 已知技术债务

暂无

## 历史教训

暂无
