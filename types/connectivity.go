package types

type ConnectivityStatus struct {
	CanConnect  bool     `json:"canConnect"`
	FailedTests []string `json:"failedTests,omitempty"`
}

type NodeConnectivity struct {
	Nodes map[string]ConnectivityStatus `json:"nodes"`
}
