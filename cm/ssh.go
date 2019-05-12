package cm

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/appleboy/easyssh-proxy"
)

// RemoteResponse represents an ssh command output
type RemoteResponse struct {
	StdOut string
	StdErr string
	Done   bool
}

// RunGatewayCMCommand run command on CM CB gateway
func (c CMServer) RunGatewayCMCommand(command string, printStdOut bool, printEmptyResponse bool) RemoteResponse {
	connectionProfileID := c.ConnectionProfile
	if len(connectionProfileID) == 0 {
		fmt.Println("No connection profile is attached for the active CM server entry!")
		os.Exit(1)
	}
	connectionProfile := GetConnectionProfileByID(connectionProfileID)
	var wg sync.WaitGroup
	wg.Add(1)

	responses := make(map[string]RemoteResponse)
	ssh := createSSHConfig(connectionProfile, c.Hostname, "")
	go func(ssh *easyssh.MakeConfig, command string, responses map[string]RemoteResponse) {
		defer wg.Done()
		stdout, stderr, done, err := ssh.Run(command)
		// Handle errors
		if err != nil {
			panic("Can't run remote command: " + err.Error())
		} else {
			if len(stdout) == 0 && printEmptyResponse {
				fmt.Println("Stdout response is empty for remote command.")
			} else {
				if printStdOut && len(strings.TrimSpace(stdout)) > 0 {
					fmt.Println(stdout)
				}
			}
			if len(stderr) > 0 {
				fmt.Println("Error during gateway ssh command:")
				fmt.Println(stderr)
				os.Exit(1)
			}
			responses["server"] = RemoteResponse{StdOut: stdout, StdErr: stderr, Done: done}
		}
	}(ssh, command, responses)
	wg.Wait()
	response := RemoteResponse{}
	if respVal, ok := responses["server"]; ok {
		response = respVal
	}
	return response
}

// CopyToCMGateway copy local file to gateway host
func (c CMServer) CopyToCMGateway(source string, dest string) {
	connectionProfileID := c.ConnectionProfile
	if len(connectionProfileID) == 0 {
		fmt.Println("No connection profile is attached for the active ambari server entry!")
		os.Exit(1)
	}
	connectionProfile := GetConnectionProfileByID(connectionProfileID)

	var wg sync.WaitGroup
	wg.Add(1)
	ssh := createSSHConfig(connectionProfile, c.Hostname, "")
	go func(ssh *easyssh.MakeConfig, source string, dest string, host string) {
		defer wg.Done()
		err := ssh.Scp(source, dest)
		// Handle errors
		if err != nil {
			errMsg := fmt.Sprintf("Can't run remote command on host '%v (scp %v to %v)", host, source, dest)
			fmt.Println(errMsg)
		} else {
			succMsg := fmt.Sprintf("Copying to remote host '%v' is successful. (from - %v, to %v)", host, source, dest)
			fmt.Println(succMsg)
		}
	}(ssh, source, dest, c.Hostname)
	wg.Wait()
}

// RunRemoteHostCommand executes bash commands on CM hosts
func (c CMServer) RunRemoteHostCommand(command string, filteredHosts map[string]bool, skipJump bool, inventory *Inventory) map[string]RemoteResponse {
	connectionProfileID := c.ConnectionProfile
	if len(connectionProfileID) == 0 {
		fmt.Println("No connection profile is attached for the active CM server entry!")
		os.Exit(1)
	}
	connectionProfile := GetConnectionProfileByID(connectionProfileID)
	var hosts map[string]bool
	if len(filteredHosts) > 0 {
		hosts = filteredHosts
	} else {
		hosts = c.GetFilteredHosts(Filter{}, inventory)
	}
	response := make(map[string]RemoteResponse)
	var wg sync.WaitGroup
	wg.Add(len(hosts))
	gatewayHost := ""
	if c.UseGateway {
		gatewayHost = c.Hostname
	}
	for host := range hosts {
		ssh := createSSHConfig(connectionProfile, host, gatewayHost)
		go func(ssh *easyssh.MakeConfig, command string, host string, response map[string]RemoteResponse) {
			defer wg.Done()
			stdout, stderr, done, err := ssh.Run(command)
			// Handle errors
			msgHeader := fmt.Sprintf("%v (done: %v) - output:", host, done)
			fmt.Println(msgHeader)
			if err != nil {
				panic("Can't run remote command: " + err.Error())
			} else {
				if len(stdout) > 0 {
					fmt.Println(stdout)
				}
				if len(stderr) > 0 {
					fmt.Println("std error:")
					fmt.Println(stderr)
				}
				response[host] = RemoteResponse{StdOut: stdout, StdErr: stderr, Done: done}
			}
		}(ssh, command, host, response)
	}
	wg.Wait()
	return response
}

