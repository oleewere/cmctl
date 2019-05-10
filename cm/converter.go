package cm

// ConvertHostsResponse convert items response to Hosts response
func (c Items) ConvertHostsResponse() []Host {
	hosts := []Host{}
	for _, item := range c.Items {
		hosts = createHostsType(item, hosts)
	}
	return hosts
}

// ConvertClustersResponse convert items response to Clusters response
func (c Items) ConvertClustersResponse() []Cluster {
	clusters := []Cluster{}
	for _, item := range c.Items {
		clusters = createClustersType(item, clusters)
	}
	return clusters
}

// ConvertDeploymentResponse convert map to Deployment response
func ConvertDeploymentResponse(responseMap map[string]interface{}) Deployment {
	deployment := Deployment{}
	clusters := []Cluster{}

	clusterServiceRoleMap := make(map[string]ServiceRolesMap)
	deployment.ClusterServiceRoleMap = clusterServiceRoleMap

	if clustersVal, ok := responseMap["clusters"]; ok {
		clustersI := clustersVal.([]interface{})
		for _, clusterItemVal := range clustersI {
			clusterItem := clusterItemVal.(map[string]interface{})
			clusters = createClustersType(clusterItem, clusters)
			deployment = fillServicesAndRolesFromCluster(clusterItem, deployment)
		}
	}
	deployment.Clusters = clusters
	return deployment
}

// ConvertServicesResponse convert items response to Services response
func (c Items) ConvertServicesResponse(cluster string) []Service {
	services := []Service{}
	for _, item := range c.Items {
		services = createServiceType(item, services, cluster)
	}
	return services
}

// ConvertRolesResponse convert items response to Roles response
func (c Items) ConvertRolesResponse(cluster string, service string) []Role {
	roles := []Role{}
	for _, item := range c.Items {
		roles = createRoleType(item, roles, cluster, service)
	}
	return roles
}

// ConvertUsersResponse convert items response to Users response
func (c Items) ConvertUsersResponse() []User {
	users := []User{}
	for _, item := range c.Items {
		users = createUserType(item, users)
	}
	return users
}

// ConvertServiceConfigResponse convert items response to config items
func (c Items) ConvertServiceConfigResponse() []ConfigItem {
	configs := []ConfigItem{}
	for _, item := range c.Items {
		configs = createConfigType(item, configs)
	}
	return configs
}

// ConvertRoleConfigGroupsResponse convert items to role configuration groups
func (c Items) ConvertRoleConfigGroupsResponse() []RoleConfigGroup {
	roleConfigGroups := []RoleConfigGroup{}
	for _, item := range c.Items {
		roleConfigGroups = createRoleConfigGroupType(item, roleConfigGroups)
	}
	return roleConfigGroups
}

func createRoleConfigGroupType(item Item, roleConfigGroups []RoleConfigGroup) []RoleConfigGroup {
	roleConfigGroup := RoleConfigGroup{}
	if name, ok := item["name"]; ok {
		roleConfigGroup.Name = name.(string)
	}
	if displayName, ok := item["displayName"]; ok {
		roleConfigGroup.DisplayName = displayName.(string)
	}
	if roleType, ok := item["roleType"]; ok {
		roleConfigGroup.RoleType = roleType.(string)
	}
	if base, ok := item["base"]; ok {
		roleConfigGroup.Base = base.(bool)
	}

	if serviceRefVal, ok := item["serviceRef"]; ok {
		serviceRef := serviceRefVal.(map[string]interface{})
		if clusterName, ok := serviceRef["clusterName"]; ok {
			roleConfigGroup.ClusterName = clusterName.(string)
		}
		if serviceName, ok := serviceRef["serviceName"]; ok {
			roleConfigGroup.ServiceName = serviceName.(string)
		}
	}
	configItems := make([]ConfigItem, 0)
	if configVal, ok := item["config"]; ok {
		config := configVal.(map[string]interface{})
		if itemsVal, ok := config["items"]; ok {
			items := itemsVal.([]interface{})
			for _, item := range items {
				configItem := item.(map[string]interface{})
				configItems = createConfigType(configItem, configItems)
			}
		}
	}
	roleConfigGroup.ConfigItems = configItems

	roleConfigGroups = append(roleConfigGroups, roleConfigGroup)
	return roleConfigGroups
}

