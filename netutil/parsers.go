package netutil

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	httpScheme  = "http"
	httpsScheme = "https"
	tcpScheme   = "tcp"
)

// AddressesFromNodes takes a list of node hosts and attempts to return a list of hosts in host:port
// format along with any error encountered.
//
// The function accepts node hosts in URL, ip, ip:port, resolvable-name and resolvable-name:port
// formats and will append the default port value if needed. For hosts where the scheme has been omitted,
// the scheme for the first host will be used. If the first host has no scheme, it will default to http.
func AddressesFromNodes(nodes []string) ([]string, error) {
	addresses := make([]string, 0, len(nodes))

	for _, node := range nodes {
		address, err := AddressFromNode(node)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

// AddressFromNode takes a node host and attempts to return a host in host:port
// format along with any error encountered
func AddressFromNode(node string) (string, error) {
	var scheme string
	address := node
	// If no scheme present, set the first scheme
	if !strings.Contains(address, "://") {
		if scheme == "" {
			scheme = httpScheme
		}
		address = strings.Join([]string{scheme, address}, "://")
	}

	url, err := url.Parse(address)
	if err != nil {
		return "", newInvalidNodeError(err)
	}

	switch url.Scheme {
	case tcpScheme:
		url.Scheme = httpScheme
		fallthrough
	case httpScheme, httpsScheme:
		if scheme == "" {
			scheme = url.Scheme
		}
	default:
		return "", newInvalidNodeError(errUnsupportedScheme)
	}

	host := url.Hostname()
	if host == "" {
		return "", newInvalidNodeError(errInvalidHostName)
	}

	port := url.Port()
	if port == "" {
		url.Host += ":" + DefaultDialPort
	}
	if !validPort(url.Port()) {
		return "", newInvalidNodeError(errInvalidPortNumber)
	}

	return strings.TrimRight(url.String(), "/"), nil
}

func validPort(port string) bool {
	intPort, err := strconv.Atoi(port)

	return (err == nil) &&
		(intPort > 0) &&
		(intPort <= 65535)
}
