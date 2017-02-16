package types

import "context"

// VolumeCreateOptions are available parameters for creating new volumes.
type VolumeCreateOptions struct {

	// Name is the name of the volume to create.
	// Required: true
	Name string

	// Description describes the volume.
	Description string

	// Size in GB.
	// Required: true
	Size int

	// Pool is the name or id of capacity pool to provision the volume in.
	Pool string

	// Namespace is the object scope, such as for teams and projects.
	Namespace string

	// Labels are user-defined key/value metadata.
	Labels map[string]string

	// Context can be set with a timeout or can be used to cancel a request.
	Context context.Context
}
