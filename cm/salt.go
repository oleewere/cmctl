package cm

import "fmt"

const saltBinaryPrefixDefault = "sudo /opt/salt_*/bin"
const saltBinaryDefault = "salt"

// ExecuteSaltCommand execute a command on all minions
func (c CMServer) ExecuteSaltCommand(command string, saltBinaryPathPrefix string, saltBinary string) {
	if len(saltBinary) == 0 {
		saltBinary = saltBinaryDefault
	}
	if len(saltBinaryPathPrefix) == 0 {
		saltBinaryPathPrefix = saltBinaryPrefixDefault
	}
	saltCommand := fmt.Sprintf("sudo %v/%v '*' cmd.run '%v'", saltBinaryPathPrefix, saltBinary, command)
	c.RunGatewayCMCommand(saltCommand, true)
}
