package desens

import "strings"

func DesensitizationEmail(email string) string {
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

func DesensitizationIP(ip string) string {
	if ip == "" {
		return ""
	}

	// IPv4处理：192.168.1.1 -> 192.***.1
	if strings.Contains(ip, ".") {
		parts := strings.Split(ip, ".")
		if len(parts) == 4 {
			return parts[0] + "." + "***" + "." + parts[3]
		}
	}

	// IPv6处理：2001:0db8:85a3::8a2e:0370:7334 -> 2001:****
	if strings.Contains(ip, ":") {
		return ip[:6] + "****"
	}

	return "***"
}
