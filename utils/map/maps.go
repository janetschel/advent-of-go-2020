package maps

func All(check map[string]bool, eq bool) bool {
	for _, curr:= range check {
		if curr != eq {
			return false
		}
	}

	return true
}

func Pop(pop map[string]int) map[string]int {
	popped := make(map[string]int)

	index := -1
	for key, value := range pop {
		index++
		if index != 0 {
			popped[key] = value
		}
	}

	return popped
}

func Peek(peek map[string]int) string {
	for key, _ := range peek {
		return key
	}

	// never reacher
	return ""
}