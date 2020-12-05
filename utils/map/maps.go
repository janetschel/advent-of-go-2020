package maps

func All(check map[string]bool, eq bool) bool {
	for _, curr:= range check {
		if curr != eq {
			return false
		}
	}

	return true
}
