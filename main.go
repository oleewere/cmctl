package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/oleewere/cmctl/cm"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// Version that will be generated during the build as a constant
var Version string

// GitRevString that will be generated during the build as a constant - represents git revision value
var GitRevString string

// DefaultApiVersion number that is used as api versions in REST api requests e.g. (32): api/v32/..
const DefaultApiVersion = 32

func main() {
	app := cli.NewApp()
	app.Name = "cmctl"
	app.Usage = "CLI tool for handle CM clusters"
	app.EnableBashCompletion = true
	app.UsageText = "cmctl command [command options] [arguments...]"
	if len(Version) > 0 {
		app.Version = Version
	} else {
		app.Version = "0.1.0"
	}
	if len(GitRevString) > 0 {
		app.Version = app.Version + fmt.Sprintf(" (git short hash: %v)", GitRevString)
	}

	app.Commands = []cli.Command{}

	registryCommand := cli.Command{
		Name:  "registry",
		Usage: "CM server registry DB related operations",
		Subcommands: []cli.Command{
			{
				Name:  "init",
				Usage: "Initialize CM server database",
				Action: func(c *cli.Context) error {
					cm.CreateCMRegistryDb()
					fmt.Println("CM registry DB has been initialized.")
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "Register new CM server entry",
				Action: func(c *cli.Context) error {
					name := cm.GetStringFlag(c.String("name"), "", "Enter CM server name")
					cmEntryId := cm.GetCMEntryId(name)
					if len(cmEntryId) > 0 {
						fmt.Println("CM server entry already exists with id " + name)
						os.Exit(1)
					}
					host := cm.GetStringFlag(c.String("host"), "", "Enter CM server/gateway address")
					useGatewayStr := cm.GetStringFlag(c.String("gateway"), "y", "Is CM Cloudbreak managed?")
					useGateway := cm.EvaluateBoolValueFromString(useGatewayStr)
					var port int
					var protocol string
					if useGateway {
						protocol = "https"
						port = 7180
					} else {
						protocol = strings.ToLower(cm.GetStringFlag(c.String("protocol"), "http", "Enter CM protocol"))
						if protocol != "http" && protocol != "https" {
							fmt.Println("Use 'http' or 'https' value for protocol option")
							os.Exit(1)
						}
						defaultPort := "7180"
						if protocol == "https" {
							defaultPort = "443"
						}
						portStr := cm.GetStringFlag(c.String("port"), defaultPort, "Enter CM port")
						var err error
						port, err = strconv.Atoi(portStr)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					}
					username := strings.ToLower(cm.GetStringFlag(c.String("username"), "admin", "Enter CM user"))
					password := cm.GetPassword(c.String("password"), "Enter CM user password")

					apiVersion := cm.GetStringFlag(c.String("verion"), strconv.Itoa(DefaultApiVersion), "Enter CM API version")
					apiVersionNum, err := strconv.Atoi(apiVersion)

					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					cm.DeactiveAllCMRegistry()
					cm.RegisterNewCMEntry(name, host, port, protocol,
						username, password, useGateway, apiVersionNum)
					fmt.Println("New CM server entry has been created: " + name)
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "name", Usage: "Name of the CM server entry"},
					cli.StringFlag{Name: "host", Usage: "Hostname of the CM server"},
					cli.StringFlag{Name: "gateway", Usage: "Managed by Cloudbreak (y/n)"},
					cli.StringFlag{Name: "port", Usage: "Port for CM Server"},
					cli.StringFlag{Name: "protocol", Usage: "Protocol for CM REST API: http/https"},
					cli.StringFlag{Name: "username", Usage: "User name for CM server"},
					cli.StringFlag{Name: "password", Usage: "Password for CM user"},
					cli.StringFlag{Name: "version", Usage: "CM Api Version"},
				},
			},
			{
				Name:  "update",
				Usage: "Update a CM server entry",
				Action: func(c *cli.Context) error {
					args := c.Args()
					if len(args) == 0 {
						fmt.Println("Provide 1 argument (<cm_server_name>)")
						os.Exit(1)
					}
					cmServerName := args.Get(0)
					existingCmServer := cm.GetCMById(cmServerName)
					if len(existingCmServer.Name) == 0 {
						fmt.Println("CM server entry does not exist with id " + cmServerName)
						os.Exit(1)
					}
					useGatewayDefault := "n"
					if existingCmServer.UseGateway {
						useGatewayDefault = "y"
					}
					host := cm.GetStringFlag(c.String("host"), existingCmServer.Hostname, "Enter CM server/gateway address")
					useGatewayStr := cm.GetStringFlag(c.String("gateway"), useGatewayDefault, "Is CM Cloudbreak managed?")
					useGateway := cm.EvaluateBoolValueFromString(useGatewayStr)
					var port int
					var protocol string
					if useGateway {
						protocol = "https"
						port = 7180
					} else {
						protocol = strings.ToLower(cm.GetStringFlag(c.String("protocol"), existingCmServer.Protocol, "Enter CM protocol"))
						if protocol != "http" && protocol != "https" {
							fmt.Println("Use 'http' or 'https' value for protocol option")
							os.Exit(1)
						}
						portStr := cm.GetStringFlag(c.String("port"), strconv.Itoa(existingCmServer.Port), "Enter CM port")
						var err error
						port, err = strconv.Atoi(portStr)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					}
					username := strings.ToLower(cm.GetStringFlag(c.String("username"), existingCmServer.Username, "Enter CM user"))
					password := cm.GetPassword(c.String("password"), "Enter CM user password")

					apiVersion := cm.GetStringFlag(c.String("verion"), strconv.Itoa(DefaultApiVersion), "Enter CM API version")
					apiVersionNum, err := strconv.Atoi(apiVersion)

					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					cm.UpdateCMEntry(cmServerName, host, port, protocol,
						username, password, useGateway, apiVersionNum, existingCmServer.ConnectionProfile)
					fmt.Println("CM server entry has been updated: " + cmServerName)
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "host", Usage: "Hostname of the CM server"},
					cli.StringFlag{Name: "gateway", Usage: "Managed by Cloudbreak (y/n)"},
					cli.StringFlag{Name: "port", Usage: "Port for CM Server"},
					cli.StringFlag{Name: "protocol", Usage: "Protocol for CM REST API: http/https"},
					cli.StringFlag{Name: "username", Usage: "User name for CM server"},
					cli.StringFlag{Name: "password", Usage: "Password for CM user"},
					cli.StringFlag{Name: "version", Usage: "CM Api Version"},
				},
			},
			{
				Name:  "delete",
				Usage: "De-register an existing CM server entry",
				Action: func(c *cli.Context) error {
					if len(c.Args()) == 0 {
						fmt.Println("Provide a registry name argument for use command. e.g.: delete vagrant")
						os.Exit(1)
					}
					name := c.Args().First()
					cmEntryId := cm.GetCMEntryId(name)
					if len(cmEntryId) == 0 {
						fmt.Println("CM registry entry does not exist with id " + name)
						os.Exit(1)
					}
					cm.DeRegisterCMEntry(name)
					fmt.Println("CM registry de-registered with id: " + name)
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "name", Usage: "name of the CM registry entry"},
				},
			},
			{
				Name:  "use",
				Usage: "Use selected CM server",
				Action: func(c *cli.Context) error {
					if len(c.Args()) == 0 {
						fmt.Println("Provide a server entry name argument for use command. e.g.: use vagrant")
						os.Exit(1)
					}
					name := c.Args().First()
					cmEntryId := cm.GetCMEntryId(name)
					if len(cmEntryId) == 0 {
						fmt.Println("CM server entry does not exist with id " + name)
						os.Exit(1)
					}
					cm.DeactiveAllCMRegistry()
					cm.ActiveCMRegistry(name)
					fmt.Println("CM server entry selected with id: " + name)
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "name", Usage: "name of the CM registry entry"},
				},
			},
			{
				Name:  "clear",
				Usage: "Drop all CM server records",
				Action: func(c *cli.Context) error {
					cm.DropCMRegistryRecords()
					fmt.Println("CM server entries dropped.")
					return nil
				},
			},
			{
				Name:  "show",
				Usage: "Show active CM server details",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					var tableData [][]string
					if len(cmServer.Name) > 0 {
						gatewayVal := "false"
						if cmServer.UseGateway {
							gatewayVal = "true"
						}
						tableData = append(tableData, []string{cmServer.Name, cmServer.Hostname, gatewayVal, strconv.Itoa(cmServer.Port), cmServer.Protocol,
							cmServer.Username, "********", cmServer.ConnectionProfile, "true"})
					}
					printTable("ACTIVE CM SERVER:", []string{"Name", "HOSTNAME", "GATEWAY", "PORT", "PROTOCOL", "USER", "PASSWORD", "PROFILE", "ACTIVE"}, tableData, c)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all registered CM servers",
				Action: func(c *cli.Context) error {
					cmServerEntries := cm.ListCMRegistryEntries()
					var tableData [][]string
					for _, cmServer := range cmServerEntries {
						activeValue := "false"
						if cmServer.Active {
							activeValue = "true"
						}
						gatewayVal := "false"
						if cmServer.UseGateway {
							gatewayVal = "true"
						}
						tableData = append(tableData, []string{cmServer.Name, cmServer.Hostname, gatewayVal, strconv.Itoa(cmServer.Port), cmServer.Protocol,
							cmServer.Username, "********", cmServer.ConnectionProfile, activeValue})
					}
					printTable("CM SERVERS:", []string{"Name", "HOSTNAME", "GATEWAY", "PORT", "PROTOCOL", "USER", "PASSWORD", "PROFILE", "ACTIVE"}, tableData, c)
					return nil
				},
			},
		},
	}

	serviesCommand := cli.Command{
		Name:  "services",
		Usage: "Service related operations",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all services per cluster",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					clusters := cmServer.ListClusters()
					if len(clusters) == 0 {
						fmt.Println(fmt.Sprintf("Not found any clusters for CM server with id '%v'", cmServer.Name))
						os.Exit(1)
					}
					// TODO: fill all data first - then print them
					for index, cluster := range clusters {
						var tableData [][]string
						services := cmServer.ListServices(cluster.Name)
						for _, service := range services {
							tableData = append(tableData, []string{service.Name, service.DisplayName, service.State, cluster.Name, service.StaleConfig})
						}
						prefixData := ""
						if index > 0 {
							prefixData = "\n"
						}
						header := fmt.Sprintf("%vSERVICES - %v (%v):", prefixData, cluster.DisplayName, cluster.Name)
						printTable(header, []string{"ID", "NAME", "STATE", "CLUSTER", "CONFIG"}, tableData, c)
					}
					return nil
				},
			},
		},
	}

	profileCommand := cli.Command{
		Name:  "profiles",
		Usage: "Connection profiles related opreations",
		Subcommands: []cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create new connection profile",
				Action: func(c *cli.Context) error {
					name := cm.GetStringFlag(c.String("name"), "", "Enter connection profile name")
					connProfileId := cm.GetConnectionProfileEntryId(name)
					if len(connProfileId) > 0 {
						fmt.Println("Connection profile entry already exists with id " + name)
						os.Exit(1)
					}
					keyPath := cm.GetStringFlag(c.String("key_path"), "", "Enter ssh key path")
					usr, err := user.Current()
					if err != nil {
						panic(err)
					}
					home := usr.HomeDir
					keyPath = strings.Replace(keyPath, "~", home, -1)
					if len(keyPath) > 0 {
						if _, err := os.Stat(keyPath); err != nil {
							if os.IsNotExist(err) {
								fmt.Println(err)
								os.Exit(1)
							}
						}
					}
					portStr := cm.GetStringFlag(c.String("port"), "22", "Enter ssh port")
					port, err := strconv.Atoi(portStr)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					userName := cm.GetStringFlag(c.String("username"), "cloudbreak", "Enter ssh username")

					cm.RegisterNewConnectionProfile(name, keyPath, port, userName)
					fmt.Println("New connection profile entry has been created: " + name)
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "name", Usage: "Name of the CM server entry"},
					cli.StringFlag{Name: "key_path", Usage: "Hostname of the CM server"},
					cli.StringFlag{Name: "port", Usage: "Port for CM Server"},
					cli.StringFlag{Name: "username", Usage: "Protocol for CM REST API: http/https"},
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all connection profile entries",
				Action: func(c *cli.Context) error {
					connectionProfiles := cm.ListConnectionProfileEntries()
					var tableData [][]string
					for _, profile := range connectionProfiles {
						tableData = append(tableData, []string{profile.Name, profile.KeyPath, strconv.Itoa(profile.Port), profile.Username})
					}
					printTable("CONNECTION PROFILES:", []string{"NAME", "KEY", "PORT", "USERNAME"}, tableData, c)
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Delete a connection profile entry by id",
				Action: func(c *cli.Context) error {
					if len(c.Args()) == 0 {
						fmt.Println("Provide a profile name argument for use command. e.g.: delete vagrant")
						os.Exit(1)
					}
					name := c.Args().First()
					profileEntryId := cm.GetConnectionProfileEntryId(name)
					if len(profileEntryId) == 0 {
						fmt.Println("Connection profile entry does not exist with id " + name)
						os.Exit(1)
					}
					cm.DeRegisterConnectionProfile(profileEntryId)
					msg := fmt.Sprintf("Connection profile '%s' has been deleted successfully", profileEntryId)
					fmt.Println(msg)
					return nil
				},
			},
			{
				Name:    "clear",
				Aliases: []string{"cl"},
				Usage:   "Delete all connection profile entries",
				Action: func(c *cli.Context) error {
					cm.DropConnectionProfileRecords()
					fmt.Println("All connection profile records has been dropped")
					return nil
				},
			},
		},
	}

	attachCommand := cli.Command{
		Name:  "attach",
		Usage: "Attach a profile to a CM server entry",
		Action: func(c *cli.Context) error {
			args := c.Args()
			if len(args) == 0 {
				fmt.Println("Provide at least 1 argument (<profile>), or 2 (<profile> and <cmEntry>)")
				os.Exit(1)
			}
			profileId := args.Get(0)
			var cmServer cm.CMServer
			if len(args) == 1 {
				cmServer = cm.GetActiveCM()
				if len(cmServer.Name) == 0 {
					fmt.Println("No active CM selected")
					os.Exit(1)
				}
			} else {
				cmServerId := args.Get(1)
				cm.GetCMById(cmServerId)
				if len(cmServer.Name) == 0 {
					fmt.Println("Cannot find specific CM server entry")
					os.Exit(1)
				}
			}
			profile := cm.GetConnectionProfileById(profileId)
			if len(profile.Name) == 0 {
				fmt.Println("Cannot find specific connection profile entry")
				os.Exit(1)
			}

			cm.SetProfileIdForCMEntry(cmServer.Name, profile.Name)
			msg := fmt.Sprintf("Attach profile '%s' to '%s'", profile.Name, cmServer.Name)
			fmt.Println(msg)
			return nil
		},
	}

	clustersCommand := cli.Command{
		Name:  "clusters",
		Usage: "Cluster related opterations",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all registered CM clusters",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					clusters := cmServer.ListClusters()
					var tableData [][]string
					for _, cluster := range clusters {
						tableData = append(tableData, []string{cluster.Name, cluster.DisplayName, cluster.Version, cluster.Type})
					}
					printTable("CLUSTERS:", []string{"NAME", "DISPLAY NAME", "VERSION", "TYPE"}, tableData, c)
					return nil
				},
			},
			{
				Name:    "export",
				Aliases: []string{"x"},
				Usage:   "Export cluster configuration to a cluster template json",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					var clusterTemplate []byte

					clusters := cmServer.ListClusters()
					if len(clusters) == 0 {
						fmt.Println("Not found any cluster resources!")
						os.Exit(1)
					}
					if len(clusters) > 1 && len(c.String("cluster")) == 0 {
						fmt.Println("Please provide cluster filter!")
						os.Exit(1)
					}
					var clusterFilter string
					if len(clusters) > 1 {
						clusterFilter = c.String("cluster")
					} else {
						clusterFilter = clusters[0].Name
					}

					clusterTemplate = cmServer.ExportClusterTemplate(clusterFilter)
					if len(c.String("file")) > 0 {
						err := ioutil.WriteFile(c.String("file"), clusterTemplate, 0644)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						return nil
					}
					printJson(clusterTemplate)
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "cluster, c", Usage: "Cluster name for the template"},
					cli.StringFlag{Name: "file, f", Usage: "File output for the generated JSON"},
				},
			},
		},
	}

	hostsCommand := cli.Command{
		Name:  "hosts",
		Usage: "Hosts resource related operations",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all registered CM hosts",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					hosts := cmServer.ListHosts()
					var tableData [][]string
					for _, host := range hosts {
						tableData = append(tableData, []string{host.HostName, host.IPAddress, host.ClusterName, host.CommissionState, host.RackID})
					}
					printTable("HOSTS:", []string{"HOSTNAME", "IP", "CLUSTER", "STATE", "RACK ID"}, tableData, c)
					return nil
				},
			},
		},
	}

	rolesCommand := cli.Command{
		Name:  "roles",
		Usage: "Role related opreations",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all services per cluster",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					filter := cm.CreateFilter(c.String("clusters"), c.String("services"), "", false)

					deployment := cmServer.GetDeployment()
					clusterServiceRoleMap := deployment.ClusterServiceRoleMap
					var counter int
					for cluster, serviceMap := range clusterServiceRoleMap {
						if len(filter.Clusters) > 0 && !contains(cluster, filter.Clusters) {
							continue
						}
						for service, roles := range serviceMap.RolesMap {
							if len(filter.Services) > 0 && !contains(service, filter.Services) {
								continue
							}
							var tableData [][]string
							for _, role := range roles {
								tableData = append(tableData, []string{role.Name, role.Type, role.HostName, role.ServiceName, role.ClusterName})
							}
							prefixData := ""
							if counter > 0 {
								prefixData = "\n"
							}
							counter++
							header := fmt.Sprintf("%vRoles - %v (cluster: %v):", prefixData, service, cluster)
							printTable(header, []string{"NAME", "TYPE", "HOSTNAME", "SERVICE", "CLUSTER"}, tableData, c)
						}

					}
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "clusters, c", Usage: "Cluster filter for roles"},
					cli.StringFlag{Name: "services, s", Usage: "Service filter for roles"},
				},
			},
		},
	}

	usersCommand := cli.Command{
		Name:  "users",
		Usage: "User related opreations",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all services per cluster",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					users := cmServer.GetUsers()
					if len(users) > 0 {
						var tableData [][]string
						for _, user := range users {
							authRoles := user.AuthRoles
							authRolesStrList := make([]string, 0)
							for _, authRole := range authRoles {
								authRolesStrList = append(authRolesStrList, authRole.Name)
							}
							authRolesStr := strings.Join(authRolesStrList[:], ",")
							tableData = append(tableData, []string{user.Name, authRolesStr})
						}
						header := "USERS: "
						printTable(header, []string{"NAME", "AUTH ROLES"}, tableData, c)
					} else {
						fmt.Println("Not found users!")
						os.Exit(1)
					}

					return nil
				},
			},
		},
	}

	saltCommand := cli.Command{
		Name:  "salt",
		Usage: "Execute salt commands on the CB gateway",
		Subcommands: []cli.Command{
			{
				Name:    "exec",
				Aliases: []string{"x"},
				Usage:   "Execute salt commands on the CB gateway",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					if !cmServer.UseGateway {
						fmt.Println("Cannot run salt commands as CM Server is not CB managed!")
						os.Exit(1)
					}
					command := c.String("command")
					if len(command) == 0 {
						fmt.Println("Command paramter is missing! (use 'command' or 'c')")
						os.Exit(1)
					}
					cmServer.ExecuteSaltCommand(command, c.String("prefix"), c.String("binary"))
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "command, c", Usage: "Command to execute on the minions"},
					cli.StringFlag{Name: "binary, b", Usage: "Salt binary to use (default: salt)"},
					cli.StringFlag{Name: "prefix, p", Usage: "Salt binary prefix path (default: /opt/salt_*/bin)"},
				},
			},
			{
				Name:  "sync",
				Usage: "Synchronize salt resources from CB source",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					if !cmServer.UseGateway {
						fmt.Println("Cannot run salt commands as CM Server is not CB managed!")
						os.Exit(1)
					}
					fmt.Println("TODO: sync salt resources ...")
					return nil
				},
			},
		},
	}

	execCommand := cli.Command{
		Name:  "exec",
		Usage: "Execute commands on all (or specific) hosts",
		Action: func(c *cli.Context) error {
			cmServer := cm.GetActiveCM()
			validateActiveCM(cmServer)
			args := c.Args()
			command := ""
			for _, arg := range args {
				command += arg
			}
			filter := cm.CreateFilter(c.String("clusters"), c.String("services"), c.String("hosts"), c.Bool("server"))
			hosts := cmServer.GetFilteredHosts(filter)
			cmServer.RunRemoteHostCommand(command, hosts, filter.Server)
			return nil
		},
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "server", Usage: "Filter on CM server"},
			cli.StringFlag{Name: "clusters, c", Usage: "Filter on clusters (comma separated)"},
			cli.StringFlag{Name: "services, s", Usage: "Filter on services (comma separated)"},
			cli.StringFlag{Name: "hosts", Usage: "Filter on hosts (comma separated)"},
		},
	}

	app.Commands = append(app.Commands, registryCommand)
	app.Commands = append(app.Commands, profileCommand)
	app.Commands = append(app.Commands, attachCommand)
	app.Commands = append(app.Commands, clustersCommand)
	app.Commands = append(app.Commands, hostsCommand)
	app.Commands = append(app.Commands, serviesCommand)
	app.Commands = append(app.Commands, rolesCommand)
	app.Commands = append(app.Commands, usersCommand)
	app.Commands = append(app.Commands, saltCommand)
	app.Commands = append(app.Commands, execCommand)

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printTable(title string, headers []string, data [][]string, c *cli.Context) {
	fmt.Println(title)
	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers)
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	} else {
		for i := 1; i <= len(title); i++ {
			fmt.Print("-")
		}
		fmt.Println()
		fmt.Println("NO ENTRIES FOUND!")
	}
}

func printJson(b []byte) {
	fmt.Println(formatJson(b).String())
}

func formatJson(b []byte) *bytes.Buffer {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &out
}

func validateActiveCM(cmServer cm.CMServer) {
	if len(cmServer.Name) == 0 {
		fmt.Println("No active CM server selected. (see 'registry use' command)")
		os.Exit(1)
	} else {
		if cmServer.UseGateway && len(cmServer.ConnectionProfile) == 0 {
			fmt.Println("No connection profile attached to CM server/gateway. (see 'profiles' command)")
			os.Exit(1)
		}
	}
}

func contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
