# cmctl

[![GoDoc Widget](https://godoc.org/github.com/oleewere/cmctl/cm?status.svg)](https://godoc.org/github.com/oleewere/cmctl/cm)
[![Build Status](https://travis-ci.org/oleewere/cmctl.svg?branch=master)](https://travis-ci.org/oleewere/cmctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/oleewere/cmctl)](https://goreportcard.com/report/github.com/oleewere/cmctl)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Execute REST API calls and remote (salt) commands on Cloudbreak deployed CM clusters

### Installation 

#### Installation on Mac OSX
```bash
brew tap oleewere/repo
brew install cmctl
```

#### Installation on Linux

Using wget:
```bash
CMCTL_VERSION=0.1.0
wget -qO- "https://github.com/oleewere/cmctl/releases/download/v${CMCTL_VERSION}/cmctl_${CMCTL_VERSION}_linux_64-bit.tar.gz" | tar -C /usr/bin -zxv cmctl
```

Using curl:
```bash
CMCTL_VERSION=0.1.0
curl -L -s "https://github.com/oleewere/cmctl/releases/download/v${CMCTL_VERSION}/cmctl_${CMCTL_VERSION}_linux_64-bit.tar.gz" | tar -C /usr/bin -xzv cmctl
```

### Usage

#### Creating a CM server record

First you need to create CM server registry, that will be used to save the details of your Cloudera Manager servers.

```bash
cmctl servers init
```

Then you can create a CM server. (`salt` related operations will be available only if you will set the CM as CB managed, other commands an be used freely againts a non-CB managed cluster). It will ask you to fill the details for the CM server:

```bash
cmctl servers create
```

You can create multiple CM server entries. The tool will always use that one, which has been created last time. You can switch to use an another one by using `servers use` command with selecting the CM server name that you defined in the create command. (see `servers show` output to get the active CM server details)

#### Creating a connection profile

In order to use commands through CB gateway or just use `exec` or `salt` commands, it is required to set up an SSH connection profile:

```bash
cmctl profiles create
```

Then you can attach that profile to your (active_ CM server by `attach` command:

```bash
cmctl attach myprofilename
# or for a non-active one
cmctl attach myprofilename mycmentry
```

#### Execute remote commands on hosts

The `exec` command can be used against any CM clusters:

```bash
cmctl exec -c "echo hello"
```

Against CB managed CM servers, you can also use `salt` (in order to avoid gateway host jump for `exec` commands):

```bash
cmctl salt exec -c "echo hello"
```

#### Generate Ansible host inventory file(s)

If you prefer to use [ansible](https://github.com/ansible/ansible), it is possible to generate ansible compatible host inventory files:

```bash
cmctl inventory generate -d /my/output/directory
# or cmctl inventory generate -o hosts -c mycluster
```

It will generate inventory files per cluster. (`<clustername>.ini` format, can be overriden by `-o` flag, but it requires to use cluster filter if CM registered multiple clusters)

Then you can use the generated file(s) as an ansible inventory file:

```bash
ansible -i my-generated-inventory server -m shell -a "echo hello"
```