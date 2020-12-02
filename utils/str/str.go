package str

import "errors"

func CharAt(str string, pos int) (string, error) {
	if pos >= len(str) {
		return "", errors.New("invalid index")
	}

	return string(str[pos]), nil
}
