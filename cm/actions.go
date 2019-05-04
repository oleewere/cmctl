package cm

// ListClusters get all the registered clusters
func (c CMServer) ListClusters() []Cluster {
	request := c.CreateGetRequest("clusters")
	cmItems := ProcessCMItems(request)
	return cmItems.ConvertResponse().Clusters
}

// ListHosts get all the registered hosts
func (c CMServer) ListHosts() []Host {
	request := c.CreateGetRequest("hosts")
	cmItems := ProcessCMItems(request)
	return cmItems.ConvertResponse().Hosts
}
