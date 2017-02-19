package types

import "time"

// DefaultNamespace is used when a namespace hasn't been specified.
const DefaultNamespace = "default"

// Volume represents storage volume.
// swagger:model Volume
type Volume struct {

	// Volume unique ID.
	// Read Only: true
	ID string `json:"id"`

	// Block device inode.
	// Read Only: true
	Inode uint32 `json:"inode"`

	// Volume name.
	// Required: true
	Name string `json:"name"`

	// Size in GB.
	// Required: true
	Size int `json:"size"`

	// Name of capacity pool to provision the volume in, or the name of the current pool.
	Pool string `json:"pool"`

	// Filesystem type to mount.  May be set on create, or set by rules to influence client.
	FSType string `json:"fsType"`

	// Volume description.
	Description string `json:"description"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"labels"`

	// Namespace is the object name and authentication scope, such as for teams and projects.
	Namespace string `json:"namespace"`

	// Volume deployment information for the master volume.
	// Read Only: true
	Master *Deployment `json:"master,omitempty"`

	// Flag indicating if the volume is mounted and in use.
	// Read Only: true
	Mounted bool `json:"mounted"`

	// When the volume was mounted.
	// Read Only: true
	MountedAt string `json:"mountedAt,omitempty"`

	// Reference to the node that has the volume mounted.
	// Read Only: true
	MountedBy string `json:"mountedBy,omitempty"`

	// Volume deployment information for the replica volumes.
	// Read Only: true
	Replicas []*Deployment `json:"replicas"`

	// Volume health, one of: healthy, degraded or dead.
	// Read Only: true
	Health string `json:"health"`

	// Short status, one of: pending, evaluating, deploying, active, unavailable, failed, updating, deleting.
	// Read Only: true
	Status string `json:"status"`

	// Status message explaining current status.
	// Read Only: true
	StatusMessage string `json:"statusMessage"`

	// When the volume was created.
	// Read Only: true
	CreatedAt time.Time `json:"createdAt"`

	// User that created the volume.
	// Read Only: true
	CreatedBy string `json:"createdBy"`
}
