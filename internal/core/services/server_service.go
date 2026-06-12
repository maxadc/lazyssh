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

package services

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Adembc/lazyssh/internal/core/domain"
	"github.com/Adembc/lazyssh/internal/core/ports"
	"go.uber.org/zap"
)

// SSH config value constants
const (
	sshYes   = "yes"
	sshNo    = "no"
	sshForce = "force"
	sshAuto  = "auto"

	// SessionType values
	sessionTypeNone      = "none"
	sessionTypeSubsystem = "subsystem"
)

type serverService struct {
	serverRepository ports.ServerRepository
	logger           *zap.SugaredLogger

	fwMu     sync.Mutex
	forwards map[string][]*os.Process
}

// NewServerService creates a new instance of serverService.
func NewServerService(logger *zap.SugaredLogger, sr ports.ServerRepository) ports.ServerService {
	return &serverService{
		logger:           logger,
		serverRepository: sr,
	}
}

// ListServers returns a list of servers sorted with pinned on top.
func (s *serverService) ListServers(query string) ([]domain.Server, error) {
	servers, err := s.serverRepository.ListServers(query)
	if err != nil {
		s.logger.Errorw("failed to list servers", "error", err)
		return nil, err
	}

	// Sort: pinned first (PinnedAt non-zero), then by PinnedAt desc, then by Alias asc.
	sort.SliceStable(servers, func(i, j int) bool {
		pi := !servers[i].PinnedAt.IsZero()
		pj := !servers[j].PinnedAt.IsZero()
		if pi != pj {
			return pi
		}
		if pi && pj {
			return servers[i].PinnedAt.After(servers[j].PinnedAt)
		}
		return servers[i].Alias < servers[j].Alias
	})

	return servers, nil
}

// validateServer performs core validation of server fields.
func validateServer(srv domain.Server) error {
	if strings.TrimSpace(srv.Alias) == "" {
		return fmt.Errorf("alias is required")
	}
	if ok, _ := regexp.MatchString(`^[A-Za-z0-9_.-]+$`, srv.Alias); !ok {
		return fmt.Errorf("alias may contain letters, digits, dot, dash, underscore")
	}
	if strings.TrimSpace(srv.Host) == "" {
		return fmt.Errorf("Host/IP is required")
	}
	if ip := net.ParseIP(srv.Host); ip == nil {
		if strings.Contains(srv.Host, " ") {
			return fmt.Errorf("host must not contain spaces")
		}
		if ok, _ := regexp.MatchString(`^[A-Za-z0-9.-]+$`, srv.Host); !ok {
			return fmt.Errorf("host contains invalid characters")
		}
		if strings.HasPrefix(srv.Host, ".") || strings.HasSuffix(srv.Host, ".") {
			return fmt.Errorf("host must not start or end with a dot")
		}
		for _, lbl := range strings.Split(srv.Host, ".") {
			if lbl == "" {
				return fmt.Errorf("host must not contain empty labels")
			}
			if strings.HasPrefix(lbl, "-") || strings.HasSuffix(lbl, "-") {
				return fmt.Errorf("hostname labels must not start or end with a hyphen")
			}
		}
	}
	if srv.Port != 0 && (srv.Port < 1 || srv.Port > 65535) {
		return fmt.Errorf("port must be a number between 1 and 65535")
	}
	return nil
}

// UpdateServer updates an existing server with new details.
func (s *serverService) UpdateServer(server domain.Server, newServer domain.Server) error {
	if err := validateServer(newServer); err != nil {
		s.logger.Warnw("validation failed on update", "error", err, "server", newServer)
		return err
	}
	err := s.serverRepository.UpdateServer(server, newServer)
	if err != nil {
		s.logger.Errorw("failed to update server", "error", err, "server", server)
	}
	return err
}

// AddServer adds a new server to the repository.
func (s *serverService) AddServer(server domain.Server) error {
	if err := validateServer(server); err != nil {
		s.logger.Warnw("validation failed on add", "error", err, "server", server)
		return err
	}
	err := s.serverRepository.AddServer(server)
	if err != nil {
		s.logger.Errorw("failed to add server", "error", err, "server", server)
	}
	return err
}

