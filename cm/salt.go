package cm

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const saltBinaryPrefixDefault = "sudo /opt/salt_*/bin"
const saltBinaryDefault = "salt"
const saltScriptsFolderDefault = "/srv"

// ExecuteSaltCommand execute a command on all minions
func (c CMServer) ExecuteSaltCommand(command string, rawCommand string, saltBinaryPathPrefix string, saltBinary string) {
	if len(saltBinary) == 0 {
		saltBinary = saltBinaryDefault
	}
	if len(saltBinaryPathPrefix) == 0 {
		saltBinaryPathPrefix = saltBinaryPrefixDefault
	}
	if len(rawCommand) == 0 {
		rawCommand = fmt.Sprintf("'*' cmd.run '%v'", command)
	}
	saltCommand := fmt.Sprintf("sudo %v/%v %v", saltBinaryPathPrefix, saltBinary, rawCommand)
	c.RunGatewayCMCommand(saltCommand, true, true)
}

// SyncSaltScripts sync salt scripts - delete/backup old scripts then upload new ones from local salt folder
func (c CMServer) SyncSaltScripts(source string, target string, filter string, clear bool) {
	fmt.Println("Starting salt sync ...")
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	home := usr.HomeDir
	if len(target) == 0 {
		target = saltScriptsFolderDefault
	}

	sourcePath := strings.Replace(source, "~", home, -1)
	if len(sourcePath) > 0 {
		checkFolder(sourcePath)
	}
	sourceTar := "/tmp/salt.tar"
	if len(filter) == 0 {
		fmt.Println(fmt.Sprintf("Creating tar archive from '%v' folder.", sourcePath))
		RunLocalCommand("tar", "-cf", sourceTar, "-C", filepath.Dir(sourcePath), "salt")
	} else {
		subFolder := path.Join(sourcePath, filter)
		fmt.Println(fmt.Sprintf("Creating tar archive from '%v' folder.", subFolder))
		checkFolder(subFolder)
		sourceTar = fmt.Sprintf("/tmp/%v.tar", filter)
		RunLocalCommand("tar", "-cf", sourceTar, "-C", sourcePath, filter)
	}
	c.CopyToCMGateway(sourceTar, sourceTar)

	remoteSaltPath := fmt.Sprintf("%v", target)
	if len(filter) > 0 {
		remoteSaltPath = fmt.Sprintf("%v/%v", remoteSaltPath, filter)
	}
	t := time.Now()

	if !clear {
		var backupTar string
		var remoteCommand string
		timestamp := t.Format("20060102150405")
		if len(filter) > 0 {
			backupTar = fmt.Sprintf("/tmp/salt-%v-%v.tar", filter, timestamp)
			remoteCommand = fmt.Sprintf("sudo tar -cf %v -C %v %v", backupTar, target, filter)
		} else {
			backupTar = fmt.Sprintf("/tmp/salt-%v.tar", timestamp)
			remoteCommand = fmt.Sprintf("sudo tar -cf %v -C %v .", backupTar, target)
		}
		fmt.Println(fmt.Sprintf("Creating tar archive ('%v') from '%v' remote folder.", backupTar, remoteSaltPath))
		fmt.Println(fmt.Sprintf("Running remote command: %v", remoteCommand))
		c.RunGatewayCMCommand(remoteCommand, true, false)
	}
	deleteSaltFolderCommand := fmt.Sprintf("sudo rm -rf %v/*", remoteSaltPath)
	c.RunGatewayCMCommand(deleteSaltFolderCommand, true, false)

	unpackSaltCommand := fmt.Sprintf("sudo mkdir -p %v && sudo tar xf %v -C %v --strip 1", remoteSaltPath, sourceTar, remoteSaltPath)
	fmt.Println(fmt.Sprintf("Running remote command: %v", unpackSaltCommand))
	c.RunGatewayCMCommand(unpackSaltCommand, true, false)
	setPermissionsCommand := fmt.Sprintf("sudo chown -R root:root %v && sudo chmod -R 644 %v", remoteSaltPath, remoteSaltPath)
	fmt.Println(fmt.Sprintf("Running remote command: %v", setPermissionsCommand))
	c.RunGatewayCMCommand(setPermissionsCommand, true, false)
	fmt.Println("Syncing salt definitions operation is finished.")

}

func checkFolder(path string) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
