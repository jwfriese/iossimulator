package stringslice

func Contains(slice []string, value string) bool {
	for _, next := range slice {
		if next == value {
			return true
		}
	}

	return false
}