// DeleteServer removes a server from the repository.
func (s *serverService) DeleteServer(server domain.Server) error {
	err := s.serverRepository.DeleteServer(server)
	if err != nil {
		s.logger.Errorw("failed to delete server", "error", err, "server", server)
	}
	return err
}

// SetPinned sets or clears a pin timestamp for the server alias.
func (s *serverService) SetPinned(alias string, pinned bool) error {
	err := s.serverRepository.SetPinned(alias, pinned)
	if err != nil {
		s.logger.Errorw("failed to set pin state", "error", err, "alias", alias, "pinned", pinned)
	}
	return err
}

// SSH starts an interactive SSH session to the given alias using the system's ssh client.
func (s *serverService) SSH(alias string) error {
	s.logger.Infow("ssh start", "alias", alias)
	cmd := exec.Command("ssh", alias)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	var stderrBuf strings.Builder
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		s.logger.Errorw("ssh command failed", "alias", alias, "error", err)
		errMsg := strings.TrimSpace(stderrBuf.String())
		if errMsg != "" {
			return fmt.Errorf("%w: %s", err, errMsg)
		}
		return err
	}

	if err := s.serverRepository.RecordSSH(alias); err != nil {
		s.logger.Errorw("failed to record ssh metadata", "alias", alias, "error", err)
	}

	s.logger.Infow("ssh end", "alias", alias)
	return nil
}

// SSHWithArgs runs system ssh with provided extra args (e.g., -L/-R/-D) for the given alias.
func (s *serverService) SSHWithArgs(alias string, extraArgs []string) error {
	s.logger.Infow("ssh start (with args)", "alias", alias, "args", extraArgs)
	args := append([]string{}, extraArgs...)
	args = append(args, alias)
	// #nosec G204
	cmd := exec.Command("ssh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	var stderrBuf strings.Builder
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		s.logger.Errorw("ssh (with args) failed", "alias", alias, "error", err)
		errMsg := strings.TrimSpace(stderrBuf.String())
		if errMsg != "" {
			return fmt.Errorf("%w: %s", err, errMsg)
		}
		return err
	}
	if err := s.serverRepository.RecordSSH(alias); err != nil {
		s.logger.Errorw("failed to record ssh metadata", "alias", alias, "error", err)
	}
	s.logger.Infow("ssh end (with args)", "alias", alias)
	return nil
}

// StartForward starts ssh port forwarding in the background and tracks the process.
func (s *serverService) StartForward(alias string, extraArgs []string) (int, error) {
	s.fwMu.Lock()
	if s.forwards == nil {
		s.forwards = make(map[string][]*os.Process)
	}
	s.fwMu.Unlock()

	extraArgs = append(extraArgs, "-N", alias)

	// #nosec G204
	cmd := exec.Command("ssh", extraArgs...)

	// Detach from TTY: discard stdio
	devNull, err := os.OpenFile(os.DevNull, os.O_RDWR, 0)
	if err != nil {
		return 0, fmt.Errorf("failed to open devnull: %w", err)
	}
	defer func() {
		if devNull != nil {
			_ = devNull.Close()
		}
	}()

	cmd.Stdin = devNull
	cmd.Stdout = devNull
	cmd.Stderr = devNull
	// Set SysProcAttr in an OS-specific way (see sysprocattr_* files)
	sysProcAttr := &syscall.SysProcAttr{}
	setDetach(sysProcAttr)
	cmd.SysProcAttr = sysProcAttr

	if err := cmd.Start(); err != nil {
		return 0, fmt.Errorf("failed to start ssh: %w", err)
	}

	proc := cmd.Process
	if proc == nil {
		return 0, fmt.Errorf("process is nil after start")
	}
	pid := proc.Pid

	// Track process
	s.fwMu.Lock()
	s.forwards[alias] = append(s.forwards[alias], proc)
	s.fwMu.Unlock()

	// Cleanup on exit
	go func(a string, c *exec.Cmd, dn *os.File) {
		_ = c.Wait()
		_ = dn.Close()

		s.fwMu.Lock()
		defer s.fwMu.Unlock()

		procs := s.forwards[a]
		if len(procs) == 0 {
			return
		}

		filtered := make([]*os.Process, 0, len(procs))
		for _, p := range procs {
			if p != nil && p.Pid != pid {
				filtered = append(filtered, p)
			}
		}

		if len(filtered) == 0 {
			delete(s.forwards, a)
		} else {
			s.forwards[a] = filtered
		}
	}(alias, cmd, devNull)

	devNull = nil // Prevent defer from closing it

	return pid, nil
}

