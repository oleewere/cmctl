package cm

import "strings"

// Filter represents filter on agent hosts (by component / service / hosts)
type Filter struct {
	Hosts    []string
	Clusters []string
	Services []string
	Server   bool
}

// CreateFilter will make a Filter object from filter strings (hosts)
func CreateFilter(clusterFilter string, serviceFilter string, hostFilter string, cmServer bool) Filter {
	filter := Filter{}
	if len(hostFilter) > 0 {
		hosts := strings.Split(hostFilter, ",")
		filter.Hosts = hosts
	}
	if len(clusterFilter) > 0 {
		clusters := strings.Split(clusterFilter, ",")
		filter.Clusters = clusters
	}
	if len(serviceFilter) > 0 {
		services := strings.Split(serviceFilter, ",")
		filter.Services = services
	}
	filter.Server = cmServer
	return filter
}

// GetFilteredHosts obtain specific hosts based on different filters
func (c CMServer) GetFilteredHosts(filter Filter) map[string]bool {
	finalHosts := make(map[string]bool)
	hosts := make(map[string]bool) // use boolean map as a set
	// TODO: filter on services and clusters
	if filter.Server {
		hosts[c.Hostname] = true
		finalHosts[c.Hostname] = true
	} else {
		cmAgents := c.ListHosts()
		calculateAndFillFinalHosts(cmAgents, filter, hosts, finalHosts)
	}
	return finalHosts
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
