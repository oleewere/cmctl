package cm

// ConvertResponse converts the response items to specific types
func (c CMItems) ConvertResponse() Response {
	response := Response{}
	hosts := []Host{}
	clusters := []Cluster{}
	for _, item := range c.Items {
		clusters = createClustersType(item, clusters)
		hosts = createHostsType(item, hosts)
	}
	if len(hosts) > 0 {
		response.Hosts = hosts
	}
	if len(hosts) > 0 {
		response.Hosts = hosts
	}
	return response
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
	if rackID, ok := item["commissionState"]; ok {
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