// StopForwarding kills all active forward processes for the alias.
func (s *serverService) StopForwarding(alias string) error {
	s.fwMu.Lock()
	procs := s.forwards[alias]
	delete(s.forwards, alias)
	s.fwMu.Unlock()

	if len(procs) == 0 {
		return nil
	}

	var errs []error
	for _, p := range procs {
		if p != nil {
			if err := p.Signal(syscall.SIGTERM); err != nil {
				// If SIGTERM fails, try SIGKILL
				if killErr := p.Kill(); killErr != nil {
					errs = append(errs, fmt.Errorf("failed to kill pid %d: %w", p.Pid, killErr))
				}
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors stopping forwards: %v", errs)
	}
	return nil
}

// IsForwarding reports whether there is at least one active forward for alias.
func (s *serverService) IsForwarding(alias string) bool {
	s.fwMu.Lock()
	defer s.fwMu.Unlock()
	return len(s.forwards[alias]) > 0
}

// Ping checks if the server is reachable on its SSH port.
func (s *serverService) Ping(server domain.Server) (bool, time.Duration, error) {
	start := time.Now()

	host, port, ok := resolveSSHDestination(server.Alias)
	if !ok {

		host = strings.TrimSpace(server.Host)
		if host == "" {
			host = server.Alias
		}
		if server.Port > 0 {
			port = server.Port
		} else {
			port = 22
		}
	}
	addr := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	dialer := net.Dialer{Timeout: 3 * time.Second}
	conn, err := dialer.Dial("tcp", addr)
	if err != nil {
		return false, time.Since(start), err
	}
	_ = conn.Close()
	return true, time.Since(start), nil
}

// resolveSSHDestination uses `ssh -G <alias>` to extract HostName and Port from the user's SSH config.
// Returns host, port, ok where ok=false if resolution failed.
func resolveSSHDestination(alias string) (string, int, bool) {
	alias = strings.TrimSpace(alias)
	if alias == "" {
		return "", 0, false
	}
	cmd := exec.Command("ssh", "-G", alias)
	out, err := cmd.Output()
	if err != nil {
		return "", 0, false
	}
	host := ""
	port := 0
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "hostname ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				host = parts[1]
			}
		}
		if strings.HasPrefix(line, "port ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				if p, err := strconv.Atoi(parts[1]); err == nil {
					port = p
				}
			}
		}
	}
	if host == "" {
		host = alias
	}
	if port == 0 {
		port = 22
	}
	return host, port, true
}

// SSHWithPassword connects to the server using sshpass for password-based authentication.
// It captures the stderr output to provide detailed error messages on failure.
func (s *serverService) SSHWithPassword(server domain.Server) error {
	s.logger.Infow("ssh with password start", "alias", server.Alias)

	// Check if sshpass is available
	sshpassPath, err := exec.LookPath("sshpass")
	if err != nil {
		s.logger.Errorw("sshpass not found", "error", err)
		return fmt.Errorf("sshpass not found in PATH")
	}

	// Build the sshpass command
	sshCmd := s.buildSSHPassCommand(server, sshpassPath)

	// Execute the command, capture stderr for error reporting
	cmd := exec.Command("sh", "-c", sshCmd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	var stderrBuf strings.Builder
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		s.logger.Errorw("ssh with password failed", "alias", server.Alias, "error", err)
		// Include stderr output in the error message
		errMsg := stderrBuf.String()
		if errMsg != "" {
			// Strip trailing newlines and trim
			errMsg = strings.TrimSpace(errMsg)
			return fmt.Errorf("%w: %s", err, errMsg)
		}
		return err
	}

	if err := s.serverRepository.RecordSSH(server.Alias); err != nil {
		s.logger.Errorw("failed to record ssh metadata", "alias", server.Alias, "error", err)
	}

	s.logger.Infow("ssh with password end", "alias", server.Alias)
	return nil
}

