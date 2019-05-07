package cm

import (
	"fmt"
	"os"
	"strings"
)

// Filter represents filter on agent hosts (by component / service / hosts)
type Filter struct {
	Hosts    []string
	Clusters []string
	Services []string
	Roles    []string
	Server   bool
}

// CreateFilter will make a Filter object from filter strings (hosts)
func CreateFilter(clusterFilter string, serviceFilter string, rolesFilter string, hostFilter string, cmServer bool) Filter {
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
	if len(rolesFilter) > 0 {
		roles := strings.Split(rolesFilter, ",")
		filter.Roles = roles
	}
	filter.Server = cmServer
	validateFilter(filter)
	return filter
}

func validateFilter(filter Filter) {
	if len(filter.Roles) > 0 && len(filter.Services) != 1 {
		fmt.Println("Error: use exactly 1 service filter with roles filter!")
		os.Exit(1)
	}
}
