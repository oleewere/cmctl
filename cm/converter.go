package cm

// ConvertHostsResponse convert items response to Hosts response
func (c CMItems) ConvertHostsResponse() []Host {
	hosts := []Host{}
	for _, item := range c.Items {
		hosts = createHostsType(item, hosts)
	}
	return hosts
}

// ConvertClustersResponse convert items response to Clusters response
func (c CMItems) ConvertClustersResponse() []Cluster {
	clusters := []Cluster{}
	for _, item := range c.Items {
		clusters = createClustersType(item, clusters)
	}
	return clusters
}

// ConvertServicesResponse convert items response to Services response
func (c CMItems) ConvertServicesResponse(cluster string) []Service {
	services := []Service{}
	for _, item := range c.Items {
		services = createServiceType(item, services, cluster)
	}
	return services
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
