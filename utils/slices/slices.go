package slices

func Unpack(slice []string, vars... *string) {
	for i, str := range slice {
		*vars[i] = str
	}
}

func Filter(slice []string, filter func(string) bool) []string {
	retSlice := []string{}

	for _, element := range slice {
		if filter(element) {
			retSlice = append(retSlice, element)
		}
	}

	return retSlice
}
