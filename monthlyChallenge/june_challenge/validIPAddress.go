package june_challenge

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	arr := strings.Split(IP, ".")
	if len(arr) == 4 {
		for _, v := range arr {
			bytes := []byte(v)
			if len(bytes) == 0 || len(bytes) > 3 || (bytes[0] == '0' && len(bytes) > 1) {
				return "Neither"
			}
			for _, b := range bytes {
				if b < 48 || b > 57 {
					return "Neither"
				}
			}
			if num, _ := strconv.ParseInt(v, 10, 64); num > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	} else {
		arr = strings.Split(IP, ":")
		if len(arr) == 8 {
			for _, v := range arr {
				bytes := []byte((v))
				if len(bytes) > 4 || len(bytes) == 0 {
					return "Neither"
				}
				for _, b := range bytes {
					if !((b >= 48 && b <= 57) || (b >= 65 && b <= 70) || (b >= 97 && b <= 102)) {
						return "Neither"
					}
				}
			}
		} else {
			return "Neither"
		}
		return "IPv6"
	}
}
