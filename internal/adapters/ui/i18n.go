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

//go:build !zh

package ui

var T = I18n{
	AliasAsc:            "Alias ↑",
	AliasDesc:           "Alias ↓",
	LastSeenAsc:         "Last SSH ↑",
	LastSeenDesc:        "Last SSH ↓",
	SearchLabel:         " 🔍 Search: ",
	SearchTitle:           " Search ",
	ServersTitle:          " Servers ",
	DetailsTitle:          " Details ",
	HelpTitle:             " Help ",
	AddServerTitle:        "Add Server",
	EditServerTitle:       "Edit Server",
	DeleteConfirmTitle:    "Delete Server",
	EditTagsTitle:         " Edit Tags: %s ",
	PortForwardTitle:      " Port Forwarding: %s ",
	NeverSeen:             "Never",
	NoServersMatch:        "No servers match the current filter.",
	SaveFailed:            "Save failed: %v",
	Pinging:               "Pinging %s…",
	PingDown:              "Ping %s: DOWN",
	PingUp:                "Ping %s: UP (%s)",
	Refreshing:            "Refreshing…",
	RefreshFailed:         "Refresh failed: %v",
	Refreshed:             "Refreshed %d servers",
	Copied:                "Copied: %s",
	CopyFailed:            "Failed to copy to clipboard",
	TagsUpdated:           "Tags updated",
	TagsLabel:             "Tags (comma):",
	PortForwardStarted:    "Port forwarding started (pid %d)",
	ForwardFailed:         "Forward failed: %s",
	StoppedForwarding:     "Stopped forwarding for %s",
	StopForwardFailed:     "Failed to stop forwarding: %s",
	InvalidPort:           "Invalid port: %s",
	InvalidBindAddress:    "Invalid bind address: %s",
	InvalidHost:           "Invalid host: %s",
	InvalidHostPort:       "Invalid host port: %s",
	StartingPortForward:   "Starting port forward…",

	BasicSettings:    "Basic Settings:",
	AdvancedSettings: "Advanced Settings:",
	Commands:         "Commands:",
	HostLabel:        "Host",
	UserLabel:        "User",
	PortLabel:        "Port",
	KeyLabel:         "Key",
	TagsLabelDetail:  "Tags",
	PinnedLabel:      "Pinned",
	LastSSHLabel:     "Last SSH",
	SSHCountLabel:    "SSH Count",
	AliasLabel:       "Alias",

	ConnectionProxy: "Connection & Proxy",
	Authentication:   "Authentication",
	Forwarding:       "Forwarding",
	SecurityCrypto:   "Security & Cryptography",
	EnvironmentExec:  "Environment & Execution",
	Debugging:        "Debugging",

	FormTabBasic:          "Basic",
	FormTabConnection:     "Connection",
	FormTabForwarding:     "Forwarding",
	FormTabAuthentication: "Authentication",
	FormTabAdvanced:       "Advanced",

	FormHintNavigate: "^H/^L Navigate  • ^S Save  • Esc Cancel",
	FormHintSave:     "Save",
	FormHintCancel:   "Cancel",

	// Status bar shortcut labels
	StatusNavigate:    "Navigate",
	StatusSSH:         "SSH",
	StatusForward:     "Forward",
	StatusStopForward: "Stop Forward",
	StatusCopySSH:     "Copy SSH",
	StatusAdd:         "Add",
	StatusEdit:        "Edit",
	StatusPing:        "Ping",
	StatusDelete:      "Delete",
	StatusPin:         "Pin/Unpin",
	StatusSearch:      "Search",
	StatusQuit:        "Quit",

	// Commands section in details panel
	CommandsTitle:     "Commands:",
	CommandsSSH:       "Enter: SSH connect",
	CommandsForward:   "f: Port forward",
	CommandsStopForward: "x: Stop forwarding",
	CommandsCopySSH:   "c: Copy SSH command",
	CommandsPing:      "g: Ping server",
	CommandsRefresh:   "r: Refresh list",
	CommandsAdd:       "a: Add new server",
	CommandsEdit:      "e: Edit entry",
	CommandsTags:      "t: Edit tags",
	CommandsDelete:    "d: Delete entry",
	CommandsPin:       "p: Pin/Unpin",
	CommandsSearch:    "/: Search",

	StatusLabel:     "Alias:",
	HostIPLabel:     "Host/IP:",
	UserLabelForm:   "User:",
	PortLabelForm:   "Port:",
	KeysLabelForm:   "Keys:",
	TagsLabelForm:   "Tags:",
	AuthMethodLabel: "Auth Method:",
	PasswordLabel:   "Password:",
	AuthMethodAuto:  "auto",
	AuthMethodKey:   "key",
	AuthMethodPassword: "password",
	SSHPassNotFound: "sshpass not found in PATH",
	PasswordAuthFailed: "Password authentication failed: %v",
	SSHConnectionFailed: "SSH connection failed: %v",
	YesOption:       "yes",
	NoOption:        "no",
	DefaultOption:   "default",
	NoneOption:      "none",
	AskOption:       "ask",
	AutoOption:      "auto",
	ForceOption:     "force",
	AnyOption:       "any",
	InetOption:      "inet",
	Inet6Option:     "inet6",
	AlwaysOption:    "always",

	ProxyJumpLabel:            "ProxyJump:",
	ProxyCommandLabel:         "ProxyCommand:",
	RemoteCommandLabel:        "RemoteCommand:",
	RequestTTYLabel:           "RequestTTY:",
	SessionTypeLabel:          "SessionType:",
	ConnectTimeoutLabel:       "ConnectTimeout:",
	ConnectionAttemptsLabel:   "ConnectionAttempts:",
	IPQoSLabel:                "IPQoS:",
	BatchModeLabel:            "BatchMode:",
	BindAddressLabel:          "BindAddress:",
	BindInterfaceLabel:        "BindInterface:",
	AddressFamilyLabel:        "AddressFamily:",
	CanonicalizeHostnameLabel: "CanonicalizeHostname:",
	CanonicalDomainsLabel:     "CanonicalDomains:",
	CanonicalizeFallbackLocalLabel: "CanonicalizeFallbackLocal:",
	CanonicalizeMaxDotsLabel:  "CanonicalizeMaxDots:",
	CanonicalizePermittedCNAMEsLabel: "CanonicalizePermittedCNAMEs:",
	ServerAliveIntervalLabel:  "ServerAliveInterval:",
	ServerAliveCountMaxLabel:  "ServerAliveCountMax:",
	CompressionLabel:          "Compression:",
	TCPKeepAliveLabel:         "TCPKeepAlive:",
	ControlMasterLabel:        "ControlMaster:",
	ControlPathLabel:          "ControlPath:",
	ControlPersistLabel:       "ControlPersist:",

	LocalForwardLabel:         "LocalForward:",
	RemoteForwardLabel:        "RemoteForward:",
	DynamicForwardLabel:       "DynamicForward:",
	ClearAllForwardingsLabel:  "ClearAllForwardings:",
	ExitOnForwardFailureLabel: "ExitOnForwardFailure:",
	GatewayPortsLabel:         "GatewayPorts:",
	ForwardAgentLabel:         "ForwardAgent:",
	ForwardX11Label:           "ForwardX11:",
	ForwardX11TrustedLabel:    "ForwardX11Trusted:",

	PubkeyAuthenticationLabel:      "PubkeyAuthentication:",
	IdentitiesOnlyLabel:            "IdentitiesOnly:",
	AddKeysToAgentLabel:            "AddKeysToAgent:",
	IdentityAgentLabel:             "IdentityAgent:",
	PasswordAuthenticationLabel:    "PasswordAuthentication:",
	KbdInteractiveAuthenticationLabel: "KbdInteractiveAuthentication:",
	NumberOfPasswordPromptsLabel:   "NumberOfPasswordPrompts:",
	PreferredAuthenticationsLabel:  "PreferredAuthentications:",
	PubkeyAcceptedAlgorithmsLabel:  "PubkeyAcceptedAlgorithms:",
	HostbasedAcceptedAlgorithmsLabel: "HostbasedAcceptedAlgorithms:",

	StrictHostKeyCheckingLabel: "StrictHostKeyChecking:",
	CheckHostIPLabel:          "CheckHostIP:",
	FingerprintHashLabel:      "FingerprintHash:",
	UserKnownHostsFileLabel:   "UserKnownHostsFile:",
	HostKeyAlgorithmsLabel:    "HostKeyAlgorithms:",
	MACsLabel:                 "MACs:",
	CiphersLabel:              "Ciphers:",
	KexAlgorithmsLabel:        "KexAlgorithms:",
	VerifyHostKeyDNSLabel:     "VerifyHostKeyDNS:",
	UpdateHostKeysLabel:       "UpdateHostKeys:",
	HashKnownHostsLabel:       "HashKnownHosts:",
	VisualHostKeyLabel:        "VisualHostKey:",

	LocalCommandLabel:       "LocalCommand:",
	PermitLocalCommandLabel: "PermitLocalCommand:",
	EscapeCharLabel:         "EscapeChar:",
	SendEnvLabel:            "SendEnv:",
	SetEnvLabel:             "SetEnv:",
	LogLevelLabel:           "LogLevel:",

	SaveButton:   "Save",
	CancelButton: "Cancel",
	StartButton:  "Start",
	DeleteButton: "Delete",
	CloseButton:  "Close",

	SortLabel: " Sort: %s ",
	TitleSort: " Servers — Sort: %s ",

	SortAscSuffix:  " ↑",
	SortDescSuffix: " ↓",

	SectionProxyCommand:          "▶ Proxy & Command",
	SectionConnectionSettings:    "▶ Connection Settings",
	SectionBindOptions:           "▶ Bind Options",
	SectionHostnameCanonicalization: "▶ Hostname Canonicalization",
	SectionKeepAlive:             "▶ Keep-Alive",
	SectionMultiplexing:          "▶ Multiplexing",
	SectionPortForwarding:        "▶ Port Forwarding",
	SectionAgentX11Forwarding:    "▶ Agent & X11 Forwarding",
	SectionPublicKeyAuth:         "▶ Public Key Authentication",
	SectionSSHAgent:              "▶ SSH Agent",
	SectionPasswordInteractive:   "▶ Password & Interactive",
	SectionSecurity:              "▶ Security",
	SectionCommandExecution:      "▶ Command Execution",
	SectionEnvironment:           "▶ Environment",

	TypeLabel:                 "Type",
	PortLabelForward:          "Port",
	HostLabelForward:          "Host",
	HostPortLabel:             "Host Port",
	BindAddressOptionalLabel:  "Bind Address (optional)",
	ModeLabel:                 "Mode",

	ForwardTypeLocal:    "Local",
	ForwardTypeRemote:   "Remote",
	ForwardTypeDynamic:  "Dynamic",
	ForwardModeOnly:     "Only forward",
	ForwardModeForwardSSH: "Forward + SSH",

	DefaultSessionType: "default",
	NoneSessionType:    "none (-N)",
	SubsystemSessionType: "subsystem (-s)",

	DeleteConfirmMsg:  "Delete server %s (%s@%s:%d)?",
	DeleteCannotUndo:  "This action cannot be undone.",

	NoHelpAvailable: "No help available for this field",
	ExamplePrefix:   "e.g., ",
	AvailableSince:  "Available since: ",

	GruConnectionProxy:  "Connection & Proxy",
	GruAuthentication:   "Authentication",
	GruForwarding:       "Forwarding",
	GruSecurityCrypto:   "Security & Cryptography",
	GruEnvironmentExec:  "Environment & Execution",
	GruDebugging:        "Debugging",

	ValidationErrorRequired: " is required",
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
	return GetFieldHelp(field)
}