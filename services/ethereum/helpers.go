package ethereum

import "regexp"

func CheckAddress(possibleAddress string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(possibleAddress)
}
