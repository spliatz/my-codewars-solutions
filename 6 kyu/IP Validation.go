package kata

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	arr := strings.Split(ip, ".")
	if len(arr) != 4 {
		return false
	}

	for _, s := range arr {
		if len(strings.Split(s, "")) > 1 && strings.Split(s, "")[0] == "0" {
			return false
		}
		if number, err := strconv.Atoi(s); err != nil || number > 255 || number < 0 {
			return false
		}
	}

	return true
}
