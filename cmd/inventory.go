package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// InventoryCommand Hosts inventory (ansible compatible) related operations
func InventoryCommand() cli.Command {
	return cli.Command{
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
}
