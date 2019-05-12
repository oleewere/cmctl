# cmctl

[![GoDoc Widget](https://godoc.org/github.com/oleewere/cmctl/cm?status.svg)](https://godoc.org/github.com/oleewere/cmctl/cm)
[![Build Status](https://travis-ci.org/oleewere/cmctl.svg?branch=master)](https://travis-ci.org/oleewere/cmctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/oleewere/cmctl)](https://goreportcard.com/report/github.com/oleewere/cmctl)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Execute REST API calls and remote (salt) commands on Cloudbreak deployed CM clusters

## Installation 

### Installation on Mac OSX
```bash
brew tap oleewere/repo
brew install cmctl
```

### Installation on Linux

Using wget:
```bash
CMCTL_VERSION=0.3.0
wget -qO- "https://github.com/oleewere/cmctl/releases/download/v${CMCTL_VERSION}/cmctl_${CMCTL_VERSION}_linux_64-bit.tar.gz" | tar -C /usr/bin -zxv cmctl
```

Using curl:
```bash
CMCTL_VERSION=0.3.0
curl -L -s "https://github.com/oleewere/cmctl/releases/download/v${CMCTL_VERSION}/cmctl_${CMCTL_VERSION}_linux_64-bit.tar.gz" | tar -C /usr/bin -xzv cmctl
```

## Usage

### Creating a CM server record

First you need to create CM server registry, that will be used to save the details of your Cloudera Manager servers.

```bash
cmctl servers init
```

Then you can create a CM server. (`salt` related operations will be available only if you will set the CM as CB managed, other commands an be used freely againts a non-CB managed cluster). It will ask you to fill the details for the CM server:

```bash
cmctl servers create
```

You can create multiple CM server entries. The tool will always use that one, which has been created last time. You can switch to use an another one by using `servers use` command with selecting the CM server name that you defined in the create command. (see `servers show` output to get the active CM server details)

### Creating a connection profile

In order to use commands through CB gateway or just use `exec` or `salt` commands, it is required to set up an SSH connection profile:

```bash
cmctl profiles create
```

Then you can attach that profile to your (active_ CM server by `profiles attach` command:

```bash
cmctl profiles attach myprofilename
# or for a non-active CM server
cmctl profiles attach myprofilename mycmentry
```

### Execute remote commands on hosts

The `exec` command can be used against any CM clusters:

```bash
cmctl exec -c "echo hello"
```

Against CB managed CM servers, you can also use `salt` (in order to avoid gateway host jump for `exec` commands):

```bash
cmctl salt exec -c "echo hello"
```

### Generate Ansible host inventory file(s)

If you prefer to use [ansible](https://github.com/ansible/ansible), it is possible to generate ansible compatible host inventory files:

```bash
cmctl inventory create -d /my/output/directory
# or cmctl inventory create -o hosts -c mycluster
```

It will generate inventory files per cluster. (`<clustername>.ini` format, can be overriden by `-o` flag, but it requires to use cluster filter if CM registered multiple clusters)

Then you can use the generated file(s) as an ansible inventory file:

```bash
ansible -i my-generated-inventory.ini server -m shell -a "echo hello"
```

The ansible inventory hosts file can be used for (some) `cmctl` commands as well (with that you can avoid REST API calls against CM for gathering proper service / role hosts):

```bash
cmctl -i my-generated-inventory.ini exec -c "echo hello" --services zookeeper
```

### Using playbooks

By a specific `yaml` file format, you can define playbooks for yourself (in order to execute a list of commands). 

```bash
# -i is optional, can work without inventory (using CM api)
# note: file:// or http:// file for files (can be comma separated)
cmctl playbook -i my-generated-inventory.ini -f examples/echo.yaml
```

The following commands are supported (see `type` task field):

#### LocalCommand

`LocalCommand` executes a local command.

```yaml
name: "Echo hello loally"
tasks:
  - name: "Echo hello loally"
    debug: true
    type: LocalCommand
    command: "echo hello local"
```

#### RemoteCommand

`RemoteCommand` executes commands on remote CM hosts (the hosts can be filtered based on server/hosts/services/role types)

```yaml
name: "Echo hello on remote hosts"
tasks:
  - name: "Echo hello from all agents"
    debug: true
    type: RemoteCommand
    command: "echo hello CM agent"
  - name: "Echo hello from server"
    debug: true
    type: RemoteCommand
    command: "echo hello CM server"
    server: true
  - name: "Echo hello from agents"
    debug: true
    hosts: myhost1,myhost2
    type: RemoteCommand
    command: "echo hello CM host agent"
  - name: "Echo hello from ZK hosts"
    debug: true
    service: zookeeper
    type: RemoteCommand
    command: "echo hello CM ZK host"
  - name: "Echo hello from ZK server hosts"
    debug: true
    service: zookeeper
    role: SERVER
    type: RemoteCommand
    command: "echo hello ZK server host"
```

#### SaltCommand

`SaltCommand` executes salt command on gateway host against salt minions. The simple command uses `'cmd.run'` on `'*'` by default. (using `raw`, you can override that behavior)

```yaml
name: "Echo hello through salt"
tasks:
  - name: "Echo hello through salt by simple command"
    debug: true
    type: SaltCommand
    command: "echo hello salt"
  - name: "Echo hello through salt by raw command"
    debug: true
    type: SaltCommand
    command: "'*' cmd.run 'echo hello raw salt'"
    parameters:
      raw: true
```

#### SaltSynCommand

`SaltSyncCommand` synchronize salt script definitions on the gateway (from local folder)

```yaml
name: "Refresh salt scripts"
inputs:
  - name: ScriptsLocation
    default: "~/Projects/cloudbreak/orchestrator-salt/src/main/resources/salt"
tasks:
  - name: "Sync salt scripts with {{ .ScriptsLocation }}"
    type: SaltSyncCommand
    parameters:
      source: "{{ .ScriptsLocation }}"
      clear: false
```

#### ServiceCpmmand

`ServiceCpmmand` exeutes service commands against CM Rest API.

```yaml
name: "Restart ZK service"
tasks:
  - name: "Restart ZK service"
    debug: true
    type: ServiceCommand
    service: "zookeeper"
    command: "restart"
```

#### RoleCommand

`RoleCommand` exeutes role commands against CM Rest API.

```yaml
name: "Restart ZK server role"
tasks:
  - name: "Restart ZK server role"
    debug: true
    type: RoleCommand
    service: "zookeeper"
    role: "SERVER"
    command: "restart"
```

#### ServiceConfigUpdate

`ServiceConfigUpdate` updates a service configuration by providing `key` and `value` in `configs` list. Instead of using `value`, you can use `clear` flag to clear config value.

```yaml
name: "Update service configs example"
tasks:
  - name: "Update ZK service config"
    debug: true
    type: ServiceConfigUpdate
    service: zookeeper
    configs:
    - key: smon_derived_configs_safety_valve
      value: true
  - name: "Clear ZK servie config"
    debug: true
    type: ServiceConfigUpdate
    service: zookeeper
    configs:
      - key: smon_derived_configs_safety_valve
    parameters:
      clear: true
```

#### RoleConfigUpdate

`RoleConfigUpdate` same as `ServiceConfigUpdate`, but that will update the role configuration, therefore you will need to provide `role` filter (as comma separated role types)

```yaml
name: "Update role configs example"
tasks:
  - name: "Update ZK role config"
    debug: true
    type: RoleConfigUpdate
    service: zookeeper
    role: SERVER
    configs:
    - key: maxSessionTimeout
      value: 60000
```

## License
GNU General Public License v3.0
