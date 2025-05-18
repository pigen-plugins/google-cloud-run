package helpers


func PrefixMapKeys(original map[string]string, prefix string) map[string]string {
	newMap := make(map[string]string, len(original))
	for k, v := range original {
		newKey := prefix + "_" + k
		newMap[newKey] = v
	}
	return newMap
}
