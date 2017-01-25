package types

import (
	"context"
	"time"
)

// Replica - defines replica
type Replica struct {
	Controller string    `json:"controller"`
	ID         string    `json:"id"`
	Inode      uint32    `json:"inode"`
	Status     string    `json:"status"`
	Health     string    `json:"health"`
	CreatedAt  time.Time `json:"created_at"`
}

// Master - master volume details
type Master struct {
	Controller string    `json:"controller"`
	ID         string    `json:"id"`
	Inode      uint32    `json:"inode"`
	Status     string    `json:"status"`
	Health     string    `json:"health"`
	CreatedAt  time.Time `json:"created_at"`
}

// Volume is used to represent a presentable storage device
type Volume struct {
	ID string `json:"id,omitempty"`

	Virtual  bool      `json:"virtual,omitempty"`
	Master   Master    `json:"master"`
	Replicas []Replica `json:"replicas"` // data about replicas

	CreatedBy string `json:"created_by"`

	Datacenter    string `json:"datacentre"`
	Tenant        string `json:"tenant"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Health        string `json:"health"`
	Pool          string `json:"pool"`
	Description   string `json:"description"`
	Size          int    `json:"size"`

	Inode uint32 `json:"inode"` // redirect inode

	Dev          string   `json:"dev,omitempty"`
	VolumeGroups []string `json:"volume_groups" mapstructure:"volume_groups"`
	Tags         []string `json:"tags"`

	Deleted bool `json:"deleted,omitempty"`

	// Controller UUIDs for replication
	// Controller string `json:"controller" mapstructure:"controller"`

	Mounted        bool `json:"mounted"`
	NumberOfMounts int  `json:"no_of_mounts"`
	// client ID
	MountedBy string    `json:"mounted_by"`
	MountedAt time.Time `json:"mounted_at"`

	CreatedAt time.Time `json:"created_at"`

	str string // memoize string
}

// ListVolumeOptions specifies parameters to the ListVolumes function.
type ListVolumeOptions struct {
	Filters map[string][]string
	All     bool
	Digests bool
	Filter  string
	Context context.Context
}

// CreateVolumeOptions specify parameters to the CreateVolume function.
type CreateVolumeOptions struct {
	Name        string
	Description string
	Pool        string
	Size        int
	Context     context.Context `json:"-"`
	Labels      map[string]string
}
