package cm

import (
	"fmt"
	"os/exec"

	"github.com/appleboy/easyssh-proxy"
)

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
