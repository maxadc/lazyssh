# 编码规范

版本 1.0 · 所有代码产出遵守 · 仅架构师可修改

本文件只管"代码长什么样"，不管"代码做什么"。

---

## 一、命名

| 类型 | 规则 | 示例 |
|------|------|------|
| 文件名 | 小写 + 下划线分隔 | `server_service.go`, `ssh_config_file_repo.go` |
| 包名 | 小写单数，无下划线 | `domain`, `ports`, `services`, `ui` |
| 类型（struct/interface） | PascalCase 公开，camelCase 私有 | `ServerService`, `serverService` |
| 函数/方法 | PascalCase 公开，camelCase 私有 | `NewServerService()`, `validateServer()` |
| 变量 | camelCase | `serverRepo`, `sshConfigFile` |
| 常量 | PascalCase（公开）或 camelCase（私有） | `RepoURL`, `MaxBackups` |
| 接口 | 单方法接口加 `-er` 后缀，多方法用名词 | `FileSystem` |
| 测试函数 | `TestXxx` | `TestValidateServer` |
| 测试文件 | `*_test.go` | `crud_test.go`, `validation_test.go` |

禁止：
- 单字母变量（循环索引 `i`, `j`, `k` 除外）
- 无意义命名（`data`, `info`, `temp`, `tmp`, `foo`, `bar`）
- 过度缩写（`usr`, `mgr`, `cfg`, `srv`）
- 包名与目录名不一致

---

## 二、文件结构

按序排列：
1. 版权头（Apache 2.0 许可注释，goheader linter 强制）
2. `package` 声明
3. `import` 分组（标准库 → 项目内部 → 第三方，goimports 自动处理）
4. 类型定义
5. 构造函数（`NewXxx`）
6. 公开方法
7. 私有方法
8. 辅助函数

| 限制 | 值 |
|------|----|
| 单文件 | ≤500 行 |
| 单函数 | ≤50 行 |
| 嵌套 | ≤3 层 |
| 参数 | ≤5 个（超用 struct 封装） |
| 圈复杂度 | ≤10（gocyclo linter 强制） |
| 行宽 | ≤120 字符（lll linter 强制，内部包豁免） |

---

## 三、类型与零值

- 所有公开 API 必须有显式返回类型
- 优先使用具体类型，仅在需要多态时使用接口
- 零值必须有意义：`sync.Mutex` 零值即可用，`nil` map/slice 的 `len()` 返回 0
- struct 字段使用具体类型，避免无意义的指针（`*string`）
- 接口定义在消费方（`core/ports/`），实现方不定义接口

---

## 四、注释与文档

- 公开类型/函数必须有 godoc 注释（revive linter 强制 `comment-spacings`）
- 注释以类型/函数名开头：`// ServerService defines the contract for server operations.`
- 复杂逻辑用行内注释解释 **为什么**，不是 **做什么**
- 版权头：Apache 2.0（goheader linter 自动检查）

---

## 五、Import

分组顺序（goimports 自动处理）：
1. 标准库
2. 项目内部（`github.com/Adembc/lazyssh/...`）
3. 第三方

禁止：
- 循环 import（Go 编译器拒绝）
- 未使用的 import（goimports 自动移除）
- 别名滥用（仅在冲突时使用）

---

## 六、并发

- 使用 `sync.Mutex`/`sync.RWMutex` 保护共享状态
- goroutine 内部错误必须通过 channel 或 log 记录
- `defer` 确保锁释放、资源清理
- 避免共享内存，优先 channel 通信
- `go test -race` 检测数据竞争

---

## 七、测试

命名：`Test{被测函数}_{场景}_{期望}`

示例：`TestValidateServer_EmptyAlias_ReturnsError`

规则：
- 外部依赖须 mock（filesystem、ssh 进程调用）
- 使用 `internal/adapters/data/ssh_config_file/crud_test.go` 为参照
- 测试间禁顺序依赖
- 禁依赖外部真实服务（SSH、网络、文件系统）
- 表格驱动测试优先
- 覆盖率目标：核心逻辑 80%+

---

## 八、错误处理

- error 作为函数最后一个返回值
- 使用 `fmt.Errorf("context: %w", err)` 包装错误，保留调用链
- 使用 `errors.Is()` / `errors.As()` 判断错误类型
- 禁止 `panic`（初始化阶段除外）
- linter 强制：errcheck, errorlint

---

## 修订记录

| 版本 | 日期 | 说明 |
|------|------|------|
| 1.0 | 2026-06-12 | 初稿 |
