package cm

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-ini/ini"
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
			cfg := ini.Empty()

			cfg.NewRawSection("server", serverHostname+"\n")

			agentRawBody := ""
			for _, agentHost := range inventory.Hosts {
				agentRawBody += (agentHost.HostName + "\n")
			}
			cfg.NewRawSection("agent", agentRawBody)

			for service, serviceHosts := range inventory.ServiceHostsMap {
				serviceSection := fmt.Sprintf("service.%v", service)
				serviceRawBody := ""
				for _, host := range serviceHosts {
					serviceRawBody += (host + "\n")
				}
				cfg.NewRawSection(serviceSection, serviceRawBody)
			}

			for service, roleHostsMap := range inventory.ServiceRoleHostsMap {
				for roleType, hostRolePairs := range roleHostsMap {
					roleSection := fmt.Sprintf("role.%v.%v", service, strings.ToLower(roleType))
					roleRawBody := ""
					for _, hostRolePair := range hostRolePairs {
						roleRawBody += (hostRolePair.HostName + "\n")
					}
					cfg.NewRawSection(roleSection, roleRawBody)
					roleVarsSection := fmt.Sprintf("%v:vars", roleSection)
					roleVarsRawBody := ""
					for _, hostRolePair := range hostRolePairs {
						roleNameHostPair := fmt.Sprintf("%v=%v\n", hostRolePair.RoleName, hostRolePair.HostName)
						roleVarsRawBody += roleNameHostPair
					}
					cfg.NewRawSection(roleVarsSection, roleVarsRawBody)
				}
			}

			cfg.NewRawSection("cluster", fmt.Sprintf("name=%v\n", inventory.ClusterName))

			allVarsRawBody := ""
			allVarsRawBody += (fmt.Sprintf("%v=%v", "ansible_ssh_user", connectionProfile.Username) + "\n")
			allVarsRawBody += (fmt.Sprintf("%v=%v", "ansible_ssh_private_key_file", connectionProfile.KeyPath) + "\n")
			allVarsRawBody += (fmt.Sprintf("%v=%v", "ansible_ssh_common_args", "'-o StrictHostKeyChecking=no'") + "\n")
			cfg.NewRawSection("all:vars", allVarsRawBody)

			cfg.NewRawSection("defaults", "host_key_checking=False\n")
			cfg.SaveTo(iniFileLocation)
			fmt.Println(fmt.Sprintf("Inventory file '%v' has been created.", iniFileLocation))
		}
	}
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
