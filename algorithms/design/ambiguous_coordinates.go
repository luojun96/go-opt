package design

// https://leetcode.cn/problems/ambiguous-coordinates/description/
// input: "(123)"
// output: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]
func ambiguousCoordinates(s string) []string {
	n := len(s) - 2
	s = s[1 : len(s)-1]

	var res []string
	for l := 1; l < n; l++ {
		lt := getPos(s[:l])
		if len(lt) == 0 {
			continue
		}
		rt := getPos(s[l:])
		if len(rt) == 0 {
			continue
		}
		for _, i := range lt {
			for _, j := range rt {
				res = append(res, "("+i+", "+j+")")
			}
		}
	}
	return res

}

func getPos(s string) (pos []string) {
	if s[0] != '0' || s == "0" {
		pos = append(pos, s)
	}
	for p := 1; p < len(s); p++ {
		if p != 1 && s[0] == '0' || s[len(s)-1] == '0' {
			continue
		}
		pos = append(pos, s[:p]+"."+s[p:])
	}
	return pos
}