// buildSSHPassCommand constructs the full sshpass command string for password auth.
func (s *serverService) buildSSHPassCommand(server domain.Server, sshpassPath string) string {
	// Build base ssh command parts
	parts := []string{fmt.Sprintf("%q", sshpassPath), "-p", quoteShellArg(server.Password), "ssh"}

	// Add proxy and connection options
	s.addProxySSHOptions(&parts, server)
	s.addConnectionTimingSSHOptions(&parts, server)

	// Add port forwarding options
	s.addPortForwardingSSHOptions(&parts, server)

	// Add authentication options
	s.addAuthSSHOptions(&parts, server)

	// Add agent and forwarding options
	s.addForwardingSSHOptions(&parts, server)

	// Add connection multiplexing options
	s.addMultiplexingSSHOptions(&parts, server)

	// Add connection reliability options
	s.addConnectionSSHOptions(&parts, server)

	// Add security options
	s.addSecuritySSHOptions(&parts, server)

	// Add command execution options
	s.addCommandExecutionSSHOptions(&parts, server)

	// Add environment options
	s.addEnvironmentSSHOptions(&parts, server)

	// Add TTY and logging options
	s.addTTYAndLoggingSSHOptions(&parts, server)

	// Port option
	if server.Port != 0 && server.Port != 22 {
		parts = append(parts, "-p", fmt.Sprintf("%d", server.Port))
	}

	// Identity file option
	if len(server.IdentityFiles) > 0 {
		for _, keyFile := range server.IdentityFiles {
			parts = append(parts, "-i", quoteShellArg(keyFile))
		}
	}

	// Host specification
	userHost := ""
	switch {
	case server.User != "" && server.Host != "":
		userHost = fmt.Sprintf("%s@%s", server.User, server.Host)
	case server.Host != "":
		userHost = server.Host
	default:
		userHost = server.Alias
	}
	parts = append(parts, userHost)

	// RemoteCommand (must come after the host)
	if server.RemoteCommand != "" {
		if server.RemoteCommand == sessionTypeNone {
			parts = append(parts, "-o", "RemoteCommand=none")
		} else {
			parts = append(parts, quoteShellArg(server.RemoteCommand))
		}
	}

	return strings.Join(parts, " ")
}

// Helper methods to add SSH options for sshpass command building

func (s *serverService) addProxySSHOptions(parts *[]string, server domain.Server) {
	if server.ProxyJump != "" {
		*parts = append(*parts, "-J", quoteShellArg(server.ProxyJump))
	}
	if server.ProxyCommand != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("ProxyCommand=%s", quoteShellArg(server.ProxyCommand)))
	}
}

func (s *serverService) addConnectionTimingSSHOptions(parts *[]string, server domain.Server) {
	if server.ConnectTimeout != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("ConnectTimeout=%s", server.ConnectTimeout))
	}
	if server.ConnectionAttempts != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("ConnectionAttempts=%s", server.ConnectionAttempts))
	}
	if server.BindAddress != "" {
		*parts = append(*parts, "-b", server.BindAddress)
	}
	if server.BindInterface != "" {
		*parts = append(*parts, "-B", server.BindInterface)
	}
	if server.AddressFamily != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("AddressFamily=%s", server.AddressFamily))
	}
	if server.IPQoS != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("IPQoS=%s", server.IPQoS))
	}
	if server.CanonicalizeHostname != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("CanonicalizeHostname=%s", server.CanonicalizeHostname))
	}
	if server.CanonicalDomains != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("CanonicalDomains=%s", server.CanonicalDomains))
	}
	if server.CanonicalizeFallbackLocal != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("CanonicalizeFallbackLocal=%s", server.CanonicalizeFallbackLocal))
	}
	if server.CanonicalizeMaxDots != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("CanonicalizeMaxDots=%s", server.CanonicalizeMaxDots))
	}
	if server.CanonicalizePermittedCNAMEs != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("CanonicalizePermittedCNAMEs=%s", quoteShellArg(server.CanonicalizePermittedCNAMEs)))
	}
}

