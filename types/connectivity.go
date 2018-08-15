package types

import "time"

type TestName string

const (
	APIConnectivity  TestName = "api"
	NatsConnectivity TestName = "nats"
	EtcdConnectivity TestName = "etcd"
)

// ConnectivityResult capture's a node connectivity report to a given target.
type ConnectivityResult struct {
	// Label is a human-readable reference for the service being tested.
	Label string `json:"label"`

	// Address is the host:port of the service being tested.
	Address string `json:"address"`

	// Source is a human-readable reference for the source host where the tests
	// were run from..
	Source string `json:"source"`

	// LatencyNS is the duration in nanoseconds that the check took to complete.
	// Will also be set on unsuccessful attempts.
	LatencyNS time.Duration `json:"latency_ns"`

	// Error is set if the test returned an error.
	Error string `json:"error"`
}

// Passes returns true iff no error
func (r ConnectivityResult) Passes() bool {
	return len(r.Error) == 0
}
