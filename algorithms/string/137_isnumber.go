package algorithms

// https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/description/
// input: s = "0"
// output: true
type State int
type CharType int

const (
	StateInitial State = iota
	StateIntSign
	StateInteger
	StatePoint
	StatePointWitoutInt
	StateFraction
	StateExp
	StateExpSign
	StateExpNumber
	StateEnd
)

const (
	CharNumber CharType = iota
	CharExp
	CharPoint
	CharSign
	CharSpace
	CharIllegal
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '+', '-':
		return CharSign
	case '.':
		return CharPoint
	case ' ':
		return CharSpace
	default:
		return CharIllegal
	}
}
func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		StateInitial: map[CharType]State{
			CharSpace:  StateInitial,
			CharNumber: StateInteger,
			CharPoint:  StatePointWitoutInt,
			CharSign:   StateIntSign,
		},
		StateIntSign: map[CharType]State{
			CharNumber: StateInteger,
			CharPoint:  StatePointWitoutInt,
		},
		StateInteger: map[CharType]State{
			CharNumber: StateInteger,
			CharExp:    StateExp,
			CharPoint:  StatePoint,
			CharSpace:  StateEnd,
		},
		StatePoint: map[CharType]State{
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StatePointWitoutInt: map[CharType]State{
			CharNumber: StateFraction,
		},
		StateFraction: map[CharType]State{
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StateExp: map[CharType]State{
			CharNumber: StateExpNumber,
			CharSign:   StateExpSign,
		},
		StateExpSign: map[CharType]State{
			CharNumber: StateExpNumber,
		},
		StateExpNumber: map[CharType]State{
			CharNumber: StateExpNumber,
			CharSpace:  StateEnd,
		},
		StateEnd: map[CharType]State{
			CharSpace: StateEnd,
		},
	}
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}

	return state == StateInteger || state == StatePoint || state == StateFraction || state == StateExpNumber || state == StateEnd
}
