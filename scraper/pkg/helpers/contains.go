package helpers

func Contains(value string, items []string) bool {
	for _, i := range items {
		if i == value {
			return true
		}
	}
	return false
}
