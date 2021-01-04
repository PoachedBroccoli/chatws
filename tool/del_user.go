package tool

// Remove : del user
func Remove(s []string, user string) []string {
	count := len(s)
	if count == 0 {
		return s
	}
	if count == 1 && s[0] == user {
		return []string{}
	}

	var tempSlice []string = []string{}

	for i := range s {
		if s[i] == user && i == count {
			return s[:count]
		} else if s[i] == user {
			tempSlice = append(s[:i], s[i+1:]...)
			break
		}
	}
	return tempSlice
}
