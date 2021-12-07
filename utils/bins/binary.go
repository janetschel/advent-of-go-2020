package bins

import "strconv"

func Pad(binary string, length int) string {
	for len(binary) < length {
		binary = "0" + binary
	}

	return binary
}

func AllBinaryNumbers(length int) []string {
	binaries := make([]string, 0)

	to := ""
	for i := 0; i < length; i++ {
		to += "1"
	}

	var binary string
	for i := 0; binary != to; i++ {
		binary = strconv.FormatInt(int64(i), 2)
		binaries = append(binaries, Pad(binary, length))
	}

	return binaries
}
