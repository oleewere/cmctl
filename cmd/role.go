package cmd

import (
	"fmt"
	"os"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// RolesCommand Role related opreations
func RolesCommand() cli.Command {
	return cli.Command{
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
							var tableData [][]string
							for roleType, hostRolePairs := range roleHostsMap {
								for _, hostRolePair := range hostRolePairs {
									tableData = append(tableData, []string{hostRolePair.RoleName, roleType, hostRolePair.HostName, service, invCluster})
								}
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
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "inventory, i", Usage: "Hosts inventory file"},
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
					cli.StringFlag{Name: "inventory, i", Usage: "Hosts inventory file"},
					cli.StringFlag{Name: "clusters", Usage: "Clusters filter (comma separated)"},
					cli.StringFlag{Name: "service, s", Usage: "Services filter"},
					cli.StringFlag{Name: "roles, r", Usage: "Role type filter (comma separated)"},
					cli.StringFlag{Name: "name, n", Usage: "Role name filter"},
				},
			},
		},
	}
}
