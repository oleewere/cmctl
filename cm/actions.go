package cm

// ListClusters get all the registered clusters
func (c CMServer) ListClusters() []Cluster {
	var cmItems CMItems
	var uri = "clusters"
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	return cmItems.ConvertClustersResponse()
}

// ListHosts get all the registered hosts
func (c CMServer) ListHosts() []Host {
	var cmItems CMItems
	var uri = "hosts"
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	return cmItems.ConvertHostsResponse()
}
