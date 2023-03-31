// +build !darwin,!linux,!windows,!solaris,!freebsd,!openbsd

package gateway

import (
	"net"
)

func discoverGatewayOSSpecific() (ip net.IP, err error) {
	return ip, errNotImplemented
}

func discoverGatewayInterfaceOSSpecific() (ip net.IP, err error) {
	return nil, errNotImplemented
}
