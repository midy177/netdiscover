package discover

import (
	"net"
)

const (
	aliyunPrivateIPv4URL = "http://100.100.100.200/latest/meta-data/private-ipv4"
	aliyunPublicIPv4URL  = "http://100.100.100.200/latest/meta-data/eipv4"
	aliyunHostnameURL    = "http://100.100.100.200/latest/meta-data/hostname"
)

// NewALIYUNDiscoverer returns a new Aliyun Web Services network discoverer
func NewALIYUNDiscoverer() Discoverer {
	return NewDiscoverer(
		PrivateIPv4DiscovererOption(aliyunPrivateIPv4),
		PublicIPv4DiscovererOption(aliyunPublicIPv4),
		PublicHostnameDiscovererOption(aliyunHostname),
	)

}

func aliyunPrivateIPv4() (net.IP, error) {
	return StandardIPFromHTTP(aliyunPrivateIPv4URL, nil)
}

func aliyunPublicIPv4() (net.IP, error) {
	return StandardIPFromHTTP(aliyunPublicIPv4URL, nil)
}

func aliyunHostname() (string, error) {
	return StandardHostnameFromHTTP(aliyunHostnameURL, nil)
}
