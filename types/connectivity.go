package types

type ConnectivityResult struct {
	Target *Node `json:"target"`

	APIConnectivity  bool `json:"apiConnectivity"`
	NatsConnectivity bool `json:"natsConnectivity"`
	EtcdConnectivity bool `json:"etcdConnectivity"`

	Pass    bool    `json:"pass"`    // logical and of all connectivity tests (eg. API && etcd && DFS etc)
	Timeout bool    `json:"timeout"` // true iff the test timed out before completion
	Errors  []error `json:"errors"`  // errors encountered during testing
}
