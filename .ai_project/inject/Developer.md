# Developer — 项目注入

版本 1.0 · 关联 constitution.md, standards.md

---

## 环境

| 项 | 命令 |
|----|------|
| 安装依赖 | `go mod download && go mod verify` |
| 运行测试 | `go test -race -coverprofile=coverage.out ./...` |
| 格式化 | `make fmt`（gofumpt + go fmt） |
| Lint | `make lint`（golangci-lint） |
| 类型检查 | `go vet ./...` + `make check`（staticcheck） |
| 全量质量 | `make quality`（fmt + vet + lint） |
| 构建 | `make build` |
| 启动项目 | `make run` |
| 覆盖率报告 | `make coverage` |

## 目录

| 目录 | 用途 |
|------|------|
| `cmd/` | 入口（main.go），依赖组装 |
| `internal/core/domain/` | 领域模型，纯数据结构 |
| `internal/core/ports/` | 接口定义，所有适配器和服务依赖于此 |
| `internal/core/services/` | 业务逻辑实现 |
| `internal/adapters/ui/` | TUI 组件（tview） |
| `internal/adapters/data/ssh_config_file/` | SSH Config 持久化适配器 |
| `internal/logger/` | 日志基础设施 |
| `docs/` | 图片资源 |

## 核心依赖

| 依赖 | 用途 |
|------|------|
| `github.com/rivo/tview` | TUI 框架 |
| `github.com/gdamore/tcell/v2` | 终端底层 |
| `github.com/kevinburke/ssh_config` | SSH Config 解析器 |
| `github.com/spf13/cobra` | CLI 框架 |
| `go.uber.org/zap` | 结构化日志 |
| `github.com/atotto/clipboard` | 剪贴板操作 |

## TDD 工作流

1. 读任务单 → 拆 GWT 验收标准
2. 写失败测试（正常/异常/边界/空值），自审质量
3. 最小实现通过测试 → 重构
4. 全量测试（`go test -race ./...`）→ DoD 自检

### 测试命名规范
`Test{被测函数}_{场景}_{期望}`
示例：`TestValidateServer_EmptyAlias_ReturnsError`

### Mock 策略
- 文件系统：注入 `FileSystem` 接口（见 `file_system.go`）
- 外部命令：使用 `exec.Command` 的 mock（提取为接口或使用构建标签）
- SSH 配置库：通过 `ServerRepository` 接口 mock

## DoD 检查清单

- [ ] 新功能有单元测试
- [ ] `go test -race ./...` 全绿
- [ ] `make quality` 全绿（fmt + vet + lint）
- [ ] 关键路径无 `// TODO` / `panic` 占位
- [ ] 公开 API 有 godoc 注释（`// Xxx does Y.`）
- [ ] 错误使用 `fmt.Errorf("...: %w", err)` 包装
- [ ] 对应文档（PRD/ADR/任务单）已更新或关闭
- [ ] `docs/MEMORY_BANK.md` 已记录（如有雷区/经验）

## 质量门禁

```bash
go test -race ./...       # 全部通过
make quality              # fmt + vet + lint 全绿
```

## 历史教训

暂无
