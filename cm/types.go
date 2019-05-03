package cm

// CMServer represents registered CM server entry details
type CMServer struct {
	Name              string `json:"name"`
	Hostname          string `json:"hostname"`
	Port              int    `json:"port"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	Protocol          string `json:"protocol"`
	Cluster           string `json:"cluster"`
	Active            bool   `json:"active"`
	ConnectionProfile string `json:"profile"`
}

// ConnectionProfile represents ssh/connection descriptions which is used to communicate with CM server and agents
type ConnectionProfile struct {
	Name     string `json:"name"`
	KeyPath  string `json:"key_path"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}
