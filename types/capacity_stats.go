package types

// ErrCapacityStatsUnchanged can be used when comparing whether stats
const ErrCapacityStatsUnchanged = "no changes"

// CapacityStats is used to report capacity statistics.
type CapacityStats struct {
	TotalCapacityBytes       uint64 `json:"total_capacity_bytes"`
	AvailableCapacityBytes   uint64 `json:"available_capacity_bytes"`
	ProvisionedCapacityBytes uint64 `json:"provisioned_capacity_bytes"`
}

// IsEqual checks if capacity values are the same
func (c CapacityStats) IsEqual(n CapacityStats) bool {
	if c == n {
		return true
	}
	return false
}
