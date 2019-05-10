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
	ApiVersion        int    `json:"version"`
}

// ConnectionProfile represents ssh/connection descriptions which is used to communicate with CM server and agents
type ConnectionProfile struct {
	Name     string `json:"name"`
	KeyPath  string `json:"key_path"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}

// Items global items from CM Server rest API response
type Items struct {
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
	TotalMemory        string `json:"totalPhysMemBytes,omitempty"`
}

// Service holds service details
type Service struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Type        string `json:"type,omitempty"`
	State       string `json:"serviceState,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	StaleConfig string `json:"configStalenessStatus,omitempty"`
}

// ConfigItem represents a config entry
type ConfigItem struct {
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	Sensitive bool   `json:"sensitive,omitempty"`
}

// RoleConfigGroup holds role configuration group
type RoleConfigGroup struct {
	Name        string       `json:"name,omitempty"`
	DisplayName string       `json:"displayName,omitempty"`
	RoleType    string       `json:"roleType,omitempty"`
	ClusterName string       `json:"clusterName,omitempty"`
	ServiceName string       `json:"serviceName,omitempty"`
	Base        bool         `json:"base,omitempty"`
	ConfigItems []ConfigItem `json:"items,omitempty"`
}

// Role holds role details
type Role struct {
	Name           string `json:"name,omitempty"`
	Type           string `json:"type,omitempty"`
	HostName       string `json:"hostName,omitempty"`
	ConfigGroup    string `json:"configGroup,omitempty"`
	State          string `json:"roleState,omitempty"`
	StaleConfig    string `json:"configStalenessStatus,omitempty"`
	ComissionState string `json:"commissionState,omitempty"`
	ClusterName    string `json:"clusterName,omitempty"`
	ServiceName    string `json:"serviceName,omitempty"`
}

// Deployment hold full CM server deployment data
type Deployment struct {
	Clusters              []Cluster
	Services              []Service
	Hosts                 []Host
	ClusterServiceRoleMap map[string]ServiceRolesMap
}

// ServiceRolesMap holds service - roles map
type ServiceRolesMap struct {
	RolesMap map[string][]Role
}

// User holds user details
type User struct {
	Name      string         `json:"name,omitempty"`
	AuthRoles []UserAuthRole `json:"authRoles,omitempty"`
}

// UserAuthRole holds user auth role details
type UserAuthRole struct {
	Name string `json:"name,omitempty"`
}

// Inventory holds hosts mappings for resources
type Inventory struct {
	ClusterName         string
	Hosts               []Host
	ServiceRoleHostsMap map[string]map[string][]HostRoleNamePair
	ServiceHostsMap     map[string][]string
}

// HostRoleNamePair holds host name and role name for inventory resource
type HostRoleNamePair struct {
	HostName string
	RoleName string
}
