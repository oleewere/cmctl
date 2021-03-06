package cm

// GetFilteredHosts obtain specific hosts based on different filters
func (c CMServer) GetFilteredHosts(filter Filter, inventory *Inventory) map[string]bool {
	finalHosts := make(map[string]bool)
	hosts := make(map[string]bool) // use boolean map as a set
	if filter.Server {
		hosts[c.Hostname] = true
		finalHosts[c.Hostname] = true
	} else if len(filter.Hosts) > 0 {
		cmAgents := c.getCMHosts(inventory)
		calculateAndFillFinalHosts(cmAgents, filter, hosts, finalHosts)
	} else if len(filter.Clusters) > 0 {
		cmAgents := c.getCMHosts(inventory)
		for _, cluster := range filter.Clusters {
			if inventory == nil {
				inventory := Inventory{}
				inventory.EnrichInventoryWithHostsForCluster(cluster, cmAgents)
			}
			for _, host := range inventory.Hosts {
				hosts[host.HostName] = true
			}
			calculateAndFillFinalHosts(cmAgents, filter, hosts, finalHosts)
		}
	} else if len(filter.Roles) > 0 {
		cmAgents := c.getCMHosts(inventory)
		inventories := c.getInventories(inventory, cmAgents)
		for _, inventory := range inventories {
			if len(filter.Clusters) > 0 && !SliceContains(inventory.ClusterName, filter.Clusters) {
				continue
			}
			rolesHostsMap := inventory.ServiceRoleHostsMap[filter.Services[0]]
			for role, rolesHosts := range rolesHostsMap {
				if SliceContains(role, filter.Roles) {
					for _, host := range rolesHosts {
						hosts[host.HostName] = true
					}
				}
			}
		}
		calculateAndFillFinalHosts(cmAgents, filter, hosts, finalHosts)
	} else if len(filter.Services) > 0 {
		cmAgents := c.getCMHosts(inventory)
		inventories := c.getInventories(inventory, cmAgents)
		for _, inventory := range inventories {
			if len(filter.Clusters) > 0 && !SliceContains(inventory.ClusterName, filter.Clusters) {
				continue
			}
			serviceHostsMap := inventory.ServiceHostsMap
			for service, serviceHosts := range serviceHostsMap {
				if SliceContains(service, filter.Services) {
					for _, host := range serviceHosts {
						hosts[host] = true
					}
				}
			}
		}
		calculateAndFillFinalHosts(cmAgents, filter, hosts, finalHosts)
	} else {
		cmAgents := c.getCMHosts(inventory)
		calculateAndFillFinalHosts(cmAgents, filter, hosts, finalHosts)
	}
	return finalHosts
}

func (c CMServer) getCMHosts(inventory *Inventory) []Host {
	if inventory != nil {
		return inventory.Hosts
	}
	return c.ListHosts()
}

func (c CMServer) getInventories(inventory *Inventory, cmAgents []Host) []Inventory {
	inventories := make([]Inventory, 0)
	if inventory != nil {
		inventories = append(inventories, *inventory)
		return inventories
	}
	deployment := c.GetDeployment()
	return CreateInventoriesFromDeploymentsAndHosts(deployment, cmAgents)
}

func calculateAndFillFinalHosts(cmAgents []Host, filter Filter, hosts map[string]bool, finalHosts map[string]bool) {
	for _, cmAgent := range cmAgents {
		if len(filter.Hosts) > 0 {
			filteredHosts := filter.Hosts
			containsHost := false
			for _, filteredHost := range filteredHosts {
				if filteredHost == cmAgent.HostName {
					containsHost = true
				}
			}
			if !containsHost {
				continue
			}
		}
		if len(hosts) > 0 {
			_, ok := hosts[cmAgent.HostName]
			if ok {
				finalHosts[cmAgent.IPAddress] = true
			}
			_, ok = hosts[cmAgent.IPAddress]
			if ok {
				finalHosts[cmAgent.IPAddress] = true
			}

		} else {
			finalHosts[cmAgent.IPAddress] = true
		}
	}
}
