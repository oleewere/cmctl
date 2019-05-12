package cm

import (
	"fmt"
	"os"
	"path"
	"strings"
)

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
	serviceRolesHostsMap := make(map[string]map[string][]HostRoleNamePair)

	for serviceKey, roles := range serviceRoleMap.RolesMap {
		rolesHostsPairMap := make(map[string][]HostRoleNamePair)
		for _, role := range roles {
			if roleHostsPairs, ok := rolesHostsPairMap[role.Type]; ok {
				if !containsRoleHosts(role.HostName, roleHostsPairs) {
					hostRolePair := HostRoleNamePair{HostName: role.HostName, RoleName: role.Name}
					roleHostsPairs = append(roleHostsPairs, hostRolePair)
					rolesHostsPairMap[role.Type] = roleHostsPairs
				}
			} else {
				roleHostsPairs = make([]HostRoleNamePair, 0)
				hostRolePair := HostRoleNamePair{HostName: role.HostName, RoleName: role.Name}
				roleHostsPairs = append(roleHostsPairs, hostRolePair)
				rolesHostsPairMap[role.Type] = roleHostsPairs
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
		serviceRolesHostsMap[serviceKey] = rolesHostsPairMap
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

// CreateInventoryFiles generate ansible compatible inventory files from CM deployment descriptor
func (c CMServer) CreateInventoryFiles(targetDir string, outputFile string, cluster string, serverHostname string) {
	connectionProfileID := c.ConnectionProfile
	if len(connectionProfileID) == 0 {
		fmt.Println("No connection profile is attached for the active CM server entry!")
		os.Exit(1)
	}
	connectionProfile := GetConnectionProfileById(connectionProfileID)

	allHosts := c.ListHosts()
	deployment := c.GetDeployment()
	inventories := CreateInventoriesFromDeploymentsAndHosts(deployment, allHosts)

	for _, inventory := range inventories {
		if len(cluster) > 0 && inventory.ClusterName != cluster {
			fmt.Println(fmt.Sprintf("Skip writing '%v.ini' inventory file. (filter: %v)", inventory.ClusterName, cluster))
		} else {
			var iniFileLocation string
			if len(targetDir) > 0 {
				iniFileLocation = path.Join(targetDir, fmt.Sprintf("%v.ini", inventory.ClusterName))
			} else {
				iniFileLocation = outputFile
			}
			cfg := CreateIniConfiguration()

			cfg.AddSection("server").AddSectionValue(serverHostname)

			agentSection := cfg.AddSection("agent")
			for _, agentHost := range inventory.Hosts {
				agentSection.AddSectionValue(agentHost.HostName)
			}

			for service, serviceHosts := range inventory.ServiceHostsMap {
				serviceSectionName := fmt.Sprintf("service.%v", service)
				serviceSection := cfg.AddSection(serviceSectionName)
				for _, host := range serviceHosts {
					serviceSection.AddSectionValue(host)
				}
			}

			for service, roleHostsMap := range inventory.ServiceRoleHostsMap {
				for roleType, hostRolePairs := range roleHostsMap {
					roleSectionName := fmt.Sprintf("role.%v.%v", service, strings.ToLower(roleType))
					roleSection := cfg.AddSection(roleSectionName)
					for _, hostRolePair := range hostRolePairs {
						roleSection.AddSectionValue(hostRolePair.HostName)
					}
					roleVarsSectionName := fmt.Sprintf("%v:vars", roleSectionName)
					roleVarSection := cfg.AddSection(roleVarsSectionName)
					for _, hostRolePair := range hostRolePairs {
						roleVarSection.AddSectionKeyValue(hostRolePair.RoleName, hostRolePair.HostName)
					}
				}
			}
			cfg.AddSection("cluster").AddSectionKeyValue("name", inventory.ClusterName)

			allVarsSection := cfg.AddSection("all:vars")
			allVarsSection.AddSectionKeyValue("ansible_ssh_user", connectionProfile.Username)
			allVarsSection.AddSectionKeyValue("ansible_ssh_private_key_file", connectionProfile.KeyPath)
			allVarsSection.AddSectionKeyValue("ansible_ssh_common_args", "'-o StrictHostKeyChecking=no'")

			cfg.AddSection("defaults").AddSectionKeyValue("host_key_checking", "False")
			cfg.WriteIniConfiguration(iniFileLocation, []string{"cluster", "server", "agent"})
			fmt.Println(fmt.Sprintf("Inventory file '%v' has been created.", iniFileLocation))
		}
	}
}

// ReadInventoryFromFile parse inventory file and create an inventory resource from it
func ReadInventoryFromFile(inventoryFilePath string) Inventory {
	inventory := Inventory{}
	cfg, err := LoadIniFile(inventoryFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sections := *cfg.Sections
	serviceHostMap := make(map[string][]string)
	serviceRoleHostsMap := make(map[string]map[string][]HostRoleNamePair)
	for _, section := range sections {
		if section.Name == "cluster" {
			inventory.ClusterName = section.KeyValueMap["name"]
		} else if section.Name == "server" {
			serverHosts := *section.Values
			inventory.ServerAddress = serverHosts[0]
		} else if section.Name == "agent" {
			agentHosts := *section.Values
			hosts := make([]Host, 0)
			for _, agentHost := range agentHosts {
				host := Host{HostName: agentHost, IPAddress: agentHost}
				hosts = append(hosts, host)
			}
			inventory.Hosts = hosts
		} else if strings.HasPrefix(section.Name, "service.") {
			serviceName := strings.TrimPrefix(section.Name, "service.")
			serviceHostMap[serviceName] = *section.Values
		} else if strings.HasPrefix(section.Name, "role.") && strings.HasSuffix(section.Name, ":vars") {
			roleParts := strings.Split(section.Name, ".")
			serviceName := roleParts[1]
			withoutVarRoleType := strings.TrimSuffix(roleParts[2], ":vars")
			roleType := strings.ToUpper(withoutVarRoleType)
			roleHostMap := section.KeyValueMap
			var hostRoleNamePairMap map[string][]HostRoleNamePair
			if servicRoleMapVal, ok := serviceRoleHostsMap[serviceName]; ok {
				hostRoleNamePairMap = servicRoleMapVal
			} else {
				hostRoleNamePairMap = make(map[string][]HostRoleNamePair)
			}
			hostRoleNamePairs := make([]HostRoleNamePair, 0)
			for key, value := range roleHostMap {
				hostRoleNamePair := HostRoleNamePair{RoleName: key, HostName: value}
				hostRoleNamePairs = append(hostRoleNamePairs, hostRoleNamePair)
			}
			hostRoleNamePairMap[roleType] = hostRoleNamePairs
			serviceRoleHostsMap[serviceName] = hostRoleNamePairMap
		}
	}
	inventory.ServiceHostsMap = serviceHostMap
	inventory.ServiceRoleHostsMap = serviceRoleHostsMap

	return inventory
}

func containsRoleHosts(host string, roleHostsPairs []HostRoleNamePair) bool {
	contains := false
	for _, hostRolePair := range roleHostsPairs {
		if hostRolePair.HostName == host {
			return true
		}
	}
	return contains
}
