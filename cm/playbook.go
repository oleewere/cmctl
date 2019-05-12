package cm

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

const (
	// RemoteCommand remote command type for running commands on the agent hosts
	RemoteCommand = "RemoteCommand"
	// LocalCommand local command type for running commands on localhost
	LocalCommand = "LocalCommand"
	// SaltCommand salt command type for running salt commands on the gateway
	SaltCommand = "SaltCommand"
	// SaltSyncCommand sync salt scripts command
	SaltSyncCommand = "SaltSyncCommand"
	// Download command type for downloading a file from an url
	Download = "Download"
	// Upload command type for uploading files to the agent hosts
	Upload = "Upload"
	// ServiceConfig command type for managing (update) service configuration
	ServiceConfig = "ServiceConfig"
	// RoleConfig command type for managing (update) role configuration
	RoleConfig = "RoleConfig"
	// ServiceCpmmand runs a CM service command (like start or stop) against services
	ServiceCpmmand = "ServiceCommand"
	// RoleCommand runs a CM role command (like start or stop) against roles
	RoleCommand = "RoleCommand"
)

// Playbook contains an array of tasks that will be executed on CM hosts
type Playbook struct {
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	Tasks       []Task  `yaml:"tasks"`
	Inputs      []Input `yaml:"inputs"`
}

// Task represents a task that can be executed on an CM hosts
type Task struct {
	Name           string            `yaml:"name"`
	Type           string            `yaml:"type"`
	Command        string            `yaml:"command"`
	CMServerFilter bool              `yaml:"server"`
	CMAgentFilter  bool              `yaml:"agent"`
	ClusterFilter  string            `yaml:"clusters"`
	HostFilter     string            `yaml:"hosts"`
	ServiceFilter  string            `yaml:"services"`
	RoleTypeFilter string            `yaml:"role_types"`
	RoleNames      string            `yaml:"roles"`
	Parameters     map[string]string `yaml:"parameters,omitempty"`
}

// Input represents a variable that needs to be provided by users (if default value is empty)
type Input struct {
	Name    string `yaml:"name"`
	Default string `yaml:"default,omitempty"`
}

// LoadPlaybookFile read a playbook yaml file and transform it to a Playbook object
func LoadPlaybookFile(location string, varsInput string) Playbook {
	varInputMap := createVarMap(varsInput)
	var dataBytes []byte
	if strings.HasPrefix(location, "http://") || strings.HasPrefix(location, "https://") {
		data, err := DownloadFileInMemory(location)
		fmt.Println("here")
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		dataBytes = data
	} else {
		fileLocation := location
		if strings.HasPrefix("file://", location) {
			fileLocation = strings.TrimPrefix("file://", location)
		}
		data, err := ioutil.ReadFile(fileLocation)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		dataBytes = data
	}
	playbookTempl := Playbook{}
	err := yaml.Unmarshal(dataBytes, &playbookTempl)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(string(dataBytes))
	if len(playbookTempl.Inputs) > 0 {
		for _, input := range playbookTempl.Inputs {
			if varVal, ok := varInputMap[input.Name]; ok {
				fmt.Println(fmt.Sprintf("Found input: %v - %v", input.Name, varVal))
				continue
			}
			if len(input.Default) == 0 {
				varSetByUser := GetStringFlag("", "", fmt.Sprintf("Enter %v", input.Name))
				varInputMap[input.Name] = varSetByUser
				continue
			}
			varInputMap[input.Name] = input.Default
		}
	}
	templ := template.New("playbook template")
	textTemplate, _ := templ.Parse(fmt.Sprintf("%s", dataBytes))
	var tpl bytes.Buffer
	textTemplate.Execute(&tpl, varInputMap)

	playbook := Playbook{}
	err = yaml.Unmarshal(tpl.Bytes(), &playbook)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("[Executing playbook: %v, file: %v]", playbook.Name, location))
	return playbook
}

// ExecutePlaybook runs tasks on CM hosts based on a playbook object
func (c CMServer) ExecutePlaybook(playbook Playbook, inventory *Inventory) {
	tasks := playbook.Tasks
	for _, task := range tasks {
		if len(task.Type) > 0 {
			filteredHosts := make(map[string]bool)
			if !task.CMAgentFilter {
				filter := CreateFilter(task.ClusterFilter, task.ServiceFilter, task.RoleTypeFilter, task.HostFilter, task.CMServerFilter)
				filteredHosts = c.GetFilteredHosts(filter, inventory)
			}
			if task.Type == RemoteCommand {
				c.ExecuteRemoteCommandTask(task, filteredHosts, inventory)
			}
			if task.Type == LocalCommand {
				ExecuteLocalCommandTask(task)
			}
			if task.Type == Download {
				ExecuteDownloadFileTask(task)
			}
			if task.Type == Upload {
				c.ExecuteUploadFileTask(task, filteredHosts, inventory)
			}
			if task.Type == SaltSyncCommand {
				c.ExecuteSaltSyncCommand(task)
			}
			/*
				if task.Type == Config {
					c.ExecuteConfigCommand(task)
				}
				if task.Type == RoleCommand {
					c.ExecuteCMCommand(task)
				}
			*/
		} else {
			if len(task.Name) > 0 {
				fmt.Println(fmt.Sprintf("Type field for task '%s' is required!", task.Name))
			} else {
				fmt.Println("Type field for task is required!")
			}
			os.Exit(1)
		}
	}
}

