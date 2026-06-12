# 编程宪法

版本 1.0 · 所有 Agent 遵守 · 仅架构师可修改
违反即不通过。

---

## 零、核心哲学

### D0：优雅优于兼容

本项目处于早期开发阶段（v0.1.0），不考虑向后兼容。架构决策优先健壮性与一致性。代码以简洁、可测试为第一优先级，不为兼容旧版本保留冗余逻辑。

---

## 一、异常体系

### D1：异常优于返回码

用 Go 错误值表达所有可恢复错误，禁止吞没错误或返回 nil 表示成功/失败模糊态。

错误包装规则：

| 场景 | 规则 |
|------|------|
| 可恢复预期错误 | 返回 error，调用方决定是否继续 |
| 不可恢复错误 | `os.Exit(1)` 或 `panic`（仅在初始化阶段） |
| 外部调用失败 | `fmt.Errorf("xxx: %w", err)` 保留错误链 |
| 边界输入校验 | 入口层返回明确错误信息 |
| goroutine 错误 | 必通过 channel/log 捕获，不可静默丢弃 |

禁止：
- 裸 `_` 忽略 error
- `fmt.Println` 代 logging
- 返回 nil 表错误
- 底层吞异常

---

## 二、设计模式

### D2：模式约束

架构文档引用以下模式时，本节约束自动生效。架构师新增模式须在此补充。

#### 事件驱动
通过回调函数（OnSearch、OnEscape 等）实现 fire-and-forget。handler 自行处理错误并 log，禁止魔法字符串。

#### 插件化
（本项目暂无插件系统，保留约束位）

#### 管线化
（构建器模式：`initializeTheme().buildComponents().buildLayout().bindEvents().loadInitialData()`）每阶段单一职责，失败在对应阶段返回 error。

#### Repository 模式
数据访问统一通过 `ports.ServerRepository` 接口，服务层不直接依赖文件系统或 ssh_config 库。仓储实现位于 `adapters/data/` 下。

---

## 三、依赖规则

### D3：单向自上而下

```
cmd/main.go          ← 组装层
    ↓
adapters/ui/         ← TUI 适配器
adapters/data/       ← 持久化适配器
    ↓
core/services/       ← 业务服务
    ↓
core/ports/          ← 接口定义
    ↓
core/domain/         ← 领域模型
```

- 上层可依赖下层，下层禁止依赖上层
- 循环依赖 = 架构缺陷，必须消除
- 跨层引用必须通过 `core/ports/` 接口

---

## 四、安全基线

### D4：不可违反

- 输入校验在边界层完成（`server_service.go:validateServer()`）
- 敏感信息禁出现在日志/错误/返回值中
- 密钥只通过 `~/.ssh/config` 或环境变量注入，禁硬编码
- 禁 `os/exec` 处理不可信用户输入（SSH 参数由用户通过 TUI 表单提供，非外部输入）
- 生产环境应关闭 Debug 日志级别（当前 `zap.DebugLevel` 仅为开发期设置）
- 文件写入使用原子操作（写临时文件 → rename）

---

## 五、配置管理

### D5：优先级

命令行参数（cobra flags） > 环境变量 > SSH 配置文件（`~/.ssh/config`）> 代码默认值

所有配置须有默认值（密钥类除外）。运行时不可修改配置路径（`configPath` 在 main 组装后冻结）。

---

## 六、日志级别

### D6：级别标准

| 级别 | 场景 |
|------|------|
| DEBUG | 开发调试（当前默认，生产应改为 INFO） |
| INFO | 关键业务节点（SSH 开始/结束、服务启动） |
| WARN | 可恢复异常、降级、配置缺失有默认值 |
| ERROR | 操作失败但进程运行（含 stack trace 上下文） |
| FATAL | 进程无法继续（已用 `os.Exit(1)`） |

日志输出目标：`~/.lazyssh/lazyssh.log`（生产 JSON 格式 + ISO8601 时间戳）

禁止：
- 循环内 INFO+
- 日志含敏感信息（密钥路径、密码、私钥内容）
- `fmt.Println` 代 logging

---

## 修订记录

| 版本 | 日期 | 说明 |
|------|------|------|
| 1.0 | 2026-06-12 | 初稿：7 项决策(D0-D6) |
