package storageos

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/storageos/go-api/types"
	"net/http"
	"net/url"
)

var (

	// PolicyAPIPrefix is a partial path to the HTTP endpoint.
	PolicyAPIPrefix = "policies"

	// ErrNoSuchPolicy is the error returned when the policy does not exist.
	ErrNoSuchPolicy = errors.New("no such policy")
)

type nopMarshaler []byte

func (n *nopMarshaler) MarshalJSON() ([]byte, error) {
	return *n, nil
}

// PolicyCreate creates a policy on the server.
func (c *Client) PolicyCreate(jsonl []byte, ctx context.Context) error {
	nopm := nopMarshaler(jsonl)
	_, err := c.do("POST", PolicyAPIPrefix, doOptions{
		data:    &nopm,
		context: ctx,
		headers: map[string]string{"Content-Type": "application/x-jsonlines"},
	})
	return err
}

func (c *Client) Policy(id string) (*types.Policy, error) {
	path := fmt.Sprintf("%s/%s", PolicyAPIPrefix, id)
	resp, err := c.do("GET", path, doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok && e.Status == http.StatusNotFound {
			return nil, ErrNoSuchPolicy
		}
		return nil, err
	}
	defer resp.Body.Close()

	var policy *types.Policy
	if err := json.NewDecoder(resp.Body).Decode(&policy); err != nil {
		return nil, err
	}
	return policy, nil
}

func (c *Client) PolicyList(opts types.ListOptions) (types.PolicySet, error) {
	listOpts := doOptions{
		fieldSelector: opts.FieldSelector,
		labelSelector: opts.LabelSelector,
		namespace:     opts.Namespace,
		context:       opts.Context,
	}

	if opts.LabelSelector != "" {
		query := url.Values{}
		query.Add("labelSelector", opts.LabelSelector)
		listOpts.values = query
	}

	resp, err := c.do("GET", PolicyAPIPrefix, listOpts)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var policies types.PolicySet
	if err := json.NewDecoder(resp.Body).Decode(&policies); err != nil {
		return nil, err
	}
	return policies, nil
}

func (c *Client) PolicyDelete(opts types.DeleteOptions) error {
	resp, err := c.do("DELETE", PolicyAPIPrefix+"/"+opts.Name, doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok {
			if e.Status == http.StatusNotFound {
				return ErrNoSuchPolicy
			}
		}
		return err
	}
	defer resp.Body.Close()
	return nil
}
