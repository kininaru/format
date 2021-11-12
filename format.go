package format

import (
	"errors"
	"fmt"
)

func IntToString(n int64, base int) (string, error) {
	if base > len(dict) {
		return "", errors.New(fmt.Sprintf("Invalid base. Must smaller than %d", len(dict)))
	}

	if n == 0 {
		return "0", nil
	}

	negative := false
	if n < 0 {
		n = -n
		negative = true
	}

	ret := ""
	for n > 0 {
		ret = string(dict[n%int64(base)]) + ret
		n /= int64(base)
	}

	if negative {
		ret = "-" + ret
	}

	return ret, nil
}
