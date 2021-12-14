package wires

func validColors() []string {
	return []string{"R", "B", "W", "S", "G"}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
