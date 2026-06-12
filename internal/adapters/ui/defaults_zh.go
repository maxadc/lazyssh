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

// zhFieldPlaceholders returns Chinese placeholder texts for form fields
var zhFieldPlaceholders = map[string]string{
	// Required fields
	"Alias": "必填",
	"Host":  "必填",

	// Fields that show default value in placeholder
	"Port":            "默认: 22",
	"User":            "默认: 当前用户名",
	"ConnectTimeout":  "秒数 (默认: 无)",
	"ConnectionAttempts": "默认: 1",
	"ServerAliveInterval": "秒数 (默认: 0)",
	"ServerAliveCountMax":   "默认: 3",
	"NumberOfPasswordPrompts": "默认: 3",
	"CanonicalizeMaxDots":   "默认: 1",
	"IPQoS":               "默认: af21 cs1",
	"EscapeChar":          "默认: ~",
	"IdentityAgent":       "默认: SSH_AUTH_SOCK",
	"UserKnownHostsFile":  "默认: ~/.ssh/known_hosts",

	// Fields that show examples in placeholder
	"Keys":       "例如: ~/.ssh/id_rsa, ~/.ssh/id_ed25519",
	"Tags":       "逗号分隔的标签",
	"ProxyJump":  "例如: bastion.example.com",
	"ProxyCommand": "例如: ssh -W %h:%p jump.example.com",
	"RemoteCommand": "例如: tmux attach",
	"LocalForward":    "例如: 8080:localhost:80, 3000:localhost:3000",
	"RemoteForward":   "例如: 80:localhost:8080",
	"DynamicForward":  "例如: 1080, 1081",
	"ControlPath":     "例如: ~/.ssh/master-%r@%h:%p",
	"ControlPersist":  "例如: 10m, 4h, 是, 否",
	"PreferredAuthentications": "例如: publickey,password",
	"PubkeyAcceptedAlgorithms":  "算法 (支持 +/-/^ 前缀)",
	"HostbasedAcceptedAlgorithms": "算法 (支持 +/-/^ 前缀)",
	"HostKeyAlgorithms":         "算法 (支持 +/-/^ 前缀)",
	"Ciphers":                   "算法 (支持 +/-/^ 前缀)",
	"MACs":                      "算法 (支持 +/-/^ 前缀)",
	"KexAlgorithms":             "算法 (支持 +/-/^ 前缀)",
	"BindAddress":               "IP, 主机名, * (全部), 或 localhost",
	"CanonicalDomains":          "例如: example.com, internal.net",
	"CanonicalizePermittedCNAMEs": "例如: *.example.com:example.net",
	"LocalCommand":              "例如: echo '已连接到 %h'",
	"SendEnv":                   "例如: LANG, LC_*, TERM",
	"SetEnv":                    "例如: FOO=bar, DEBUG=1",
}

// GetFieldPlaceholder zh build version
func GetFieldPlaceholder(fieldName string) string {
	if placeholder, exists := zhFieldPlaceholders[fieldName]; exists {
		return placeholder
	}
	return ""
}

// zhFormSectionHeaders returns Chinese form section header strings
var zhFormSectionHeaders = map[string]string{
	"Proxy & Command":          "▶ 代理 & 命令",
	"Connection Settings":      "▶ 连接设置",
	"Bind Options":             "▶ 绑定选项",
	"Hostname Canonicalization": "▶ 主机名规范化",
	"Keep-Alive":               "▶ 保活",
	"Multiplexing":             "▶ 多路复用",
	"Port Forwarding":          "▶ 端口转发",
	"Agent & X11 Forwarding":   "▶ Agent & X11转发",
	"Public Key Authentication": "▶ 公钥认证",
	"SSH Agent":                "▶ SSH Agent",
	"Password & Interactive":   "▶ 密码 & 交互",
	"Security":                 "▶ 安全",
	"Advanced":                 "▶ 高级",
	"Cryptography":             "▶ 密码学",
	"Command Execution":        "▶ 命令执行",
	"Environment":              "▶ 环境",
}

// GetFormSectionHeader zh build version - returns Chinese section header if available
func GetFormSectionHeader(key string) string {
	if zh, exists := zhFormSectionHeaders[key]; exists {
		return zh
	}
	return "▶ " + key
}

