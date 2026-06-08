func scoreOfString(s string) int {
	total := 0

	for i := 0; i < len(s) - 1; i++ {
		dif := int(s[i]) - int(s[i+1])

		if dif < 0 {
			dif = -dif
		}
		total += dif
	}

	return total
}
