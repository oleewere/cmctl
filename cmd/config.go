package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// ConfigsCommand Config group related opreations
func ConfigsCommand() cli.Command {
	return cli.Command{
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
}