func (s *serverService) addPortForwardingSSHOptions(parts *[]string, server domain.Server) {
	for _, forward := range server.LocalForward {
		*parts = append(*parts, "-L", forward)
	}
	for _, forward := range server.RemoteForward {
		*parts = append(*parts, "-R", forward)
	}
	for _, forward := range server.DynamicForward {
		*parts = append(*parts, "-D", forward)
	}
	if server.ClearAllForwardings == sshYes {
		*parts = append(*parts, "-o", "ClearAllForwardings=yes")
	}
	if server.ExitOnForwardFailure == sshYes {
		*parts = append(*parts, "-o", "ExitOnForwardFailure=yes")
	}
	if server.GatewayPorts != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("GatewayPorts=%s", server.GatewayPorts))
	}
}

func (s *serverService) addAuthSSHOptions(parts *[]string, server domain.Server) {
	if server.PubkeyAuthentication != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("PubkeyAuthentication=%s", server.PubkeyAuthentication))
	}
	if server.PubkeyAcceptedAlgorithms != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("PubkeyAcceptedAlgorithms=%s", server.PubkeyAcceptedAlgorithms))
	}
	if server.HostbasedAcceptedAlgorithms != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("HostbasedAcceptedAlgorithms=%s", server.HostbasedAcceptedAlgorithms))
	}
	if server.PasswordAuthentication != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("PasswordAuthentication=%s", server.PasswordAuthentication))
	}
	if server.PreferredAuthentications != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("PreferredAuthentications=%s", server.PreferredAuthentications))
	}
	if server.IdentitiesOnly != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("IdentitiesOnly=%s", server.IdentitiesOnly))
	}
	if server.AddKeysToAgent != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("AddKeysToAgent=%s", server.AddKeysToAgent))
	}
	if server.IdentityAgent != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("IdentityAgent=%s", quoteShellArg(server.IdentityAgent)))
	}
	if server.KbdInteractiveAuthentication != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("KbdInteractiveAuthentication=%s", server.KbdInteractiveAuthentication))
	}
	if server.NumberOfPasswordPrompts != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("NumberOfPasswordPrompts=%s", server.NumberOfPasswordPrompts))
	}
}

func (s *serverService) addForwardingSSHOptions(parts *[]string, server domain.Server) {
	if server.ForwardAgent != "" {
		if server.ForwardAgent == sshYes {
			*parts = append(*parts, "-A")
		} else if server.ForwardAgent == sshNo {
			*parts = append(*parts, "-a")
		}
	}
	if server.ForwardX11 != "" {
		if server.ForwardX11 == sshYes {
			*parts = append(*parts, "-X")
		} else if server.ForwardX11 == sshNo {
			*parts = append(*parts, "-x")
		}
	}
	if server.ForwardX11Trusted == sshYes {
		*parts = append(*parts, "-Y")
	}
}

func (s *serverService) addMultiplexingSSHOptions(parts *[]string, server domain.Server) {
	if server.ControlMaster != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("ControlMaster=%s", server.ControlMaster))
	}
	if server.ControlPath != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("ControlPath=%s", quoteShellArg(server.ControlPath)))
	}
	if server.ControlPersist != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("ControlPersist=%s", server.ControlPersist))
	}
}

func (s *serverService) addConnectionSSHOptions(parts *[]string, server domain.Server) {
	if server.ServerAliveInterval != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("ServerAliveInterval=%s", server.ServerAliveInterval))
	}
	if server.ServerAliveCountMax != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("ServerAliveCountMax=%s", server.ServerAliveCountMax))
	}
	if server.Compression == sshYes {
		*parts = append(*parts, "-C")
	}
	if server.TCPKeepAlive != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("TCPKeepAlive=%s", server.TCPKeepAlive))
	}
	if server.BatchMode == sshYes {
		*parts = append(*parts, "-o", "BatchMode=yes")
	}
}

