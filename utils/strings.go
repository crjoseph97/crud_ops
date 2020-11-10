package utils

// UniqueStrings remove duplicates in the string and make it unique.
//// TODO: Make this function be generic, by supporting all other types
func UniqueStrings(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// CompareStrings compares the match string with the matchWith string and return true, if it exist.
func CompareStrings(match string, matchWith []string) bool {
	for _, m := range matchWith {
		if m == match {
			return true
		}
	}

	return false
}
