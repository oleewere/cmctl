package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// ClustersCommand Cluster related opterations
func ClustersCommand() cli.Command {
	return cli.Command{
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
}
