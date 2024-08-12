package web

func EmptyOrNewlines(s string) bool {
	if s == "" {
		return true
	}
	for _, v := range s {
		if v != '\n' {
			return false
		}
	}
	return true
}
