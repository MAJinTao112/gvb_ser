package utils

func InList(key string, list []string) bool {
	for _, value := range list {
		if value == key {
			return true
		}
	}
	return false
}