// zhFormFieldLabels returns Chinese form field label suffixes
var zhFormFieldLabels = map[string]string{
	"Alias:":            "别名:",
	"Host/IP:":          "主机/IP:",
	"User:":             "用户:",
	"Port:":             "端口:",
	"Keys:":             "密钥:",
	"Tags:":             "标签:",
	"AuthMethod:":       "认证方式:",
	"Password:":         "密码:",
	"ProxyJump:":        "代理跳板:",
	"ProxyCommand:":     "代理命令:",
	"RemoteCommand:":    "远程命令:",
	"RequestTTY:":       "请求TTY:",
	"SessionType:":      "会话类型:",
	"ConnectTimeout:":   "连接超时:",
	"ConnectionAttempts:": "连接尝试:",
	"IPQoS:":            "IP服务质量:",
	"BatchMode:":        "批处理模式:",
	"BindAddress:":      "绑定地址:",
	"BindInterface:":    "绑定接口:",
	"AddressFamily:":    "地址族:",
	"CanonicalizeHostname:": "规范化主机名:",
	"CanonicalDomains:":   "规范域名:",
	"CanonicalizeFallbackLocal:": "规范化回退本地:",
	"CanonicalizeMaxDots:":  "规范化最大点数:",
	"CanonicalizePermittedCNAMEs:": "允许的CNAME:",
	"ControlMaster:":      "控制主连接:",
	"ControlPath:":        "控制路径:",
	"ControlPersist:":     "控制持久化:",
	"Compression:":        "压缩:",
	"TCPKeepAlive:":       "TCP保活:",
	"ServerAliveInterval:": "服务器存活间隔:",
	"ServerAliveCountMax:": "服务器存活计数上限:",
	"PubkeyAuthentication:": "公钥认证:",
	"IdentitiesOnly:":    "仅使用指定密钥:",
	"AddKeysToAgent:":    "添加密钥到Agent:",
	"IdentityAgent:":     "认证Agent:",
	"PasswordAuthentication:": "密码认证:",
	"KbdInteractiveAuthentication:": "键盘交互认证:",
	"ForwardAgent:":      "转发Agent:",
	"ForwardX11:":        "转发X11:",
	"ForwardX11Trusted:": "可信X11转发:",
	"ClearAllForwardings:": "清除所有转发:",
	"ExitOnForwardFailure:": "转发失败退出:",
	"GatewayPorts:":      "网关端口:",
	"LocalForward:":      "本地转发:",
	"RemoteForward:":     "远程转发:",
	"DynamicForward:":    "动态转发:",
	"StrictHostKeyChecking:": "严格主机密钥检查:",
	"CheckHostIP:":       "检查主机IP:",
	"FingerprintHash:":   "指纹哈希:",
	"UserKnownHostsFile:": "用户已知主机文件:",
	"HostKeyAlgorithms:": "主机密钥算法:",
	"MACs:":              "MAC算法:",
	"Ciphers:":           "加密算法:",
	"KexAlgorithms:":     "密钥交换算法:",
	"VerifyHostKeyDNS:":  "DNS验证主机密钥:",
	"UpdateHostKeys:":    "更新主机密钥:",
	"HashKnownHosts:":    "哈希已知主机:",
	"VisualHostKey:":     "可视化主机密钥:",
	"LocalCommand:":      "本地命令:",
	"PermitLocalCommand:": "允许本地命令:",
	"EscapeChar:":        "转义字符:",
	"SendEnv:":           "发送环境变量:",
	"SetEnv:":            "设置环境变量:",
	"LogLevel:":          "日志级别:",
	"NumberOfPasswordPrompts:": "密码提示次数:",
	"PreferredAuthentications:": "首选认证方式:",
	"PubkeyAcceptedAlgorithms:": "公钥接受算法:",
	"HostbasedAcceptedAlgorithms:": "主机认证接受算法:",
}

// GetFormLabel zh build version - returns Chinese field label if available
func GetFormLabel(key string) string {
	if zh, exists := zhFormFieldLabels[key]; exists {
		return zh
	}
	return key
}

// zhTabAbbrevs returns Chinese tab abbreviations
var zhTabAbbrevs = map[string]string{
	"Basic":     "基本",
	"Connection": "连接",
	"Forwarding": "转发",
	"Authentication": "认证",
	"Advanced":  "高级",
}

// GetTabAbbrev zh build version
func GetTabAbbrev(fullName string) string {
	if zh, exists := zhTabAbbrevs[fullName]; exists {
		return zh
	}
	return fullName
}

// zhDetailGroupNames returns Chinese detail panel group names
var zhDetailGroupNames = map[string]string{
	"Connection & Proxy":     "连接 & 代理",
	"Authentication":         "认证",
	"Forwarding":             "转发",
	"Security & Cryptography": "安全 & 密码学",
	"Environment & Execution": "环境 & 执行",
	"Debugging":              "调试",
}

// GetDetailGroupName zh build version
func GetDetailGroupName(key string) string {
	if zh, exists := zhDetailGroupNames[key]; exists {
		return zh
	}
	return key
}

// zhNoHelpAvailable returns zh text for "no help available"
func GetNoHelpText() string {
	return "该字段暂无帮助信息"
}

// zhExamplePrefix returns zh text for "e.g."
func GetExamplePrefix() string {
	return "例如: "
}

// zhAvailableSince returns zh text for "Available since"
func GetAvailableSince() string {
	return "OpenSSH 版本: "
}

// GetAuthMethodDefault returns the default auth method for zh build.
func GetAuthMethodDefault() string {
	return ""
}

// GetPasswordDefault returns the default password value for zh build.
func GetPasswordDefault() string {
	return ""
}