func (s *serverService) addSecuritySSHOptions(parts *[]string, server domain.Server) {
	if server.StrictHostKeyChecking != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("StrictHostKeyChecking=%s", server.StrictHostKeyChecking))
	}
	if server.CheckHostIP != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("CheckHostIP=%s", server.CheckHostIP))
	}
	if server.FingerprintHash != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("FingerprintHash=%s", server.FingerprintHash))
	}
	if server.UserKnownHostsFile != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("UserKnownHostsFile=%s", quoteShellArg(server.UserKnownHostsFile)))
	}
	if server.HostKeyAlgorithms != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("HostKeyAlgorithms=%s", server.HostKeyAlgorithms))
	}
	if server.MACs != "" {
		*parts = append(*parts, "-m", server.MACs)
	}
	if server.Ciphers != "" {
		*parts = append(*parts, "-c", server.Ciphers)
	}
	if server.KexAlgorithms != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("KexAlgorithms=%s", server.KexAlgorithms))
	}
	if server.VerifyHostKeyDNS != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("VerifyHostKeyDNS=%s", server.VerifyHostKeyDNS))
	}
	if server.UpdateHostKeys != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("UpdateHostKeys=%s", server.UpdateHostKeys))
	}
	if server.HashKnownHosts != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("HashKnownHosts=%s", server.HashKnownHosts))
	}
	if server.VisualHostKey != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("VisualHostKey=%s", server.VisualHostKey))
	}
}

func (s *serverService) addCommandExecutionSSHOptions(parts *[]string, server domain.Server) {
	if server.LocalCommand != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("LocalCommand=%s", quoteShellArg(server.LocalCommand)))
	}
	if server.PermitLocalCommand != "" {
		*parts = append(*parts, "-o", fmt.Sprintf("PermitLocalCommand=%s", server.PermitLocalCommand))
	}
	if server.EscapeChar != "" {
		*parts = append(*parts, "-e", server.EscapeChar)
	}
}

func (s *serverService) addEnvironmentSSHOptions(parts *[]string, server domain.Server) {
	for _, env := range server.SendEnv {
		*parts = append(*parts, "-o", fmt.Sprintf("SendEnv=%s", env))
	}
	for _, env := range server.SetEnv {
		*parts = append(*parts, "-o", fmt.Sprintf("SetEnv=%s", quoteShellArg(env)))
	}
}

func (s *serverService) addTTYAndLoggingSSHOptions(parts *[]string, server domain.Server) {
	if server.RequestTTY != "" {
		switch server.RequestTTY {
		case sshYes:
			*parts = append(*parts, "-t")
		case sshNo:
			*parts = append(*parts, "-T")
		case sshForce:
			*parts = append(*parts, "-tt")
		case sshAuto:
			// auto is the default behavior, no flag needed
		default:
			*parts = append(*parts, "-o", fmt.Sprintf("RequestTTY=%s", server.RequestTTY))
		}
	}

	if server.LogLevel != "" {
		switch strings.ToLower(server.LogLevel) {
		case "quiet":
			*parts = append(*parts, "-q")
		case "verbose":
			*parts = append(*parts, "-v")
		case "debug", "debug1":
			*parts = append(*parts, "-v")
		case "debug2":
			*parts = append(*parts, "-vv")
		case "debug3":
			*parts = append(*parts, "-vvv")
		}
	}

	if server.SessionType != "" {
		switch server.SessionType {
		case sessionTypeNone:
			*parts = append(*parts, "-N")
		case sessionTypeSubsystem:
			*parts = append(*parts, "-s")
		default:
			*parts = append(*parts, "-o", fmt.Sprintf("SessionType=%s", server.SessionType))
		}
	}
}

// quoteShellArg returns the value quoted if it contains shell-special characters.
func quoteShellArg(val string) string {
	if strings.ContainsAny(val, " \t\n\"'\\$!") {
		return fmt.Sprintf("'%s'", strings.ReplaceAll(val, "'", "'\\''"))
	}
	return val
}
