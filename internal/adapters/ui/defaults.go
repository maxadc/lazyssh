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

// SSHFieldDefaults contains the default values for all SSH configuration fields
// This centralizes all default values to ensure consistency across the application
var SSHFieldDefaults = map[string]string{
	// Basic fields
	"Port": "22",
	"User": "", // Empty means current username (OpenSSH default)

	// Connection fields
	"ConnectTimeout":     "", // none (system default)
	"ConnectionAttempts": "1",
	"IPQoS":              "af21 cs1",
	"BatchMode":          "no",
	"Compression":        "no",
	"AddressFamily":      "any",
	"RequestTTY":         "auto",
	"SessionType":        "default",

	// Proxy fields
	"ProxyJump":     "", // none
	"ProxyCommand":  "", // none
	"RemoteCommand": "", // none

	// Port forwarding fields
	"LocalForward":         "", // none
	"RemoteForward":        "", // none
	"DynamicForward":       "", // none
	"ForwardAgent":         "no",
	"ForwardX11":           "no",
	"ForwardX11Trusted":    "no",
	"ClearAllForwardings":  "no",
	"ExitOnForwardFailure": "no",
	"GatewayPorts":         "no",

	// Authentication fields
	"PubkeyAuthentication":         "yes",
	"PasswordAuthentication":       "yes",
	"PreferredAuthentications":     "gssapi-with-mic,hostbased,publickey,keyboard-interactive,password",
	"IdentitiesOnly":               "no",
	"AddKeysToAgent":               "no",
	"IdentityAgent":                "SSH_AUTH_SOCK",
	"KbdInteractiveAuthentication": "yes",
	"NumberOfPasswordPrompts":      "3",
	"PubkeyAcceptedAlgorithms":     "", // all supported
	"HostbasedAcceptedAlgorithms":  "", // all supported

	// Multiplexing fields
	"ControlMaster":  "no",
	"ControlPath":    "", // none
	"ControlPersist": "no",

	// Keep-alive fields
	"ServerAliveInterval": "0", // disabled
	"ServerAliveCountMax": "3",
	"TCPKeepAlive":        "yes",

	// Security fields
	"StrictHostKeyChecking": "ask",
	"UserKnownHostsFile":    "~/.ssh/known_hosts",
	"HostKeyAlgorithms":     "", // default algorithms
	"Ciphers":               "", // default ciphers
	"MACs":                  "", // default MACs
	"CheckHostIP":           "no",
	"FingerprintHash":       "SHA256", // OpenSSH uses uppercase SHA256
	"VerifyHostKeyDNS":      "no",
	"UpdateHostKeys":        "no",
	"HashKnownHosts":        "no",
	"VisualHostKey":         "no",

	// Cryptography fields
	"KexAlgorithms": "", // all supported

	// Hostname canonicalization fields
	"CanonicalizeHostname":        "no",
	"CanonicalDomains":            "", // none
	"CanonicalizeFallbackLocal":   "yes",
	"CanonicalizeMaxDots":         "1",
	"CanonicalizePermittedCNAMEs": "", // none

	// Command execution fields
	"LocalCommand":       "", // none
	"PermitLocalCommand": "no",
	"EscapeChar":         "~",

	// Environment fields
	"SendEnv": "", // none
	"SetEnv":  "", // none

	// Debugging fields
	"LogLevel": "INFO",

	// Bind options
	"BindAddress":   "", // none
	"BindInterface": "", // none
}

// GetSSHFieldDefault returns the default value for a given SSH field
// Returns empty string if no default is defined
func GetSSHFieldDefault(fieldName string) string {
	if value, exists := SSHFieldDefaults[fieldName]; exists {
		return value
	}
	return ""
}

// GetSSHFieldDefaultWithFallback returns the default value for a given SSH field
// with a fallback value if no default is defined
func GetSSHFieldDefaultWithFallback(fieldName, fallback string) string {
	if value, exists := SSHFieldDefaults[fieldName]; exists {
		return value
	}
	return fallback
}
