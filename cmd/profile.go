package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// ProfilesCommand Connection profiles related opreations
func ProfilesCommand() cli.Command {
	return cli.Command{
		Name:  "profiles",
		Usage: "Connection profiles related opreations",
		Subcommands: []cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create new connection profile",
				Action: func(c *cli.Context) error {
					name := cm.GetStringFlag(c.String("name"), "", "Enter connection profile name")
					connProfileID := cm.GetConnectionProfileEntryID(name)
					if len(connProfileID) > 0 {
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
						cm.GetCMByID(cmServerID)
						if len(cmServer.Name) == 0 {
							fmt.Println("Cannot find specific CM server entry")
							os.Exit(1)
						}
					}
					profile := cm.GetConnectionProfileByID(profileID)
					if len(profile.Name) == 0 {
						fmt.Println("Cannot find specific connection profile entry")
						os.Exit(1)
					}

					cm.SetProfileIDForCMEntry(cmServer.Name, profile.Name)
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
					profileEntryID := cm.GetConnectionProfileEntryID(name)
					if len(profileEntryID) == 0 {
						fmt.Println("Connection profile entry does not exist with id " + name)
						os.Exit(1)
					}
					cm.DeRegisterConnectionProfile(profileEntryID)
					msg := fmt.Sprintf("Connection profile '%s' has been deleted successfully", profileEntryID)
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
}
