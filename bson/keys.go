package tagit

func keys(input map[string]struct{}) []string {
	keys := make([]string, len(input))
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	return keys
}
