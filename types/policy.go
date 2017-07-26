package types

import (
	"encoding/json"
)

type Policy struct {
	Spec struct {
		User            string `json:"user,omitempty"`
		Group           string `json:"group,omitempty"`
		Readonly        bool   `json:"readonly,omitempty"`
		APIGroup        string `json:"apiGroup,omitempty"`
		Resource        string `json:"resource,omitempty"`
		Namespace       string `json:"namespace,omitempty"`
		NonResourcePath string `json:"nonResourcePath,omitempty"`
	} `json:"spec"`
}

type PolicyWithID struct {
	*Policy
	ID string
}

func (p *PolicyWithID) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Policy)
}

type PolicySet map[string]*Policy

func (p PolicySet) GetPoliciesWithID() []*PolicyWithID {
	rtn := make([]*PolicyWithID, 0, len(p))

	for k, v := range p {
		rtn = append(rtn, &PolicyWithID{
			Policy: v,
			ID:     k,
		})
	}

	return rtn
}
