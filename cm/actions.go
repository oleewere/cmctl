package cm

import (
	"bytes"
	"fmt"
	"strings"
)

// ListClusters get all the registered clusters
func (c CMServer) ListClusters() []Cluster {
	var cmItems Items
	var uri = "clusters"
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	return cmItems.ConvertClustersResponse()
}

// ListHosts get all the registered hosts
func (c CMServer) ListHosts() []Host {
	var cmItems Items
	var uri = "hosts"
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	return cmItems.ConvertHostsResponse()
}

// ListServices get all the registered services per cluster
func (c CMServer) ListServices(cluster string) []Service {
	var cmItems Items
	var uri = fmt.Sprintf("clusters/%v/services", cluster)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	return cmItems.ConvertServicesResponse(cluster)
}

// GetDeployment get full deployement data for CM server
func (c CMServer) GetDeployment() Deployment {
	var deploymentMap map[string]interface{}
	var uri = fmt.Sprintf("cm/deployment")
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		deploymentMap = ProcessAsMapFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		deploymentMap = ProcessAsMap(request)
	}
	return ConvertDeploymentResponse(deploymentMap)
}

// ExportClusterTemplate exporting template for a specific cluster
func (c CMServer) ExportClusterTemplate(cluster string) []byte {
	var uri = fmt.Sprintf("clusters/%v/export", cluster)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		return []byte(c.RunGatewayCMCommand(curlCommand, false, true).StdOut)
	}
	request := c.CreateGetRequest(uri)
	return ProcessRequest(request)
}

// GetUsers returns a list of the user names configured in the system.
func (c CMServer) GetUsers() []User {
	var cmItems Items
	var uri = "users"
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	return cmItems.ConvertUsersResponse()
}

// GetExternalAccountCategories returns a list of the supported external account categories
func (c CMServer) GetExternalAccountCategories() []string {
	var cmItems Items
	var uri = "externalAccounts/supportedCategories"
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	supportedCategories := make([]string, 0)
	for _, item := range cmItems.Items {
		if nameVal, ok := item["name"]; ok {
			name := nameVal.(string)
			supportedCategories = append(supportedCategories, name)
		}
	}
	return supportedCategories
}

// GetExternalAccountTypesByCategory returns a list of the external accounts types by category
func (c CMServer) GetExternalAccountTypesByCategory(category string) []string {
	var cmItems Items
	var uri = fmt.Sprintf("externalAccounts/supportedTypes/%v", category)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	accountTypes := make([]string, 0)
	for _, item := range cmItems.Items {
		if nameVal, ok := item["type"]; ok {
			name := nameVal.(string)
			accountTypes = append(accountTypes, name)
		}
	}
	return accountTypes
}

// GetExternalAccountsByType gather external accounts with configs by type
func (c CMServer) GetExternalAccountsByType(accountType string) map[string]map[string]string {
	var cmItems Items
	var uri = fmt.Sprintf("externalAccounts/type/%v", accountType)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	return cmItems.ConvertExternalAccounts()
}

// RunServiceOperation run a service operation (start / stop / restart)
func (c CMServer) RunServiceOperation(cluster string, service string, command string, verbose bool) []byte {
	var response []byte
	var uri = fmt.Sprintf("clusters/%v/services/%s/commands/%v", cluster, service, command)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlPostCommand(uri, "")
		response = []byte(c.RunGatewayCMCommand(curlCommand, verbose, true).StdOut)
	} else {
		request := c.CreatePostRequest(bytes.Buffer{}, uri)
		response = ProcessRequest(request)
		if verbose {
			fmt.Println(string(response))
		}
	}
	return response
}

