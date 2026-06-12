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

package ui

// I18n contains all translatable strings for the UI.
// Values are populated by i18n.go (English) or i18n_zh.go (Chinese).
type I18n struct {
	AliasAsc            string
	AliasDesc           string
	LastSeenAsc         string
	LastSeenDesc        string
	SearchLabel         string
	SearchTitle         string
	ServersTitle        string
	DetailsTitle        string
	HelpTitle           string
	AddServerTitle      string
	EditServerTitle     string
	DeleteConfirmTitle  string
	EditTagsTitle       string
	PortForwardTitle    string
	NeverSeen           string
	NoServersMatch      string
	SaveFailed          string
	Pinging             string
	PingDown            string
	PingUp              string
	Refreshing          string
	RefreshFailed       string
	Refreshed           string
	Copied              string
	CopyFailed          string
	TagsUpdated         string
	TagsLabel           string
	PortForwardStarted  string
	ForwardFailed       string
	StoppedForwarding   string
	StopForwardFailed   string
	InvalidPort         string
	InvalidBindAddress  string
	InvalidHost         string
	InvalidHostPort     string
	StartingPortForward string

	BasicSettings    string
	AdvancedSettings string
	Commands         string
	HostLabel        string
	UserLabel        string
	PortLabel        string
	KeyLabel         string
	TagsLabelDetail  string
	PinnedLabel      string
	LastSSHLabel     string
	SSHCountLabel    string
	AliasLabel       string

	ConnectionProxy    string
	Authentication     string
	Forwarding         string
	SecurityCrypto     string
	EnvironmentExec    string
	Debugging          string

	FormTabBasic          string
	FormTabConnection     string
	FormTabForwarding     string
	FormTabAuthentication string
	FormTabAdvanced       string

	FormHintNavigate string
	FormHintSave     string
	FormHintCancel   string

	StatusNavigate     string
	StatusSSH          string
	StatusForward      string
	StatusStopForward  string
	StatusCopySSH      string
	StatusAdd          string
	StatusEdit         string
	StatusPing         string
	StatusDelete       string
	StatusPin          string
	StatusSearch       string
	StatusQuit         string

	CommandsTitle     string
	CommandsSSH       string
	CommandsForward   string
	CommandsStopForward string
	CommandsCopySSH   string
	CommandsPing      string
	CommandsRefresh   string
	CommandsAdd       string
	CommandsEdit      string
	CommandsTags      string
	CommandsDelete    string
	CommandsPin       string
	CommandsSearch    string

	StatusLabel           string
	HostIPLabel           string
	UserLabelForm         string
	PortLabelForm         string
	KeysLabelForm         string
	TagsLabelForm         string
	AuthMethodLabel       string
	PasswordLabel         string
	AuthMethodAuto        string
	AuthMethodKey         string
	AuthMethodPassword    string
	SSHPassNotFound       string
	PasswordAuthFailed    string
	SSHConnectionFailed   string
	ProxyJumpLabel        string
	ProxyCommandLabel     string
	RemoteCommandLabel    string
	RequestTTYLabel       string
	SessionTypeLabel      string
	ConnectTimeoutLabel   string
	ConnectionAttemptsLabel string
	IPQoSLabel            string
	BatchModeLabel        string
	BindAddressLabel      string
	BindInterfaceLabel    string
	AddressFamilyLabel    string
	CanonicalizeHostnameLabel string
	CanonicalDomainsLabel     string
	CanonicalizeFallbackLocalLabel string
	CanonicalizeMaxDotsLabel  string
	CanonicalizePermittedCNAMEsLabel string
	ServerAliveIntervalLabel  string
	ServerAliveCountMaxLabel  string
	CompressionLabel          string
	TCPKeepAliveLabel         string
	ControlMasterLabel        string
	ControlPathLabel          string
	ControlPersistLabel       string
	LocalForwardLabel         string
	RemoteForwardLabel        string
	DynamicForwardLabel       string
	ClearAllForwardingsLabel  string
	ExitOnForwardFailureLabel string
	GatewayPortsLabel         string
	ForwardAgentLabel         string
	ForwardX11Label           string
	ForwardX11TrustedLabel    string
	PubkeyAuthenticationLabel string
	IdentitiesOnlyLabel       string
	AddKeysToAgentLabel       string
	IdentityAgentLabel        string
	PasswordAuthenticationLabel string
	KbdInteractiveAuthenticationLabel string
	NumberOfPasswordPromptsLabel  string
	PreferredAuthenticationsLabel string
	PubkeyAcceptedAlgorithmsLabel string
	HostbasedAcceptedAlgorithmsLabel string
	StrictHostKeyCheckingLabel string
	CheckHostIPLabel           string
	FingerprintHashLabel       string
	UserKnownHostsFileLabel    string
	HostKeyAlgorithmsLabel     string
	MACsLabel                  string
	CiphersLabel               string
	KexAlgorithmsLabel         string
	VerifyHostKeyDNSLabel      string
	UpdateHostKeysLabel        string
	HashKnownHostsLabel        string
	VisualHostKeyLabel         string
	LocalCommandLabel          string
	PermitLocalCommandLabel    string
	EscapeCharLabel            string
	SendEnvLabel               string
	SetEnvLabel                string
	LogLevelLabel              string
	SaveButton   string
	CancelButton string
	StartButton  string
	DeleteButton string
	CloseButton  string
	SortLabel    string
	TitleSort    string
	SortAscSuffix  string
	SortDescSuffix string
	SectionProxyCommand        string
	SectionConnectionSettings  string
	SectionBindOptions         string
	SectionHostnameCanonicalization string
	SectionKeepAlive           string
	SectionMultiplexing        string
	SectionPortForwarding      string
	SectionAgentX11Forwarding  string
	SectionPublicKeyAuth       string
	SectionSSHAgent            string
	SectionPasswordInteractive string
	SectionSecurity            string
	SectionCommandExecution    string
	SectionEnvironment         string
	TypeLabel                  string
	PortLabelForward           string
	HostLabelForward           string
	HostPortLabel              string
	BindAddressOptionalLabel   string
	ModeLabel                  string
	ForwardTypeLocal           string
	ForwardTypeRemote          string
	ForwardTypeDynamic         string
	ForwardModeOnly            string
	ForwardModeForwardSSH      string
	YesOption      string
	NoOption       string
	DefaultOption  string
	NoneOption     string
	AskOption      string
	AutoOption     string
	ForceOption    string
	AnyOption      string
	InetOption     string
	Inet6Option    string
	AlwaysOption   string
	DefaultSessionType     string
	NoneSessionType        string
	SubsystemSessionType   string
	DeleteConfirmMsg       string
	DeleteCannotUndo       string
	ValidationErrorRequired string
	NoHelpAvailable        string
	ExamplePrefix          string
	AvailableSince         string
	GruConnectionProxy     string
	GruAuthentication      string
	GruForwarding          string
	GruSecurityCrypto      string
	GruEnvironmentExec     string
	GruDebugging           string
}
