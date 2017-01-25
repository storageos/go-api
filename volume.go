package storageos

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/storageos/go-api/types"
)

var (
	// ErrNoSuchVolume is the error returned when the volume does not exist.
	ErrNoSuchVolume = errors.New("no such volume")

	// ErrVolumeInUse is the error returned when the volume requested to be removed is still in use.
	ErrVolumeInUse = errors.New("volume in use and cannot be removed")
)

// ListVolumes returns the list of available volumes.
func (c *Client) ListVolumes(opts types.ListVolumeOptions) ([]types.Volume, error) {
	path := "/volumes?" + queryString(opts)
	resp, err := c.do("GET", path, doOptions{context: opts.Context})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var volumes []types.Volume
	if err := json.NewDecoder(resp.Body).Decode(&volumes); err != nil {
		return nil, err
	}
	return volumes, nil
}

// CreateVolume creates a volume on the server.
func (c *Client) CreateVolume(opts types.CreateVolumeOptions) (*types.Volume, error) {
	resp, err := c.do("POST", "/volumes", doOptions{
		data:    opts,
		context: opts.Context,
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var volume types.Volume
	if err := json.NewDecoder(resp.Body).Decode(&volume); err != nil {
		return nil, err
	}
	return &volume, nil
}

// GetVolume returns a volume by its reference.
func (c *Client) GetVolume(ref string) (*types.Volume, error) {
	resp, err := c.do("GET", "/volumes/"+ref, doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok && e.Status == http.StatusNotFound {
			return nil, ErrNoSuchVolume
		}
		return nil, err
	}
	defer resp.Body.Close()
	var volume types.Volume
	if err := json.NewDecoder(resp.Body).Decode(&volume); err != nil {
		return nil, err
	}
	return &volume, nil
}

// RemoveVolume removes a volume by its reference.
func (c *Client) RemoveVolume(ref string) error {
	resp, err := c.do("DELETE", "/volumes/"+ref, doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok {
			if e.Status == http.StatusNotFound {
				return ErrNoSuchVolume
			}
			if e.Status == http.StatusConflict {
				return ErrVolumeInUse
			}
		}
		return nil
	}
	defer resp.Body.Close()
	return nil
}
