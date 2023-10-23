package string

import "strconv"

func countSeniors(details []string) int {
	count := 0
	for _, detial := range details {
		age, err := strconv.Atoi(detial[11:13])
		if err != nil {
			continue
		}
		if age > 60 {
			count++
		}
	}

	return count
}