func createConfigType(item Item, configs []ConfigItem) []ConfigItem {
	configItem := ConfigItem{}

	if name, ok := item["name"]; ok {
		configItem.Name = name.(string)
	}
	if value, ok := item["value"]; ok {
		configItem.Value = value.(string)
	}
	if displayName, ok := item["displayName"]; ok {
		configItem.DisplayName = displayName.(string)
	}
	if description, ok := item["description"]; ok {
		configItem.Description = description.(string)
	}
	if defaultVal, ok := item["default"]; ok {
		configItem.Default = defaultVal.(string)
	}
	if sensitive, ok := item["sensitive"]; ok {
		configItem.Sensitive = sensitive.(bool)
	}
	if required, ok := item["required"]; ok {
		configItem.Required = required.(bool)
	}
	if relatedName, ok := item["relatedName"]; ok {
		configItem.RelatedName = relatedName.(string)
	}
	if validationWarningsSuppressed, ok := item["validationWarningsSuppressed"]; ok {
		configItem.ValidationWarningsSuppressed = validationWarningsSuppressed.(bool)
	}
	if validationState, ok := item["validationState"]; ok {
		configItem.ValidationState = validationState.(string)
	}

	configs = append(configs, configItem)
	return configs
}

func createUserType(item Item, users []User) []User {
	user := User{}

	if name, ok := item["name"]; ok {
		user.Name = name.(string)
	}

	authRoles := make([]UserAuthRole, 0)

	if authRolesVal, ok := item["authRoles"]; ok {
		authRolesI := authRolesVal.([]interface{})
		for _, authRoleItem := range authRolesI {
			authRoleMap := authRoleItem.(map[string]interface{})
			if authRoleName, ok := authRoleMap["name"]; ok {
				authRole := UserAuthRole{Name: authRoleName.(string)}
				authRoles = append(authRoles, authRole)
			}
		}
	}
	user.AuthRoles = authRoles
	users = append(users, user)
	return users
}

func createRoleType(item Item, roles []Role, cluster string, service string) []Role {
	role := Role{}

	if name, ok := item["name"]; ok {
		role.Name = name.(string)
	}
	if typeVal, ok := item["type"]; ok {
		role.Type = typeVal.(string)
	}
	if state, ok := item["roleState"]; ok {
		role.State = state.(string)
	}
	if staleVal, ok := item["configStalenessStatus"]; ok {
		role.StaleConfig = staleVal.(string)
	}
	if commissionState, ok := item["commissionState"]; ok {
		role.ComissionState = commissionState.(string)
	}
	if roleConfigGroupRefVal, ok := item["roleConfigGroupRef"]; ok {
		roleConfigGroupRef := roleConfigGroupRefVal.(map[string]interface{})
		if roleConfigGroupName, ok := roleConfigGroupRef["roleConfigGroupName"]; ok {
			role.ConfigGroup = roleConfigGroupName.(string)
		}
	}
	if hostRefVal, ok := item["hostRef"]; ok {
		hostRef := hostRefVal.(map[string]interface{})
		if hostname, ok := hostRef["hostname"]; ok {
			role.HostName = hostname.(string)
		}
	}
	role.ClusterName = cluster
	role.ServiceName = service

	roles = append(roles, role)
	return roles
}

