package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// DefaultAPIVersion number that is used as api versions in REST api requests e.g. (31): api/v31/..
const DefaultAPIVersion = 31

// ServersCommand CM server registry DB related operations
func ServersCommand() cli.Command {
	return cli.Command{
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
					cmEntryID := cm.GetCMEntryID(name)
					if len(cmEntryID) > 0 {
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
					existingCmServer := cm.GetCMByID(cmServerName)
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
					cmEntryID := cm.GetCMEntryID(name)
					if len(cmEntryID) == 0 {
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
					cmEntryID := cm.GetCMEntryID(name)
					if len(cmEntryID) == 0 {
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
}