// DownloadViaScp downloads file from remote to local
func DownloadViaScp(sshConfig *easyssh.MakeConfig, source string, dest string, skipJump bool) error {
	userAndRemote := fmt.Sprintf("%v@%v", sshConfig.User, sshConfig.Server)
	var args []string
	if len(sshConfig.Proxy.Server) > 0 && !skipJump {
		args = []string{"-o", fmt.Sprintf("ProxyJump=%v", sshConfig.Proxy.Server), "-o", "StrictHostKeyChecking=no", "-q", "-P", sshConfig.Port, "-i", sshConfig.KeyPath, userAndRemote + ":" + source, dest}
	} else {
		args = []string{"-o", "StrictHostKeyChecking=no", "-q", "-P", sshConfig.Port, "-i", sshConfig.KeyPath, userAndRemote + ":" + source, dest}
	}
	cmd := exec.Command("scp", args...)
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Copy %v (host: %v) to location: %v", source, sshConfig.Server, dest))
	return nil
}

// CopyToRemote copy local file to remote host(s)
func (c CMServer) CopyToRemote(source string, dest string, filteredHosts map[string]bool, gateway string, inventory *Inventory) {
	connectionProfileID := c.ConnectionProfile
	if len(connectionProfileID) == 0 {
		fmt.Println("No connection profile is attached for the active CM server entry!")
		os.Exit(1)
	}
	connectionProfile := GetConnectionProfileByID(connectionProfileID)
	var hosts map[string]bool
	if len(filteredHosts) > 0 {
		hosts = filteredHosts
	} else {
		hosts = c.GetFilteredHosts(Filter{}, inventory)
	}
	var wg sync.WaitGroup
	wg.Add(len(hosts))
	for host := range hosts {
		ssh := createSSHConfig(connectionProfile, host, gateway)
		go func(ssh *easyssh.MakeConfig, source string, dest string, host string) {
			defer wg.Done()
			err := ssh.Scp(source, dest)
			// Handle errors
			if err != nil {
				errMsg := fmt.Sprintf("Can't run remote command on host '%v (scp %v to %v)", host, source, dest)
				fmt.Println(errMsg)
			} else {
				succMsg := fmt.Sprintf("Copying to remote host '%v' is successful. (from - %v, to %v)", host, source, dest)
				fmt.Println(succMsg)
			}
		}(ssh, source, dest, host)
	}
	wg.Wait()
}

func createSSHConfig(connectionProfile ConnectionProfile, host string, gatewayHost string) *easyssh.MakeConfig {
	if len(gatewayHost) > 0 {
		return &easyssh.MakeConfig{
			User:    connectionProfile.Username,
			Server:  host,
			KeyPath: connectionProfile.KeyPath,
			Port:    strconv.Itoa(connectionProfile.Port),
			Timeout: 60 * time.Second,
			Proxy: easyssh.DefaultConfig{
				User:    connectionProfile.Username,
				Server:  gatewayHost,
				Port:    strconv.Itoa(connectionProfile.Port),
				KeyPath: connectionProfile.KeyPath,
				Timeout: 60 * time.Second,
			},
		}
	}
	return &easyssh.MakeConfig{
		User:    connectionProfile.Username,
		Server:  host,
		KeyPath: connectionProfile.KeyPath,
		Port:    strconv.Itoa(connectionProfile.Port),
		Timeout: 60 * time.Second,
	}
}
