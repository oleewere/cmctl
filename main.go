package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	initCommand := cli.Command{
		Name:  "init",
		Usage: "Initialize CM server database",
		Action: func(c *cli.Context) error {
			cm.CreateCMRegistryDb()
			fmt.Println("CM registry DB has been initialized.")
			return nil
		},
	}

	createCommand := cli.Command{
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
	}

	updateCommand := cli.Command{
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
	}

	deleteCommand := cli.Command{
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
	}

	useCommand := cli.Command{
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
	}

	clearCommand := cli.Command{
		Name:  "clear",
		Usage: "Drop all CM server records",
		Action: func(c *cli.Context) error {
			cm.DropCMRegistryRecords()
			fmt.Println("CM server entries dropped.")
			return nil
		},
	}

	showCommand := cli.Command{
		Name:  "show",
		Usage: "Show active CM server details",
		Action: func(c *cli.Context) error {
			cmServer := cm.GetActiveCM()
			var tableData [][]string
			if len(cmServer.Name) > 0 {
				tableData = append(tableData, []string{cmServer.Name, cmServer.Hostname, strconv.Itoa(cmServer.Port), cmServer.Protocol,
					cmServer.Username, "********", cmServer.ConnectionProfile, "true"})
			}
			printTable("ACTIVE CM SERVER:", []string{"Name", "HOSTNAME", "PORT", "PROTOCOL", "USER", "PASSWORD", "PROFILE", "ACTIVE"}, tableData, c)
			return nil
		},
	}

	listCommand := cli.Command{
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
	}

	serviesCommand := cli.Command{
		Name:  "services",
		Usage: "Service related commands",
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
							tableData = append(tableData, []string{service.DisplayName, service.State, cluster.DisplayName, service.StaleConfig})
						}
						prefixData := ""
						if index > 0 {
							prefixData = "\n"
						}
						header := fmt.Sprintf("%vSERVICES - %v (%v):", prefixData, cluster.DisplayName, cluster.Name)
						printTable(header, []string{"NAME", "STATE", "CLUSTER", "CONFIG"}, tableData, c)
					}
					return nil
				},
			},
		},
	}

	profileCommand := cli.Command{
		Name:  "profiles",
		Usage: "Connection profiles related commands",
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
		Usage: "Attach a profile to an CM server entry",
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

	listClustersCommand := cli.Command{
		Name:  "clusters",
		Usage: "Print all registered CM clusters",
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
	}

	listHostsCommand := cli.Command{
		Name:  "hosts",
		Usage: "Print all registered CM hosts",
		Action: func(c *cli.Context) error {
			cmServer := cm.GetActiveCM()
			validateActiveCM(cmServer)
			hosts := cmServer.ListHosts()
			var tableData [][]string
			for _, host := range hosts {
				tableData = append(tableData, []string{host.HostName, host.IPAddress, host.ClusterDisplayName, host.CommissionState, host.RackID})
			}
			printTable("HOSTS:", []string{"HOSTNAME", "IP", "CLUSTER", "STATE", "RACK ID"}, tableData, c)
			return nil
		},
	}

	app.Commands = append(app.Commands, initCommand)
	app.Commands = append(app.Commands, createCommand)
	app.Commands = append(app.Commands, listCommand)
	app.Commands = append(app.Commands, updateCommand)
	app.Commands = append(app.Commands, deleteCommand)
	app.Commands = append(app.Commands, clearCommand)
	app.Commands = append(app.Commands, useCommand)
	app.Commands = append(app.Commands, showCommand)
	app.Commands = append(app.Commands, profileCommand)
	app.Commands = append(app.Commands, attachCommand)
	app.Commands = append(app.Commands, listClustersCommand)
	app.Commands = append(app.Commands, listHostsCommand)
	app.Commands = append(app.Commands, serviesCommand)

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
		fmt.Println("No active CM server selected. (see 'use' command)")
		os.Exit(1)
	} else {
		if cmServer.UseGateway && len(cmServer.ConnectionProfile) == 0 {
			fmt.Println("No connection profile attached to CM server/gateway. (see 'profiles' command)")
			os.Exit(1)
		}
	}
}