func createServiceType(item Item, services []Service, cluster string) []Service {
	service := Service{}
	if name, ok := item["name"]; ok {
		service.Name = name.(string)
	}
	if displayName, ok := item["displayName"]; ok {
		service.DisplayName = displayName.(string)
	}
	if serviceType, ok := item["type"]; ok {
		service.Type = serviceType.(string)
	}
	if state, ok := item["serviceState"]; ok {
		service.State = state.(string)
	}
	if staleConfig, ok := item["configStalenessStatus"]; ok {
		service.StaleConfig = staleConfig.(string)
	}
	service.ClusterName = cluster
	services = append(services, service)
	return services
}

func createClustersType(item Item, clusters []Cluster) []Cluster {
	cluster := Cluster{}
	if name, ok := item["name"]; ok {
		cluster.Name = name.(string)
	}
	if displayName, ok := item["displayName"]; ok {
		cluster.DisplayName = displayName.(string)
	}
	if fullVersion, ok := item["fullVersion"]; ok {
		cluster.Version = fullVersion.(string)
	}
	if clusterType, ok := item["clusterType"]; ok {
		cluster.Type = clusterType.(string)
	}
	if uuid, ok := item["uuid"]; ok {
		cluster.UUID = uuid.(string)
	}
	clusters = append(clusters, cluster)
	return clusters
}

func createHostsType(item Item, hosts []Host) []Host {
	host := Host{}
	if hostID, ok := item["hostId"]; ok {
		host.HostID = hostID.(string)
	}
	if hostName, ok := item["hostname"]; ok {
		host.HostName = hostName.(string)
	}
	if ip, ok := item["ipAddress"]; ok {
		host.IPAddress = ip.(string)
	}
	if commissionState, ok := item["commissionState"]; ok {
		host.CommissionState = commissionState.(string)
	}
	if rackID, ok := item["rackId"]; ok {
		host.RackID = rackID.(string)
	}
	if totalMemVal, ok := item["totalPhysMemBytes"]; ok {
		totalMemBytes := int64(totalMemVal.(float64))
		host.TotalMemory = ByteCountDecimal(totalMemBytes)
	}
	if clusterRefVal, ok := item["clusterRef"]; ok {
		clusterRef := clusterRefVal.(map[string]interface{})
		if clusterNameVal, ok := clusterRef["clusterName"]; ok {
			host.ClusterName = clusterNameVal.(string)
		}
		if clusterDisplayNameVal, ok := clusterRef["displayName"]; ok {
			host.ClusterDisplayName = clusterDisplayNameVal.(string)
		}
	}
	hosts = append(hosts, host)
	return hosts
}

func fillServicesAndRolesFromCluster(clusterMap map[string]interface{}, deployment Deployment) Deployment {
	services := deployment.Services
	clusterName := clusterMap["name"].(string)
	deployment.ClusterServiceRoleMap[clusterName] = ServiceRolesMap{RolesMap: make(map[string][]Role)}
	if servicesVal, ok := clusterMap["services"]; ok {
		servicesI := servicesVal.([]interface{})
		for _, serviceItemVal := range servicesI {
			serviceItem := serviceItemVal.(map[string]interface{})
			services = createServiceType(serviceItem, services, clusterName)
			deployment = fillRolesFromClusterAndService(serviceItem, clusterName, deployment)
		}
	}
	deployment.Services = services
	return deployment
}

func fillRolesFromClusterAndService(serviceMap map[string]interface{}, clusterName string, deployment Deployment) Deployment {
	roleArray := make([]Role, 0)
	serviceName := serviceMap["name"].(string)
	serviceRolesMap := deployment.ClusterServiceRoleMap[clusterName].RolesMap
	if rolesVal, ok := serviceMap["roles"]; ok {
		rolesI := rolesVal.([]interface{})
		for _, roleItemVal := range rolesI {
			roleItem := roleItemVal.(map[string]interface{})
			roleArray = createRoleType(roleItem, roleArray, clusterName, serviceName)
		}
	}
	serviceRolesMap[serviceName] = roleArray
	return deployment
}
