package internal

import (
	"strconv"
	"strings"
)

func is_valid_hex_string(code string) bool {
	const valid string = "0123456789abcdefABCDEF"

	for _, v := range code {
		if !strings.ContainsRune(valid, v) {
			return false
		}
	}
	return true
}

func i32ToString(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}
