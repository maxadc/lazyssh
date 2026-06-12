// Copyright 2025.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build zh

package ui

var TZ = I18n{
	// Sort modes
	AliasAsc:            "别名 ↑",
	AliasDesc:           "别名 ↓",
	LastSeenAsc:         "最近 SSH ↑",
	LastSeenDesc:        "最近 SSH ↓",

	// Search & list
	SearchLabel:    " 🔍 搜索: ",
	SearchTitle:    " 搜索 ",
	ServersTitle:   " 服务器 ",
	DetailsTitle:   " 详情 ",
	HelpTitle:      " 帮助 ",
	AddServerTitle: "添加服务器",
	EditServerTitle: "编辑服务器",
	DeleteConfirmTitle: "删除服务器",
	EditTagsTitle:    " 编辑标签: %s ",
	PortForwardTitle: " 端口转发: %s ",
	NeverSeen:        "从未",
	NoServersMatch:   "没有服务器匹配当前筛选条件。",
	SaveFailed:       "保存失败: %v",
	Pinging:          "正在 Ping %s…",
	PingDown:         "Ping %s: 离线",
	PingUp:           "Ping %s: 在线 (%s)",
	Refreshing:       "正在刷新…",
	RefreshFailed:    "刷新失败: %v",
	Refreshed:        "已刷新 %d 台服务器",
	Copied:           "已复制: %s",
	CopyFailed:       "复制到剪贴板失败",
	TagsUpdated:      "标签已更新",
	TagsLabel:        "标签（逗号分隔）:",
	PortForwardStarted: "端口转发已启动（pid %d）",
	ForwardFailed:    "转发失败: %s",
	StoppedForwarding: "已停止 %s 的转发",
	StopForwardFailed: "停止转发失败: %s",
	InvalidPort:      "无效端口: %s",
	InvalidBindAddress: "无效绑定地址: %s",
	InvalidHost:      "无效主机: %s",
	InvalidHostPort:  "无效主机端口: %s",
	StartingPortForward: "正在启动端口转发…",

	// Detail panel labels
	BasicSettings:    "基本设置:",
	AdvancedSettings: "高级设置:",
	Commands:         "命令:",
	HostLabel:        "主机",
	UserLabel:        "用户",
	PortLabel:        "端口",
	KeyLabel:         "密钥",
	TagsLabelDetail:  "标签",
	PinnedLabel:      "已固定",
	LastSSHLabel:     "最近 SSH",
	SSHCountLabel:    "SSH 次数",
	AliasLabel:       "别名",

	// Section headers (detail panel grouping)
	ConnectionProxy:     "连接 & 代理",
	Authentication:      "认证",
	Forwarding:          "转发",
	SecurityCrypto:      "安全 & 密码学",
	EnvironmentExec:     "环境 & 执行",
	Debugging:           "调试",

	// Form tabs
	FormTabBasic:          "基本",
	FormTabConnection:     "连接",
	FormTabForwarding:     "转发",
	FormTabAuthentication: "认证",
	FormTabAdvanced:       "高级",

	// Form hints
	FormHintNavigate: "^H/^L 切换选项卡  • ^S 保存  • Esc 取消",
	FormHintSave:     "保存",
	FormHintCancel:   "取消",

	// Status bar shortcuts
	StatusNavigate:    "导航",
	StatusSSH:         "SSH",
	StatusForward:     "转发",
	StatusStopForward: "停止转发",
	StatusCopySSH:     "复制 SSH",
	StatusAdd:         "添加",
	StatusEdit:        "编辑",
	StatusPing:        "Ping",
	StatusDelete:      "删除",
	StatusPin:         "固定/取消",
	StatusSearch:      "搜索",
	StatusQuit:        "退出",

	// Commands section in details panel
	CommandsTitle:     "命令:",
	CommandsSSH:       "Enter: SSH 连接",
	CommandsForward:   "f: 端口转发",
	CommandsStopForward: "x: 停止转发",
	CommandsCopySSH:   "c: 复制 SSH 命令",
	CommandsPing:      "g: Ping 服务器",
	CommandsRefresh:   "r: 刷新列表",
	CommandsAdd:       "a: 添加新服务器",
	CommandsEdit:      "e: 编辑条目",
	CommandsTags:      "t: 编辑标签",
	CommandsDelete:    "d: 删除条目",
	CommandsPin:       "p: 固定/取消固定",
	CommandsSearch:    "/: 搜索",

	// Status bar labels
	StatusLabel:     "别名:",
	HostIPLabel:     "主机/IP:",
	UserLabelForm:   "用户:",
	PortLabelForm:   "端口:",
	KeysLabelForm:   "密钥:",
	TagsLabelForm:   "标签:",
	AuthMethodLabel: "认证方式:",
	PasswordLabel:   "密码:",
	AuthMethodAuto:  "自动",
	AuthMethodKey:   "密钥",
	AuthMethodPassword: "密码",
	SSHPassNotFound: "PATH 中未找到 sshpass",
	PasswordAuthFailed: "密码认证失败: %v",
	SSHConnectionFailed: "SSH 连接失败: %v",

	// Dropdown option values
	YesOption:       "是",
	NoOption:        "否",
	DefaultOption:   "默认",
	NoneOption:      "无",
	AskOption:       "询问",
	AutoOption:      "自动",
	ForceOption:     "强制",
	AnyOption:       "任意",
	InetOption:      "inet",
	Inet6Option:     "inet6",
	AlwaysOption:    "始终",

	// Proxy & Connection labels
	ProxyJumpLabel:            "代理跳板:",
	ProxyCommandLabel:         "代理命令:",
	RemoteCommandLabel:        "远程命令:",
	RequestTTYLabel:           "请求TTY:",
	SessionTypeLabel:          "会话类型:",
	ConnectTimeoutLabel:       "连接超时:",
	ConnectionAttemptsLabel:   "连接尝试:",
	IPQoSLabel:                "IP服务质量:",
	BatchModeLabel:            "批处理模式:",
	BindAddressLabel:          "绑定地址:",
	BindInterfaceLabel:        "绑定接口:",
	AddressFamilyLabel:        "地址族:",
	CanonicalizeHostnameLabel: "规范化主机名:",
	CanonicalDomainsLabel:     "规范域名:",
	CanonicalizeFallbackLocalLabel: "规范化回退本地:",
	CanonicalizeMaxDotsLabel:  "规范化最大点数:",
	CanonicalizePermittedCNAMEsLabel: "允许的CNAME:",
	ServerAliveIntervalLabel:  "服务器存活间隔:",
	ServerAliveCountMaxLabel:  "服务器存活计数上限:",
	CompressionLabel:          "压缩:",
	TCPKeepAliveLabel:         "TCP保活:",
	ControlMasterLabel:        "控制主连接:",
	ControlPathLabel:          "控制路径:",
	ControlPersistLabel:       "控制持久化:",

	// Port forwarding labels
	LocalForwardLabel:         "本地转发:",
	RemoteForwardLabel:        "远程转发:",
	DynamicForwardLabel:       "动态转发:",
	ClearAllForwardingsLabel:  "清除所有转发:",
	ExitOnForwardFailureLabel: "转发失败退出:",
	GatewayPortsLabel:         "网关端口:",
	ForwardAgentLabel:         "转发Agent:",
	ForwardX11Label:           "转发X11:",
	ForwardX11TrustedLabel:    "可信X11转发:",

	// Authentication labels
	PubkeyAuthenticationLabel:       "公钥认证:",
	IdentitiesOnlyLabel:             "仅使用指定密钥:",
	AddKeysToAgentLabel:             "添加密钥到Agent:",
	IdentityAgentLabel:              "认证Agent:",
	PasswordAuthenticationLabel:     "密码认证:",
	KbdInteractiveAuthenticationLabel: "键盘交互认证:",
	NumberOfPasswordPromptsLabel:    "密码提示次数:",
	PreferredAuthenticationsLabel:   "首选认证方式:",
	PubkeyAcceptedAlgorithmsLabel:   "公钥接受算法:",
	HostbasedAcceptedAlgorithmsLabel: "主机认证接受算法:",

	// Security labels
	StrictHostKeyCheckingLabel: "严格主机密钥检查:",
	CheckHostIPLabel:          "检查主机IP:",
	FingerprintHashLabel:      "指纹哈希:",
	UserKnownHostsFileLabel:   "用户已知主机文件:",
	HostKeyAlgorithmsLabel:    "主机密钥算法:",
	MACsLabel:                 "MAC算法:",
	CiphersLabel:              "加密算法:",
	KexAlgorithmsLabel:        "密钥交换算法:",
	VerifyHostKeyDNSLabel:     "DNS验证主机密钥:",
	UpdateHostKeysLabel:       "更新主机密钥:",
	HashKnownHostsLabel:       "哈希已知主机:",
	VisualHostKeyLabel:        "可视化主机密钥:",

	// Command & environment labels
	LocalCommandLabel:       "本地命令:",
	PermitLocalCommandLabel: "允许本地命令:",
	EscapeCharLabel:         "转义字符:",
	SendEnvLabel:            "发送环境变量:",
	SetEnvLabel:             "设置环境变量:",
	LogLevelLabel:           "日志级别:",

	// Buttons
	SaveButton:   "保存",
	CancelButton: "取消",
	StartButton:  "启动",
	DeleteButton: "删除",
	CloseButton:  "关闭",

	// Sort
	SortLabel: " 排序: %s ",
	TitleSort: " 服务器 — 排序: %s ",

	SortAscSuffix:  " ↑",
	SortDescSuffix: " ↓",

	// Section headers (form)
	SectionProxyCommand:          "▶ 代理 & 命令",
	SectionConnectionSettings:    "▶ 连接设置",
	SectionBindOptions:           "▶ 绑定选项",
	SectionHostnameCanonicalization: "▶ 主机名规范化",
	SectionKeepAlive:             "▶ 保活",
	SectionMultiplexing:          "▶ 多路复用",
	SectionPortForwarding:        "▶ 端口转发",
	SectionAgentX11Forwarding:    "▶ Agent & X11转发",
	SectionPublicKeyAuth:         "▶ 公钥认证",
	SectionSSHAgent:              "▶ SSH Agent",
	SectionPasswordInteractive:   "▶ 密码 & 交互",
	SectionSecurity:              "▶ 安全",
	SectionCommandExecution:      "▶ 命令执行",
	SectionEnvironment:           "▶ 环境",

	// Port forward form labels
	TypeLabel:                 "类型",
	PortLabelForward:          "端口",
	HostLabelForward:          "主机",
	HostPortLabel:             "主机端口",
	BindAddressOptionalLabel:  "绑定地址（可选）",
	ModeLabel:                 "模式",

	ForwardTypeLocal:    "本地",
	ForwardTypeRemote:   "远程",
	ForwardTypeDynamic:  "动态",
	ForwardModeOnly:     "仅转发",
	ForwardModeForwardSSH: "转发 + SSH",

	DefaultSessionType: "默认",
	NoneSessionType:    "无会话 (-N)",
	SubsystemSessionType: "子系统 (-s)",

	// Delete confirmation
	DeleteConfirmMsg:  "确定删除服务器 %s (%s@%s:%d)？",
	DeleteCannotUndo:  "此操作不可撤销。",

	NoHelpAvailable: "该字段暂无帮助信息",
	ExamplePrefix:   "例如: ",
	AvailableSince:  "OpenSSH 版本: ",

	GruConnectionProxy:  "连接 & 代理",
	GruAuthentication:   "认证",
	GruForwarding:       "转发",
	GruSecurityCrypto:   "安全 & 密码学",
	GruEnvironmentExec:  "环境 & 执行",
	GruDebugging:        "调试",

	ValidationErrorRequired: " 是必填项",
}

func Tr(key string) string {
	return key
}

func (i *I18n) SortModeString(mode SortMode) string {
	return mode.String()
}

func (i *I18n) CommandsText() string {
	return "\n[::b]" + i.CommandsTitle + "[-]\n  " + i.CommandsSSH + "\n  " + i.CommandsForward + "\n  " + i.CommandsStopForward + "\n  " + i.CommandsCopySSH + "\n  " + i.CommandsPing + "\n  " + i.CommandsRefresh + "\n  " + i.CommandsAdd + "\n  " + i.CommandsEdit + "\n  " + i.CommandsTags + "\n  " + i.CommandsDelete + "\n  " + i.CommandsPin
}

func (i *I18n) FieldHelp(field string) *FieldHelp {
	return GetFieldHelpZH(field)
}

// T is the active i18n instance - points to TZ for zh builds
var T = TZ
