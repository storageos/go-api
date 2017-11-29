package netutil_test

import (
	"github.com/storageos/go-api/netutil"
	"testing"
)

func TestAddressResolution(t *testing.T) {

	var tests = []struct {
		nodeAddr string
		expected string
	}{
		// Valid schemes with port number
		{"tcp://google-public-dns-a.google.com:1234", "8.8.8.8:1234"},
		{"https://google-public-dns-a.google.com:1234", "8.8.8.8:1234"},
		{"http://google-public-dns-a.google.com:1234", "8.8.8.8:1234"},

		// Valid schemes without port number
		{"tcp://google-public-dns-a.google.com", "8.8.8.8:5705"},
		{"https://google-public-dns-a.google.com", "8.8.8.8:5705"},
		{"http://google-public-dns-a.google.com", "8.8.8.8:5705"},

		// Just host with port number
		{"google-public-dns-a.google.com:1234", "8.8.8.8:1234"},
		{"8.8.8.8", "8.8.8.8:5705"},

		// Just host without port number
		{"google-public-dns-a.google.com", "8.8.8.8:5705"},
		{"8.8.8.8:1234", "8.8.8.8:1234"},
	}
	for _, tt := range tests {
		md, err := netutil.NewMultiDialer([]string{tt.nodeAddr}, nil)
		if err != nil {
			t.Fatal(err)
		}

		if len(md.Addresses) != 1 {
			t.Errorf("Unexpected number of addresses %#v", md.Addresses)
			return
		}

		got := md.Addresses[0]
		if got != tt.expected {
			t.Errorf("Got %s. Want %s.", got, tt.expected)
		}
	}
}
