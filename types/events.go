package types

import "time"

// EventType describes the type of event
type EventType string

// EventTypes are added to events to assist with type assertions
const (
	RequestType   EventType = "request"
	ResponseType            = "response"
	HeartbeatType           = "heartbeat"
	BackupType              = "backup"
)

// Event describes the fields that all events should implement.  Event is
// intended to be inherherited in more specific Event types.
type Event struct {
	ID string `json:"id"`
	// Parent is used to specify parent event
	Parent          string    `json:"parent"`
	EventType       EventType `json:"event_type"`
	Action          string    `json:"action"`
	Timestamp       int64     `json:"timestamp"`
	Status          string    `json:"status"`
	Message         string    `json:"message"`
	Log             []string  `json:"log"`
	ProgressPercent int       `json:"progress_percent"`
	CreatedBy       string    `json:"created_by"`

	Target        string      `json:"target"`
	ActionPayload interface{} `json:"action_payload"`

	// payload can be encoded into bytes as well
	// ActionPayloadType  string `json:"action_payload_type"`
	ActionPayloadBytes []byte `json:"action_payload_bts"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	// retry related value
	Retry     bool      `json:"retry"`
	RetriedAt time.Time `json:"retried_at"`
	Attempts  int       `json:"attempts"`

	// optional parameter
	Deadline time.Time `json:"deadline"`

	// optional events to dispatch
	Rollback     []*Request `json:"rollback"`
	RollbackDone bool       `json:"rollback_done"`

	Subject string `json:"subject"` // or "queue""

	// controller ID which created this event
	OriginController string `json:"origin_controller"`
}

// Request is the message structure used for sending request events
type Request struct {
	Event
}
