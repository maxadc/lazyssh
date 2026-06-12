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

// GetFieldPlaceholder returns an appropriate placeholder for a form field.
// English version. For zh build, see defaults_zh.go.
func GetFieldPlaceholder(fieldName string) string {
	switch fieldName {
	case "Alias", "Host":
		return "required"
	case "Port":
		return "default: 22"
	case "User":
		return "default: current username"
	case "Keys":
		return "e.g., ~/.ssh/id_rsa, ~/.ssh/id_ed25519"
	case "Tags":
		return "comma-separated tags"
	case "ProxyJump":
		return "e.g., bastion.example.com"
	case "ProxyCommand":
		return "e.g., ssh -W %h:%p jump.example.com"
	case "RemoteCommand":
		return "e.g., tmux attach"
	case "LocalForward":
		return "e.g., 8080:localhost:80, 3000:localhost:3000"
	case "RemoteForward":
		return "e.g., 80:localhost:8080"
	case "DynamicForward":
		return "e.g., 1080, 1081"
	case "ControlPath":
		return "e.g., ~/.ssh/master-%r@%h:%p"
	case "ControlPersist":
		return "e.g., 10m, 4h, yes, no"
	case "PreferredAuthentications":
		return "e.g., publickey,password"
	case "PubkeyAcceptedAlgorithms", "HostbasedAcceptedAlgorithms",
		"HostKeyAlgorithms", "Ciphers", "MACs", "KexAlgorithms":
		return "algorithms (+/-/^ prefix supported)"
	case "BindAddress":
		return "IP, hostname, * (all), or localhost"
	case "CanonicalDomains":
		return "e.g., example.com, internal.net"
	case "CanonicalizePermittedCNAMEs":
		return "e.g., *.example.com:example.net"
	case "LocalCommand":
		return "e.g., echo 'Connected to %h'"
	case "SendEnv":
		return "e.g., LANG, LC_*, TERM"
	case "SetEnv":
		return "e.g., FOO=bar, DEBUG=1"
	default:
		return ""
	}
}

// GetFormSectionHeader returns the section header for a form tab.
// Default returns the key with "▶ " prefix. For zh build, see defaults_zh.go
func GetFormSectionHeader(key string) string {
	return "▶ " + key
}

// GetFormLabel returns the form field label. Default is the key with ":" suffix.
// For zh build, see defaults_zh.go
func GetFormLabel(key string) string {
	return key
}

// GetTabAbbrev returns the abbreviated tab name for narrow views. Default is the key itself.
// For zh build, see defaults_zh.go
func GetTabAbbrev(fullName string) string {
	return fullName
}

// GetDetailGroupName returns the detail panel group name. Default is the key itself.
// For zh build, see defaults_zh.go
func GetDetailGroupName(key string) string {
	return key
}

// GetNoHelpText returns text for "no help available". Default is English.
// For zh build, see defaults_zh.go
func GetNoHelpText() string {
	return "No help available for this field"
}

// GetExamplePrefix returns text for "e.g.". Default is English.
// For zh build, see defaults_zh.go
func GetExamplePrefix() string {
	return "e.g., "
}

// GetAvailableSince returns text for "Available since:". Default is English.
// For zh build, see defaults_zh.go
func GetAvailableSince() string {
	return "Available since: "
}

// GetAuthMethodDefault returns the default auth method.
// Default is empty string (auto). For zh build, see defaults_zh.go.
func GetAuthMethodDefault() string {
	return ""
}

// GetPasswordDefault returns the default password value.
// Default is empty string. For zh build, see defaults_zh.go.
func GetPasswordDefault() string {
	return ""
}