// RunRolesOperation run operation on a list of roles for a specific service
func (c CMServer) RunRolesOperation(cluster string, service string, roleNames []string, roleNameFilter string, command string, verbose bool) []byte {
	var response []byte
	var uri = fmt.Sprintf("clusters/%v/services/%s/roleCommands/%v", cluster, service, command)
	var joinedRoleNames string
	if len(roleNameFilter) > 0 {
		joinedRoleNames = roleNameFilter
	} else {
		joinedRoleNames = strings.Join(AddQuots(roleNames), ",")
	}
	postBody := fmt.Sprintf("{\"items\": [%v]}", joinedRoleNames)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlPostCommand(uri, postBody)
		response = []byte(c.RunGatewayCMCommand(curlCommand, verbose, true).StdOut)
	} else {
		bodyBytesBuffer := bytes.Buffer{}
		bodyBytesBuffer.WriteString(postBody)
		request := c.CreatePostRequest(bodyBytesBuffer, uri)
		response = ProcessRequest(request)
		if verbose {
			fmt.Println(string(response))
		}
	}
	return response
}

// ListServiceConfigs gather service configurations
func (c CMServer) ListServiceConfigs(cluster string, service string, full bool) []ConfigItem {
	var cmItems Items
	fullView := ""
	if full {
		fullView = "?view=FULL"
	}
	var uri = fmt.Sprintf("clusters/%v/services/%v/config%v", cluster, service, fullView)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	return cmItems.ConvertServiceConfigResponse()
}

// ListRoleConfigGroups gather role configuration groups
func (c CMServer) ListRoleConfigGroups(cluster string, service string, full bool) []RoleConfigGroup {
	var cmItems Items
	fullView := ""
	if full {
		fullView = "?view=FULL"
	}
	var uri = fmt.Sprintf("clusters/%v/services/%v/roleConfigGroups%v", cluster, service, fullView)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlGetCommand(uri)
		cmItems = ProcessCMItemsFromSSHResponse(c.RunGatewayCMCommand(curlCommand, false, true))
	} else {
		request := c.CreateGetRequest(uri)
		cmItems = ProcessCMItems(request)
	}
	return cmItems.ConvertRoleConfigGroupsResponse()
}

// UpdateServiceConfigs refresh service configuration
func (c CMServer) UpdateServiceConfigs(cluster string, service string, keyValueMap map[string]string, clear bool, verbose bool) []byte {
	var response []byte
	var uri = fmt.Sprintf("clusters/%v/services/%v/config", cluster, service) + "?message=Service%20%27" + service + "%27%20config%20update%20by%20CMCTL."
	putBody := createItemsMapListString(keyValueMap, clear)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlPutCommand(uri, putBody)
		response = []byte(c.RunGatewayCMCommand(curlCommand, verbose, false).StdOut)
	} else {
		bodyBytesBuffer := bytes.Buffer{}
		bodyBytesBuffer.WriteString(putBody)
		request := c.CreatePostRequest(bodyBytesBuffer, uri)
		response = ProcessRequest(request)
		if verbose {
			fmt.Println(string(response))
		}
	}
	return response
}

// UpdateRoleConfigs refresh role group configuration
func (c CMServer) UpdateRoleConfigs(cluster string, service string, roleType string, roleConfigGroup string,
	keyValueMap map[string]string, clear bool, verbose bool) []byte {
	var response []byte
	var uri = fmt.Sprintf("clusters/%v/services/%v/roleConfigGroups/%v/config", cluster, service, roleConfigGroup) +
		"?message=Role%20%27" + roleConfigGroup + "%27%20config%20update%20by%20CMCTL."
	putBody := createItemsMapListString(keyValueMap, clear)
	if c.UseGateway {
		curlCommand := c.CreateGatewayCurlPutCommand(uri, putBody)
		response = []byte(c.RunGatewayCMCommand(curlCommand, verbose, false).StdOut)
	} else {
		bodyBytesBuffer := bytes.Buffer{}
		bodyBytesBuffer.WriteString(putBody)
		request := c.CreatePostRequest(bodyBytesBuffer, uri)
		response = ProcessRequest(request)
		if verbose {
			fmt.Println(string(response))
		}
	}
	return response
}

func createItemsMapListString(keyValueMap map[string]string, clear bool) string {
	itemList := make([]string, 0)
	for key, value := range keyValueMap {
		valueStr := "\"" + value + "\""
		if clear {
			valueStr = "null"
		}
		item := fmt.Sprintf("{\"name\":\"%v\",\"value\":%v}", key, valueStr)
		itemList = append(itemList, item)
	}
	return fmt.Sprintf("{\"items\": [%v]}", strings.Join(itemList, ","))
}
