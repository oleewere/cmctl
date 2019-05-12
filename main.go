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
	"gopkg.in/yaml.v2"
)

// Version that will be generated during the build as a constant
var Version string

// GitRevString that will be generated during the build as a constant - represents git revision value
var GitRevString string

// DefaultAPIVersion number that is used as api versions in REST api requests e.g. (32): api/v32/..
const DefaultAPIVersion = 32

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

	serversCommand := cli.Command{
		Name:  "servers",
		Usage: "CM server registry DB related operations",
		Subcommands: []cli.Command{
			{
				Name:  "init",
				Usage: "Initialize CM server database",
				Action: func(c *cli.Context) error {
					cm.CreateCMRegistryDb()
					fmt.Println("CM servers registry DB has been initialized.")
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

					apiVersion := cm.GetStringFlag(c.String("verion"), strconv.Itoa(DefaultAPIVersion), "Enter CM API version")
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

					apiVersion := cm.GetStringFlag(c.String("verion"), strconv.Itoa(DefaultAPIVersion), "Enter CM API version")
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
						fmt.Println("Provide a server name argument for use command. e.g.: delete vagrant")
						os.Exit(1)
					}
					name := c.Args().First()
					cmEntryId := cm.GetCMEntryId(name)
					if len(cmEntryId) == 0 {
						fmt.Println("CM server entry does not exist with id " + name)
						os.Exit(1)
					}
					cm.DeRegisterCMEntry(name)
					fmt.Println("CM server de-registered with id: " + name)
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "name", Usage: "name of the CM server entry"},
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
					cli.StringFlag{Name: "name", Usage: "name of the CM server entry"},
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
			{
				Name:  "run",
				Usage: "Run operation on service(s) - start / stop / restart / <custom>",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					clusters := cmServer.ListClusters()
					clustersFilter := c.String("clusters")
					servicesFilter := c.String("services")
					if len(clusters) == 0 {
						fmt.Println(fmt.Sprintf("Not found any clusters for CM server with id '%v'", cmServer.Name))
						os.Exit(1)
					}
					if len(clusters) > 1 && len(clustersFilter) > 0 {
						fmt.Println("Parameter 'clusters' is required!")
						os.Exit(1)
					}
					if len(clusters) == 1 {
						clustersFilter = clusters[0].Name
					}

					if len(servicesFilter) == 0 {
						fmt.Println("Parameter 'services' is required!")
						os.Exit(1)
					}
					filter := cm.CreateFilter(clustersFilter, servicesFilter, "", "", false)

					for _, cluster := range filter.Clusters {
						for _, service := range filter.Services {
							cmServer.RunServiceOperation(cluster, service, c.String("command"), true)
						}
					}

					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "command, c", Usage: "Command: start/stop/restart"},
					cli.StringFlag{Name: "clusters", Usage: "Clusters filter (comma separated)"},
					cli.StringFlag{Name: "services, s", Usage: "Services filter (comma separated)"},
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
				Name:    "attach",
				Aliases: []string{"a"},
				Usage:   "Attach a profile to a CM server entry",
				Action: func(c *cli.Context) error {
					args := c.Args()
					if len(args) == 0 {
						fmt.Println("Provide at least 1 argument (<profile>), or 2 (<profile> and <cmEntry>)")
						os.Exit(1)
					}
					profileID := args.Get(0)
					var cmServer cm.CMServer
					if len(args) == 1 {
						cmServer = cm.GetActiveCM()
						if len(cmServer.Name) == 0 {
							fmt.Println("No active CM selected")
							os.Exit(1)
						}
					} else {
						cmServerID := args.Get(1)
						cm.GetCMById(cmServerID)
						if len(cmServer.Name) == 0 {
							fmt.Println("Cannot find specific CM server entry")
							os.Exit(1)
						}
					}
					profile := cm.GetConnectionProfileById(profileID)
					if len(profile.Name) == 0 {
						fmt.Println("Cannot find specific connection profile entry")
						os.Exit(1)
					}

					cm.SetProfileIdForCMEntry(cmServer.Name, profile.Name)
					msg := fmt.Sprintf("Attach profile '%s' to '%s'", profile.Name, cmServer.Name)
					fmt.Println(msg)
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
					printJSON(clusterTemplate)
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
						tableData = append(tableData, []string{host.HostName, host.IPAddress, host.ClusterName, host.RackID, host.TotalMemory})
					}
					printTable("HOSTS:", []string{"HOSTNAME", "IP", "CLUSTER", "RACK ID", "MEMORY"}, tableData, c)
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
				Usage:   "Print all roles per service / cluster",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					inventoryFile := c.String("inventory")
					clustersFilter := c.String("clusters")

					inventories := make([]cm.Inventory, 0)
					if len(inventoryFile) > 0 {
						inventory := cm.ReadInventoryFromFile(inventoryFile)
						inventories = append(inventories, inventory)
						clustersFilter = inventory.ClusterName
					} else {
						deployment := cmServer.GetDeployment()
						hosts := cmServer.ListHosts()
						inventories = cm.CreateInventoriesFromDeploymentsAndHosts(deployment, hosts)
					}
					filter := cm.CreateFilter(clustersFilter, c.String("services"), "", "", false)

					var counter int
					for _, inventory := range inventories {
						invCluster := inventory.ClusterName
						if len(filter.Clusters) > 0 && !cm.SliceContains(invCluster, filter.Clusters) {
							continue
						}
						serviceRoleHostsMap := inventory.ServiceRoleHostsMap
						for service, roleHostsMap := range serviceRoleHostsMap {
							if len(filter.Services) > 0 && !cm.SliceContains(service, filter.Services) {
								continue
							}
							for roleType, hostRolePairs := range roleHostsMap {
								var tableData [][]string
								for _, hostRolePair := range hostRolePairs {
									tableData = append(tableData, []string{hostRolePair.RoleName, roleType, hostRolePair.HostName, service, invCluster})
								}
								prefixData := ""
								if counter > 0 {
									prefixData = "\n"
								}
								counter++
								header := fmt.Sprintf("%vRoles - %v (cluster: %v):", prefixData, service, invCluster)
								printTable(header, []string{"NAME", "TYPE", "HOSTNAME", "SERVICE", "CLUSTER"}, tableData, c)
							}
						}
					}
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "inventory, i", Usage: "Use hosts inventory file"},
					cli.StringFlag{Name: "clusters, c", Usage: "Cluster filter for roles"},
					cli.StringFlag{Name: "services, s", Usage: "Service filter for roles"},
				},
			},
			{
				Name:  "run",
				Usage: "Run operation on role(s) - start / stop / restart / <custom>",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					clusters := cmServer.ListClusters()
					clustersFilter := c.String("clusters")
					serviceFilter := c.String("service")
					rolesFilter := c.String("roles")
					command := c.String("command")
					inventoryFile := c.String("inventory")

					inventories := make([]cm.Inventory, 0)

					if len(inventoryFile) > 0 {
						inventory := cm.ReadInventoryFromFile(inventoryFile)
						inventories = append(inventories, inventory)
						clustersFilter = inventory.ClusterName
					} else {
						if len(clusters) == 0 {
							fmt.Println(fmt.Sprintf("Not found any clusters for CM server with id '%v'", cmServer.Name))
							os.Exit(1)
						}
						if len(clusters) > 1 && len(clustersFilter) > 0 {
							fmt.Println("Parameter 'clusters' is required!")
							os.Exit(1)
						}
						if len(clusters) == 1 {
							clustersFilter = clusters[0].Name
						}
					}

					if len(serviceFilter) == 0 {
						fmt.Println("Parameter 'service' is required!")
						os.Exit(1)
					}
					if len(rolesFilter) == 0 {
						fmt.Println("Parameter 'roles' is required!")
						os.Exit(1)
					}
					filter := cm.CreateFilter(clustersFilter, serviceFilter, rolesFilter, "", false)
					filter.Roles = cm.UpperAllInSlice(filter.Roles)
					if len(filter.Services) > 1 {
						fmt.Println("Only 1 service can be selected as a filter!")
						os.Exit(1)
					}
					if len(inventoryFile) == 0 {
						deployment := cmServer.GetDeployment()
						hosts := cmServer.ListHosts()
						inventories = cm.CreateInventoriesFromDeploymentsAndHosts(deployment, hosts)
					}

					roleNames := make([]string, 0)
					for _, inventory := range inventories {
						invCluster := inventory.ClusterName
						if !cm.SliceContains(invCluster, filter.Clusters) {
							fmt.Println(fmt.Sprintf("Cluster with name '%v' is filtered out.", invCluster))
							continue
						}
						if roleHostsPairMap, ok := inventory.ServiceRoleHostsMap[serviceFilter]; ok {
							for role, roleNameHostsPairs := range roleHostsPairMap {
								for _, roleNameHostsPair := range roleNameHostsPairs {
									if cm.SliceContains(role, filter.Roles) {
										roleNames = append(roleNames, roleNameHostsPair.RoleName)
									}
								}
							}
						}
					}
					roleNameFilter := c.String("name")
					if len(roleNames) > 0 || len(roleNameFilter) > 0 {
						roleNameFilter := c.String("name")
						cmServer.RunRolesOperation(clustersFilter, serviceFilter, roleNames, roleNameFilter, command, true)
					} else {
						fmt.Println("No roles are selected! (out filtered ?)")
						os.Exit(1)
					}

					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "command, c", Usage: "Command: start/stop/restart"},
					cli.StringFlag{Name: "inventory, i", Usage: "Use hosts inventory file"},
					cli.StringFlag{Name: "clusters", Usage: "Clusters filter (comma separated)"},
					cli.StringFlag{Name: "service, s", Usage: "Services filter"},
					cli.StringFlag{Name: "roles, r", Usage: "Role type filter (comma separated)"},
					cli.StringFlag{Name: "name, n", Usage: "Role name filter"},
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
				Usage:   "Print all users per cluster",
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

	configsCommand := cli.Command{
		Name:  "configs",
		Usage: "Config group related opreations",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all config groups per role / service / cluster",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)

					typeFilter := c.String("type")
					if len(typeFilter) == 0 || !(typeFilter == "service" || typeFilter == "role") {
						fmt.Println("Filter '--type' is required for this operation ('service' or 'role')")
					}

					clusters := cmServer.ListClusters()
					clustersFilter := c.String("clusters")

					if len(clusters) == 0 {
						fmt.Println(fmt.Sprintf("Not found any clusters for CM server with id '%v'", cmServer.Name))
						os.Exit(1)
					}
					if len(clusters) > 1 && len(clustersFilter) > 0 {
						fmt.Println("Parameter 'clusters' is required!")
						os.Exit(1)
					}
					if len(clusters) == 1 {
						clustersFilter = clusters[0].Name
					}

					serviceFilter := c.String("services")
					if len(serviceFilter) == 0 {
						fmt.Println("Parameter 'services' is required!")
						os.Exit(1)
					}
					fullView := c.Bool("full")

					if typeFilter == "service" {
						filter := cm.CreateFilter(clustersFilter, serviceFilter, "", "", false)
						var counter int
						for _, cluster := range filter.Clusters {
							for _, service := range filter.Services {
								serviceConfigs := cmServer.ListServiceConfigs(cluster, service, fullView)
								var tableData [][]string
								for _, configItem := range serviceConfigs {
									sensitiveStr := "false"
									if configItem.Sensitive {
										sensitiveStr = "true"
									}
									if fullView {
										requiredStr := "false"
										if configItem.Required {
											requiredStr = "true"
										}
										tableData = append(tableData, []string{configItem.Name, configItem.Value, sensitiveStr, configItem.Default, requiredStr})
									} else {
										tableData = append(tableData, []string{configItem.Name, configItem.Value, sensitiveStr})
									}
								}
								prefixData := ""
								if counter > 0 {
									prefixData = "\n"
								}
								counter++
								if fullView {
									header := fmt.Sprintf("%vConfigs - %v (cluster: %v):", prefixData, service, cluster)
									printTable(header, []string{"NAME", "VALUE", "SENSITIVE", "DEFAULT", "REQUIRED"}, tableData, c)
								} else {
									header := fmt.Sprintf("%vConfigs - %v (cluster: %v):", prefixData, service, cluster)
									printTable(header, []string{"NAME", "VALUE", "SENSITIVE"}, tableData, c)
								}
							}
						}
					}
					if typeFilter == "role" {
						rolesFilter := c.String("roles")
						filter := cm.CreateFilter(clustersFilter, serviceFilter, rolesFilter, "", false)
						if len(filter.Roles) > 0 && len(filter.Services) > 1 {
							fmt.Println("Use exactly 1 service for '--services' with roles filter!")
							os.Exit(1)
						}
						var counter int
						for _, cluster := range filter.Clusters {
							for _, service := range filter.Services {
								roleConfigGroups := cmServer.ListRoleConfigGroups(cluster, service, fullView)
								for _, roleConfigGroup := range roleConfigGroups {
									if len(filter.Roles) > 0 && !cm.SliceContains(roleConfigGroup.RoleType, filter.Roles) {
										continue
									}
									var tableData [][]string
									for _, configItem := range roleConfigGroup.ConfigItems {
										sensitiveStr := "false"
										if configItem.Sensitive {
											sensitiveStr = "true"
										}
										if fullView {
											requiredStr := "false"
											if configItem.Required {
												requiredStr = "true"
											}
											tableData = append(tableData, []string{configItem.Name, configItem.Value, sensitiveStr,
												configItem.Default, requiredStr})
										} else {
											tableData = append(tableData, []string{configItem.Name, configItem.Value, sensitiveStr,
												service, cluster})
										}
									}
									prefixData := ""
									if counter > 0 {
										prefixData = "\n"
									}
									counter++
									if fullView {
										header := fmt.Sprintf("%vRole Config Group - %v (role type: %v, service: %v, cluster:%v):", prefixData,
											roleConfigGroup.Name, roleConfigGroup.RoleType, service, cluster)
										printTable(header, []string{"NAME", "VALUE", "SENSITIVE", "DEFAULT", "REQUIRED"}, tableData, c)
									} else {
										header := fmt.Sprintf("%vRole Config Group - %v (role type: %v):", prefixData,
											roleConfigGroup.Name, roleConfigGroup.RoleType)
										printTable(header, []string{"NAME", "VALUE", "SENSITIVE", "SERVICE", "CLUSTER"}, tableData, c)
									}
								}
							}
						}
					}

					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "type, t", Usage: "Config type (service or role)"},
					cli.StringFlag{Name: "clusters, c", Usage: "Clusters filter (comma separated)"},
					cli.StringFlag{Name: "services, s", Usage: "Services filter (comma separated)"},
					cli.StringFlag{Name: "roles, r", Usage: "Role type filter (comma separated)"},
					cli.BoolFlag{Name: "full", Usage: "Show all configurations"},
				},
			},
			{
				Name:  "set",
				Usage: "Update service or role config",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)

					typeFilter := c.String("type")
					if len(typeFilter) == 0 || !(typeFilter == "service" || typeFilter == "role") {
						fmt.Println("Filter '--type' is required for this operation ('service' or 'role')")
					}
					configKeyFilter := c.String("key")
					configValueFilter := c.String("value")
					clearConfig := c.Bool("clear")

					if len(configKeyFilter) == 0 {
						fmt.Println("Parameter 'key' is required!")
						os.Exit(1)
					}
					if len(configValueFilter) == 0 && !clearConfig {
						fmt.Println("Parameter 'value' or 'clear' are required!")
						os.Exit(1)
					}

					clusters := cmServer.ListClusters()
					clusterFilter := c.String("cluster")

					if len(clusters) == 0 {
						fmt.Println(fmt.Sprintf("Not found any clusters for CM server with id '%v'", cmServer.Name))
						os.Exit(1)
					}

					if len(clusters) == 0 {
						fmt.Println(fmt.Sprintf("Not found any clusters for CM server with id '%v'", cmServer.Name))
						os.Exit(1)
					}
					if len(clusters) > 1 && len(clusterFilter) > 0 {
						fmt.Println("Parameter 'cluster' is required!")
						os.Exit(1)
					}
					if len(clusters) == 1 {
						clusterFilter = clusters[0].Name
					}

					serviceFilter := c.String("service")
					if len(serviceFilter) == 0 {
						fmt.Println("Parameter 'service' is required!")
						os.Exit(1)
					}
					keyValueMap := make(map[string]string)
					keyValueMap[configKeyFilter] = configValueFilter
					if typeFilter == "service" {
						cmServer.UpdateServiceConfigs(clusterFilter, serviceFilter, keyValueMap, clearConfig, true)
					}
					if typeFilter == "role" {
						roleTypeFilter := c.String("role")
						roleConfigGroupFilter := c.String("name")

						if len(roleTypeFilter) == 0 && len(roleConfigGroupFilter) == 0 {
							fmt.Println("Parameter 'role' or 'name' is required!")
							os.Exit(1)
						}
						if len(roleTypeFilter) > 0 {
							roleTypeFilter = strings.ToUpper(roleTypeFilter)
						}
						roleConfigGroups := cmServer.ListRoleConfigGroups(clusterFilter, serviceFilter, true)
						for _, roleConfigGroup := range roleConfigGroups {
							if len(roleConfigGroupFilter) > 0 && roleConfigGroup.Name != roleConfigGroupFilter {
								continue
							}
							if len(roleTypeFilter) > 0 && roleConfigGroup.RoleType != roleTypeFilter {
								continue
							}
							cmServer.UpdateRoleConfigs(clusterFilter, serviceFilter,
								roleConfigGroup.RoleType, roleConfigGroup.Name, keyValueMap, clearConfig, true)
						}
					}

					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "type, t", Usage: "Config type (service or role)"},
					cli.StringFlag{Name: "cluster, c", Usage: "Cluster filter"},
					cli.StringFlag{Name: "service, s", Usage: "Service filter"},
					cli.StringFlag{Name: "role, r ", Usage: "Role type filter"},
					cli.StringFlag{Name: "name, n", Usage: "Role config group name filter"},
					cli.StringFlag{Name: "key, k", Usage: "Configuration key"},
					cli.StringFlag{Name: "value, v", Usage: "Configuration value"},
					cli.BoolFlag{Name: "clear", Usage: "Clean configuration value (set to default)"},
				},
			},
		},
	}

	inventoryCommand := cli.Command{
		Name:  "inventory",
		Usage: "Hosts inventory (ansible compatible) related operations",
		Subcommands: []cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Generate host inventory files",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					outputDir := c.String("directory")
					outputFile := c.String("output")
					clusterFilter := c.String("cluster")
					serverHostname := c.String("server-hostname")
					if len(serverHostname) == 0 {
						serverHostname = cmServer.Hostname
					}
					if len(outputDir) == 0 && len(outputFile) == 0 {
						currentDir, err := os.Getwd()
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						outputDir = currentDir
						fmt.Println("Output file or directory is not provided. Inventory files will be generated in the current folder.")
					}
					if len(outputDir) > 0 && !cm.Exists(outputDir) {
						fmt.Println(fmt.Sprintf("Output directory does not exist: '%v'", outputDir))
						os.Exit(1)
					}
					clusters := cmServer.ListClusters()
					if len(outputFile) > 0 && len(clusterFilter) == 0 && len(clusters) > 1 {
						fmt.Println("Cluster filter is required! (multiple clusters found)")
					}
					if len(clusters) == 1 && len(clusterFilter) == 0 {
						clusterFilter = clusters[0].Name
					}
					cmServer.CreateInventoryFiles(outputDir, outputFile, clusterFilter, serverHostname)
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "directory, d", Usage: "Output directory for inventory files"},
					cli.StringFlag{Name: "output, o", Usage: "Output file for 1 inventory file"},
					cli.StringFlag{Name: "cluster, c", Usage: "Cluster filter (required for 'output' option)"},
					cli.StringFlag{Name: "server-hostname, s", Usage: "Override CM server hostname"},
				},
			},
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "Print inventory file mappings in JSON format",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					iniFile := c.String("inventory")
					inventory := cm.ReadInventoryFromFile(iniFile)
					outputFormat := "json"
					outputFormatInput := c.String("output-format")
					if len(outputFormatInput) > 0 {
						outputFormat = outputFormatInput
					}
					if outputFormat == "json" {
						bodyBytes, err := json.Marshal(inventory)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						printJSON(bodyBytes)
					} else if outputFormat == "yaml" || outputFormat == "yml" {
						bodyBytes, err := yaml.Marshal(inventory)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						fmt.Println(string(bodyBytes))
					} else {
						fmt.Println(fmt.Sprintf("Unsupported output format: %v", outputFormat))
						os.Exit(1)
					}
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "inventory, i", Usage: "Inventory file"},
					cli.StringFlag{Name: "output-format, o", Usage: "Output type (json or yaml)"},
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
						fmt.Println("Command parameter is missing! (use 'command' or 'c')")
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
					if len(c.String("source")) == 0 {
						fmt.Println("The 'source' parameter is required for sync salt scripts!")
						os.Exit(1)
					}
					cmServer.SyncSaltScripts(c.String("source"), c.String("target"), c.String("filter"), c.Bool("clear"))
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "source, s", Usage: "Source path for the local salt script folder to sync"},
					cli.StringFlag{Name: "target, t", Usage: "Target remote path for the salt scripts (default: /srv)"},
					cli.StringFlag{Name: "filter, f", Usage: "Filter to use only second level folders for uploading (e.g.: pillar)"},
					cli.BoolFlag{Name: "clear", Usage: "Delete salt backup(s) from the gateway"},
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
			command := c.String("command")
			if len(command) == 0 {
				fmt.Println("Command parameter is missing! (use 'command' or 'c')")
				os.Exit(1)
			}
			filter := cm.CreateFilter(c.String("clusters"), c.String("services"), c.String("roles"), c.String("hosts"), c.Bool("server"))
			filter.Roles = cm.UpperAllInSlice(filter.Roles)
			hosts := cmServer.GetFilteredHosts(filter)
			cmServer.RunRemoteHostCommand(command, hosts, filter.Server)
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "command, c", Usage: "Command to execute on the remote hosts"},
			cli.BoolFlag{Name: "server", Usage: "Filter on CM server"},
			cli.StringFlag{Name: "clusters", Usage: "Filter on clusters (comma separated)"},
			cli.StringFlag{Name: "services", Usage: "Filter on services (comma separated)"},
			cli.StringFlag{Name: "roles", Usage: "Filter on role types (comma separated)"},
			cli.StringFlag{Name: "hosts", Usage: "Filter on hosts (comma separated)"},
		},
	}

	app.Commands = append(app.Commands, serversCommand)
	app.Commands = append(app.Commands, profileCommand)
	app.Commands = append(app.Commands, clustersCommand)
	app.Commands = append(app.Commands, hostsCommand)
	app.Commands = append(app.Commands, serviesCommand)
	app.Commands = append(app.Commands, rolesCommand)
	app.Commands = append(app.Commands, usersCommand)
	app.Commands = append(app.Commands, configsCommand)
	app.Commands = append(app.Commands, inventoryCommand)
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

func printJSON(b []byte) {
	fmt.Println(formatJSON(b).String())
}

func formatJSON(b []byte) *bytes.Buffer {
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
		fmt.Println("No active CM server selected. (see 'servers use' command)")
		os.Exit(1)
	} else {
		if cmServer.UseGateway && len(cmServer.ConnectionProfile) == 0 {
			fmt.Println("No connection profile attached to CM server/gateway. (see 'profiles' command)")
			os.Exit(1)
		}
	}
}
