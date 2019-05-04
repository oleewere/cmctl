package cm

import (
	"fmt"
	"os"
	"strconv"
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

func (c CMServer) RunGatewayCMCommand(command string) RemoteResponse {
	connectionProfileId := c.ConnectionProfile
	if len(connectionProfileId) == 0 {
		fmt.Println("No connection profile is attached for the active ambari server entry!")
		os.Exit(1)
	}
	connectionProfile := GetConnectionProfileById(connectionProfileId)
	var wg sync.WaitGroup
	wg.Add(1)

	responses := make(map[string]RemoteResponse)
	ssh := createSshConfig(connectionProfile, c.Hostname, "")
	go func(ssh *easyssh.MakeConfig, command string, responses map[string]RemoteResponse) {
		defer wg.Done()
		stdout, stderr, done, err := ssh.Run(command)
		// Handle errors
		if err != nil {
			panic("Can't run remote command: " + err.Error())
		} else {
			if len(stdout) == 0 {
				fmt.Println("Stdout response is emtpy for remote command.")
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

// RunRemoteHostCommand executes bash commands on CM hosts
func (c CMServer) RunRemoteHostCommand(command string, filteredHosts map[string]bool, skipJump bool) map[string]RemoteResponse {
	connectionProfileId := c.ConnectionProfile
	if len(connectionProfileId) == 0 {
		fmt.Println("No connection profile is attached for the active ambari server entry!")
		os.Exit(1)
	}
	connectionProfile := GetConnectionProfileById(connectionProfileId)
	var hosts map[string]bool
	if len(filteredHosts) > 0 {
		hosts = filteredHosts
	} else {
		hosts = c.GetFilteredHosts(Filter{})
	}
	response := make(map[string]RemoteResponse)
	var wg sync.WaitGroup
	wg.Add(len(hosts))
	gatewayHost := ""
	if c.UseGateway {
		gatewayHost = c.Hostname
	}
	for host := range hosts {
		ssh := createSshConfig(connectionProfile, host, gatewayHost)
		go func(ssh *easyssh.MakeConfig, command string, host string, response map[string]RemoteResponse) {
			defer wg.Done()
			stdout, stderr, done, err := ssh.Run(command, 60)
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

func createSshConfig(connectionProfile ConnectionProfile, host string, gatewayHost string) *easyssh.MakeConfig {
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
