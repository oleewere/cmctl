package cmd

import (
	"fmt"
	"os"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// ServicesCommand service related operations
func ServicesCommand() cli.Command {
	return cli.Command{
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
}
