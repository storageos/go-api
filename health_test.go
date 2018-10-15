package storageos

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/storageos/go-api/types"
)

func TestCPHealth(t *testing.T) {
	data := `{
  "submodules": {
    "kv": {
      "status": "alive",
      "message": "",
      "updatedAt": "2017-08-07T12:54:29Z",
      "changedAt": "2017-08-07T12:46:43Z"
    },
    "kv_write": {
      "status": "alive",
      "message": "",
      "updatedAt": "2017-08-07T12:54:29Z",
      "changedAt": "0001-01-01T00:00:00Z"
    },
    "nats": {
      "status": "alive",
      "message": "",
      "updatedAt": "2017-08-07T12:54:29Z",
      "changedAt": "0001-01-01T00:00:00Z"
    },
    "scheduler": {
      "status": "alive",
      "message": "",
      "updatedAt": "2017-08-07T12:54:29Z",
      "changedAt": "2017-08-07T12:47:03Z"
    }
  }
}`

	var expected types.CPHealthStatus
	if err := json.Unmarshal([]byte(data), &expected); err != nil {
		t.Fatal(err)
	}

	client := newTestClient(&FakeRoundTripper{message: data, status: http.StatusOK})
	cpHealth, err := client.CPHealth(context.Background(), "someHostname")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(*cpHealth, expected) {
		t.Errorf("Wrong return value.\nWant %#v.\nGot  %#v.", expected, *cpHealth)
	}
}

func TestDPHealth(t *testing.T) {
	data := `{
  "submodules": {
    "directfs-initiator": {
      "status": "alive",
      "message": "",
      "updatedAt": "2017-08-07T12:55:57Z",
      "changedAt": "2017-08-07T12:46:58Z"
    },
    "directfs-responder": {
      "status": "alive",
      "message": "",
      "updatedAt": "2017-08-07T12:55:57Z",
      "changedAt": "2017-08-07T12:46:58Z"
    },
    "director": {
      "status": "alive",
      "message": "",
      "updatedAt": "2017-08-07T12:55:57Z",
      "changedAt": "2017-08-07T12:46:58Z"
    },
    "presentation": {
      "status": "alive",
      "message": "",
      "updatedAt": "2017-08-07T12:55:57Z",
      "changedAt": "2017-08-07T12:47:09Z"
    },
    "fs": {
      "status": "alive",
      "message": "",
      "updatedAt": "2017-08-07T12:55:57Z",
      "changedAt": "2017-08-07T12:47:09Z"
    }
  }
}
`
	var expected types.DPHealthStatus
	if err := json.Unmarshal([]byte(data), &expected); err != nil {
		t.Fatal(err)
	}

	client := newTestClient(&FakeRoundTripper{message: data, status: http.StatusOK})
	dpHealth, err := client.DPHealth(context.Background(), "someHostname")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(*dpHealth, expected) {
		t.Errorf("Wrong return value.\nWant %#v.\nGot  %#v.", expected, *dpHealth)
	}
}
