package cmd

import (
	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// HostsCommand Hosts resource related operations
func HostsCommand() cli.Command {
	return cli.Command{
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
}
