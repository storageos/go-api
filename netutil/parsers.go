package netutil

import (
	"net"
	"net/url"
	"strconv"
	"strings"
)

func addrsFromNodes(nodes []string) ([]string, error) {
	var addrs []string

	for _, n := range nodes {
		switch {
		// Assume that the node is provided as a URL
		case strings.Contains(n, "://"):
			newAddrs, err := parseURL(n)
			if err != nil {
				return nil, newInvalidNodeError(err)
			}

			addrs = append(addrs, newAddrs...)

		// Assume the node is in hostname:port or ip:port format
		case strings.Contains(n, ":"):
			newAddrs, err := parseHostPort(n)
			if err != nil {
				return nil, newInvalidNodeError(err)
			}

			addrs = append(addrs, newAddrs...)

		// Assume hostname or ip
		default:
			newAddrs, err := parseHost(n)
			if err != nil {
				return nil, newInvalidNodeError(err)
			}

			addrs = append(addrs, newAddrs...)
		}
	}

	return addrs, nil
}

func validPort(port string) bool {
	intPort, err := strconv.Atoi(port)

	return (err == nil) &&
		(intPort > 0) &&
		(intPort <= 65535)
}

func parseURL(node string) ([]string, error) {
	url, err := url.Parse(node)
	if err != nil {
		return nil, err
	}

	// Verify a valid scheme
	switch url.Scheme {
	case "tcp", "http", "https":
		host, port, err := net.SplitHostPort(url.Host)
		if err != nil {
			// We could be here as there is no port, lets try one last time with default port added
			host, port, err = net.SplitHostPort(url.Host + ":" + DefaultDialPort)
			if err != nil {
				return nil, err
			}
		}

		if port == "" {
			port = DefaultDialPort
		}

		if !validPort(port) {
			return nil, errInvalidPortNumber
		}

		// LookupHost works for IP addr too
		addrs, err := net.LookupHost(host)
		if err != nil {
			return nil, err
		}

		for i, a := range addrs {
			addrs[i] = a + ":" + port
		}

		return addrs, nil

	default:
		return nil, errUnsupportedScheme
	}
}

func parseHostPort(node string) ([]string, error) {
	host, port, err := net.SplitHostPort(node)
	if err != nil {
		return nil, err
	}

	if !validPort(port) {
		return nil, errInvalidPortNumber
	}

	// LookupHost works for IP addr too
	addrs, err := net.LookupHost(host)
	if err != nil {
		return nil, err
	}

	for i, a := range addrs {
		addrs[i] = a + ":" + port
	}

	return addrs, nil
}

func parseHost(node string) ([]string, error) {
	return parseHostPort(node + ":" + DefaultDialPort)
}
