package cmd

import (
	"fmt"
	"os"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// ExecCommand Execute commands on all (or specific) hosts
func ExecCommand() cli.Command {
	return cli.Command{
		Name:  "exec",
		Usage: "Execute commands on all (or specific) hosts",
		Action: func(c *cli.Context) error {
			cmServer := cm.GetActiveCM()
			validateActiveCM(cmServer)
			command := c.String("command")
			inventoryFile := c.String("inventory")
			if len(command) == 0 {
				fmt.Println("Command parameter is missing! (use 'command' or 'c')")
				os.Exit(1)
			}
			var inventory *cm.Inventory
			if len(inventoryFile) > 0 {
				inventoryVal := cm.ReadInventoryFromFile(inventoryFile)
				inventory = &inventoryVal
			}
			filter := cm.CreateFilter(c.String("clusters"), c.String("services"), c.String("roles"), c.String("hosts"), c.Bool("server"))
			filter.Roles = cm.UpperAllInSlice(filter.Roles)
			hosts := cmServer.GetFilteredHosts(filter, inventory)
			cmServer.RunRemoteHostCommand(command, hosts, filter.Server, inventory)
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "command, c", Usage: "Command to execute on the remote hosts"},
			cli.StringFlag{Name: "inventory, i", Usage: "Hosts inventory file"},
			cli.BoolFlag{Name: "server", Usage: "Filter on CM server"},
			cli.StringFlag{Name: "clusters", Usage: "Filter on clusters (comma separated)"},
			cli.StringFlag{Name: "services", Usage: "Filter on services (comma separated)"},
			cli.StringFlag{Name: "roles", Usage: "Filter on role types (comma separated)"},
			cli.StringFlag{Name: "hosts", Usage: "Filter on hosts (comma separated)"},
		},
	}
}
