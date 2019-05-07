package cm

// CreateInventoriesFromDeploymentsAndHosts create struct that holds hostnames by services/roles for all clusters
func CreateInventoriesFromDeploymentsAndHosts(deployment Deployment, agentHosts []Host) []Inventory {
	inventories := make([]Inventory, 0)
	clusterServiceRolesMap := deployment.ClusterServiceRoleMap
	for clusterKey, serviceRoleMapObj := range clusterServiceRolesMap {
		inventory := CreateInventory(clusterKey, serviceRoleMapObj, deployment, agentHosts)
		inventories = append(inventories, inventory)
	}
	return inventories
}

// CreateInventory create struct that holds hostnames by services/roles/cluster
func CreateInventory(clusterName string, serviceRoleMap ServiceRolesMap, deployment Deployment, agentHosts []Host) Inventory {
	inventory := Inventory{}
	serviceHostsMap := make(map[string][]string)
	serviceRolesHostsMap := make(map[string]map[string][]string)

	for serviceKey, roles := range serviceRoleMap.RolesMap {
		rolesHostsMap := make(map[string][]string)
		for _, role := range roles {
			if roleHosts, ok := rolesHostsMap[role.Type]; ok {
				if !SliceContains(role.HostName, roleHosts) {
					roleHosts = append(roleHosts, role.HostName)
					rolesHostsMap[role.Type] = roleHosts
				}
			} else {
				roleHosts = make([]string, 0)
				roleHosts = append(roleHosts, role.HostName)
				rolesHostsMap[role.Type] = roleHosts
			}
			if serviceHosts, ok := serviceHostsMap[serviceKey]; ok {
				if !SliceContains(role.HostName, serviceHosts) {
					serviceHosts = append(serviceHosts, role.HostName)
					serviceHostsMap[serviceKey] = serviceHosts
				}
			} else {
				serviceHosts = make([]string, 0)
				serviceHosts = append(serviceHosts, role.HostName)
				serviceHostsMap[serviceKey] = serviceHosts
			}
		}
		serviceRolesHostsMap[serviceKey] = rolesHostsMap
	}

	inventory.ClusterName = clusterName
	inventory = GetInventoryWithHostsForCluster(inventory, clusterName, agentHosts)
	inventory.ServiceRoleHostsMap = serviceRolesHostsMap
	inventory.ServiceHostsMap = serviceHostsMap

	return inventory
}

// GetInventoryWithHostsForCluster enrich inventory struct with hosts for a specific cluster
func GetInventoryWithHostsForCluster(inventory Inventory, cluster string, agentHosts []Host) Inventory {
	clusterHosts := make([]Host, 0)
	for _, host := range agentHosts {
		if host.ClusterName == cluster {
			clusterHosts = append(clusterHosts, host)
		}
	}
	inventory.Hosts = clusterHosts
	return inventory
}
