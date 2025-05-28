package utils

import (
	"net"
	"slices"

	"github.com/SisyphusSQ/golib/utils/stringutil"
)

func GetIP() (ipv4 string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, iface := range interfaces {
		// 忽略没有地址的接口
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP

			// 检查地址类型并提取 IP 字段
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 忽略环回地址和 IPv6 地址
			if ip == nil || ip.IsLoopback() {
				continue
			}

			// 使用 IPv4 地址
			if ip.To4() != nil {
				ipv4 = ip.String()
			}
		}
	}
	return
}

func CheckIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func LookupIP(domain string) ([]net.IP, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

func FilOfficeIP(ips string) bool {
	return stringutil.StartWith(ips, "10.2.", "10.3.", "10.4.", "10.5.", "10.6.", "10.7.", "10.8.", "10.9.", "10.10.")
}

func GetIPsFromDomain(instances []string) []string {
	ips := make([]string, 0)
	for _, instance := range instances {
		if CheckIP(instance) {
			ips = append(ips, instance)
			continue
		}

		lookupIPs, err := LookupIP(instance)
		if err != nil {
			//log.Logger.Errorf("[delFilter] lookup ip err: %v", err)
			continue
		}
		//log.Logger.Debugf("[delFilter] lookup ips: %v", lookupIPs)

		for _, lookupIP := range lookupIPs {
			if slices.Contains(ips, lookupIP.String()) {
				continue
			}
			ips = append(ips, lookupIP.String())
		}
	}
	return ips
}
