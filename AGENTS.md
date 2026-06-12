# lazyssh 项目手册

项目：lazyssh — Terminal-based interactive SSH manager (TUI)

## 强制公约

- **语言**：注释中文、标识符英文（Go 语言标准）
- **代码**：gofmt + goimports + gofumpt 格式化，golangci-lint 全量检查
- **环境**：Go Modules 管理依赖；日志输出至 `~/.lazyssh/lazyssh.log`
- **测试**：TDD 红线（需求→GWT→失败测试→最小实现→重构），禁止跳过
- **目录**：
  - `cmd/` — 入口 & 依赖组装
  - `internal/core/domain/` — 领域模型
  - `internal/core/ports/` — 接口定义
  - `internal/core/services/` — 业务服务
  - `internal/adapters/ui/` — TUI 适配器
  - `internal/adapters/data/` — 持久化适配器
  - `internal/logger/` — 日志

## 公共契约

所有 Agent 必须遵守，违反即不通过：

| 文件 | 定位 | 维护者 |
|------|------|--------|
| `docs/constitution.md` | 编程哲学（怎么思考软件） | 架构师 |
| `docs/standards.md` | 编码规范（代码长什么样） | 架构师 |
| `.ai_project/inject/*.md` | 各角色项目注入配置 | 对应角色 |

## 四角色工作流

```
初始化小助理 → 用户需求 → PM → PRD → Architect → ADR → InterfaceEngineer → 任务单 → Developer → 代码+测试
```

| 角色 | 产出 | 全局 Prompt |
|------|------|-------------|
| PM | 产品文档(PRD) | `/root/.opencode/agents/PM.md` |
| Architect | 架构文档+ADR | `/root/.opencode/agents/Architect.md` |
| InterfaceEngineer | 开发任务单 | `/root/.opencode/agents/InterfaceEngineer.md` |
| Developer | 代码+测试 | `/root/.opencode/agents/Developer.md` |

每个角色启动时读取：
1. `docs/constitution.md` + `docs/standards.md`（公共契约）
2. `.ai_project/inject/{角色}.md`（项目注入）
3. 上游产出物（PRD / ADR / 任务单）

## 角色加载

你是多 Agent 系统中的一个角色。**开始任何工作前，必须先读取你的全局 Prompt**：

- **初始化小助理** → `/root/.opencode/agents/InitAssistant.md`
- **产品经理** → `/root/.opencode/agents/PM.md`
- **架构师** → `/root/.opencode/agents/Architect.md`
- **接口工程师** → `/root/.opencode/agents/InterfaceEngineer.md`
- **开发者** → `/root/.opencode/agents/Developer.md`

## 跨角色审核

产出物在交付前须经相关角色审核，审核格式统一：

| 审核方 | 审核对象 | 关注点 |
|--------|---------|--------|
| PM | 所有阶段产出 | 需求对齐度、体验完整性 |
| Architect | ADR/代码 | 架构合规、依赖方向、模式引用 |
| InterfaceEngineer | 代码实现 | 接口签名一致性、异常处理、数据模型 |
| Developer | 任务单/架构设计 | 可实现性、成本合理性、边界条件 |

审核输出格式：

```
@{角色名} 审核反馈
| 检查项 | 结果 | 备注 |
|--------|------|------|
结果：✅通过 / ⚠️遗漏(具体项) / 🚫不一致(具体差异)
```

## 文档导航

| 文档 | 用途 |
|------|------|
| `docs/constitution.md` | 编程宪法 |
| `docs/standards.md` | 编码规范 |
| `docs/status.md` | 项目状态(PM维护) |
| `docs/MEMORY_BANK.md` | 项目记忆库 |
| `.ai_project/inject/*.md` | 各角色项目配置 |
| `docs/product/` | PRD |
| `docs/architecture/` | 架构文档 |
| `docs/decisions/` | 决策记录(ADR) |
| `docs/tasks/` | 开发任务单 |
| `docs/worknotes/architect/` | 架构师工作笔记 |

## 通用行为准则

- **不能谄媚**：用户方案有误时必须指明原因
- **极简交付**：只写最少必要代码
- **精准修改**：只改目标代码，不顺便优化
- **目标驱动**：转为可验证指标，循环直到通过

## 架构边界

- **单向依赖**：`adapters → core/services → core/ports → core/domain`，不可反向
- **接口隔离**：所有跨层调用通过 `core/ports/` 接口，UI 不碰 `ssh_config` 库
- **文件系统抽象**：所有 I/O 通过 `FileSystem` 接口，生产用 `DefaultFileSystem`，测试可 mock
- **零值安全**：`domain.Server` 不使用指针字段，零值即有效默认

---