// ExecuteCMCommand executes an CM command against services or roles
func (c CMServer) ExecuteCMCommand(task Task) {
	if len(task.Command) > 0 {
		useRoleFilter := false
		useServiceFilter := false
		if len(task.RoleTypeFilter) > 0 {
			useRoleFilter = true
		} else if len(task.ServiceFilter) > 0 {
			useServiceFilter = true
		}

		if useRoleFilter {
			filter := CreateFilter(task.ClusterFilter, "", task.RoleTypeFilter, "", false)
			fmt.Println(filter)
			//c.RunRolesOperation(task.Command, filter, useServiceFilter, useRoleFilter, task.Command, true)
		}
		if useServiceFilter {
			filter := CreateFilter(task.ClusterFilter, task.ServiceFilter, "", "", false)
			fmt.Println(filter)
			//c.RunServiceOperation(task.Command, filter, useServiceFilter, useRoleFilter, task.Command, true)
		}
	}
}

/*
// ExecuteConfigCommand executes a configuration upgrade
func (c CMServer) ExecuteConfigCommand(task Task) {
	if task.Parameters != nil {
		haveConfigType := false
		haveConfigKey := false
		haveConfigValue := false
		if configType, ok := task.Parameters["config_type"]; ok {
			haveConfigType = true
			if configKey, ok := task.Parameters["config_key"]; ok {
				haveConfigKey = true
				if configValue, ok := task.Parameters["config_value"]; ok {
					haveConfigValue = true
					a.SetConfig(configType, configKey, configValue)
				}
			}
		}
		if !haveConfigType {
			fmt.Println("'config_type' parameter is required for 'Upload' task")
			os.Exit(1)
		}
		if !haveConfigKey {
			fmt.Println("'config_key' parameter is required for 'Upload' task")
			os.Exit(1)
		}
		if !haveConfigValue {
			fmt.Println("'config_value' parameter is required for 'Upload' task")
			os.Exit(1)
		}
	}
}
*/

// ExecuteRemoteCommandTask executes a remote command on filtered hosts
func (c CMServer) ExecuteRemoteCommandTask(task Task, filteredHosts map[string]bool, inventory *Inventory) {
	if len(task.Command) > 0 {
		fmt.Println("Execute remote command: " + task.Command)
		c.RunRemoteHostCommand(task.Command, filteredHosts, task.CMServerFilter, inventory)
	}
}

// ExecuteSaltSyncCommand sync salt scripts from folder
func (c CMServer) ExecuteSaltSyncCommand(task Task) {
	if task.Parameters != nil {
		if source, ok := task.Parameters["source"]; ok {
			target := saltScriptsFolderDefault
			clear := false
			filter := ""
			if targetVal, ok := task.Parameters["target"]; ok {
				target = targetVal
			}
			if filterVal, ok := task.Parameters["filter"]; ok {
				filter = filterVal
			}
			if clearVal, ok := task.Parameters["clear"]; ok {
				if clearVal == "true" {
					clear = true
				}
			}
			c.SyncSaltScripts(source, target, filter, clear)
		} else {
			fmt.Println("'source' parameter is required for 'SaltSyncCommand' task")
			os.Exit(1)
		}
	}
}

// ExecuteUploadFileTask upload a file to specific (filtered) hosts
func (c CMServer) ExecuteUploadFileTask(task Task, filteredHosts map[string]bool, inventory *Inventory) {
	if task.Parameters != nil {
		haveSourceFile := false
		haveTargetFile := false
		if sourceVal, ok := task.Parameters["source"]; ok {
			haveSourceFile = true
			if targetVal, ok := task.Parameters["target"]; ok {
				haveTargetFile = true
				fmt.Println(fmt.Sprintf("Execute upload file command - source: %s, target: %s",
					task.Parameters["source"], task.Parameters["target"]))
				cmServerGatewayAddress := ""
				if c.UseGateway {
					cmServerGatewayAddress = c.Hostname
				}
				// TODO if server filter is used with gateway ?
				c.CopyToRemote(sourceVal, targetVal, filteredHosts, cmServerGatewayAddress, inventory)
			}
		}
		if !haveSourceFile {
			fmt.Println("'source' parameter is required for 'Upload' task")
			os.Exit(1)
		}
		if !haveTargetFile {
			fmt.Println("'target' parameter is required for 'Upload' task")
			os.Exit(1)
		}

	}
}

// ExecuteLocalCommandTask executes a local shell command
func ExecuteLocalCommandTask(task Task) {
	if len(task.Command) > 0 {
		fmt.Println("Execute local command: " + task.Command)
		splitted := strings.Split(task.Command, " ")
		if len(splitted) == 1 {
			RunLocalCommand(splitted[0])
		} else {
			RunLocalCommand(splitted[0], splitted[1:]...)
		}
	}
}

// ExecuteDownloadFileTask download a file from an url to the local filesystem
func ExecuteDownloadFileTask(task Task) {
	if task.Parameters != nil {
		haveURL := false
		haveFile := false
		if urlVal, ok := task.Parameters["url"]; ok {
			haveURL = true
			if fileVal, ok := task.Parameters["file"]; ok {
				haveFile = true
				fmt.Println(fmt.Sprintf("Execute download file command - url: %s, location: %s",
					task.Parameters["url"], task.Parameters["file"]))
				DownloadFile(fileVal, urlVal)
			}
		}
		if !haveFile {
			fmt.Println("'file' parameter is required for 'Download' task")
			os.Exit(1)
		}
		if !haveURL {
			fmt.Println("'url' parameter is required for 'Download' task")
			os.Exit(1)
		}
	}
}

func createVarMap(varMapStr string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	if len(varMapStr) > 0 {
		var ss []string
		ss = strings.Split(varMapStr, " ")
		for _, pair := range ss {
			z := strings.Split(pair, "=")
			resultMap[z[0]] = z[1]
		}
	}
	return resultMap
}
