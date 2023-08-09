package utils

import (
	"math/big"
	"strings"
)

func OneOfStrings(mainString string, stringList ...string) bool {
	for _, str := range stringList {
		if mainString == str {
			return true
		}
	}
	return false
}

func CompareStringNumber(a, b string) int {
	bigA, _ := new(big.Int).SetString(a, 10)
	bigB, _ := new(big.Int).SetString(b, 10)
	return bigA.Cmp(bigB)
}

func RemoveUUIDStrikeThrough(uuidStr string) string {
	uuidString := strings.Replace(uuidStr, "-", "", -1)
	return uuidString
}

func AddUUIDStrikeThrough(str string) string {
	strUUID := str[:8] + "-" + str[8:12] + "-" + str[12:16] + "-" + str[16:20] + "-" + str[20:]
	return strUUID
}
