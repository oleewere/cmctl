package cm

// CMServer represents registered CM server entry details
type CMServer struct {
	Name              string `json:"name"`
	Hostname          string `json:"hostname"`
	Port              int    `json:"port"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	Protocol          string `json:"protocol"`
	Active            bool   `json:"active"`
	ConnectionProfile string `json:"profile"`
	UseGateway        bool   `json:"gateway"`
}

// ConnectionProfile represents ssh/connection descriptions which is used to communicate with CM server and agents
type ConnectionProfile struct {
	Name     string `json:"name"`
	KeyPath  string `json:"key_path"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}

// CMItems global items from CM Server rest API response
type CMItems struct {
	Href  string `json:"href"`
	Items []Item `json:"items"`
}

// Item dynamic map - cast contents to specific types
type Item map[string]interface{}

// Cluster holds installed CM cluster details
type Cluster struct {
	UUID        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Version     string `json:"fullVersion,omitempty"`
	Type        string `json:"clusterType,omitempty"`
}

// Host holds host details
type Host struct {
	HostID             string `json:"hostId,omitempty"`
	IPAddress          string `json:"ipAddress,omitempty"`
	HostName           string `json:"hostName,omitempty"`
	CommissionState    string `json:"commissionState,omitempty"`
	RackID             string `json:"rackId,omitempty"`
	ClusterName        string `json:"clusterName,omitempty"`
	ClusterDisplayName string `json:"clusterDisplayName,omitempty"`
}
