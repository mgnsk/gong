// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 34: // ['"','"']
			return 2
		case r == 35: // ['#','#']
			return 3
		case r == 36: // ['$','$']
			return 4
		case r == 41: // [')',')']
			return 5
		case r == 42: // ['*','*']
			return 6
		case r == 45: // ['-','-']
			return 7
		case r == 46: // ['.','.']
			return 8
		case r == 47: // ['/','/']
			return 9
		case r == 48: // ['0','0']
			return 10
		case 49 <= r && r <= 57: // ['1','9']
			return 11
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 91: // ['[','[']
			return 13
		case r == 93: // [']',']']
			return 14
		case r == 94: // ['^','^']
			return 15
		case r == 97: // ['a','a']
			return 16
		case r == 98: // ['b','b']
			return 17
		case r == 99: // ['c','c']
			return 18
		case r == 100: // ['d','d']
			return 12
		case r == 101: // ['e','e']
			return 19
		case 102 <= r && r <= 111: // ['f','o']
			return 12
		case r == 112: // ['p','p']
			return 20
		case 113 <= r && r <= 114: // ['q','r']
			return 12
		case r == 115: // ['s','s']
			return 21
		case r == 116: // ['t','t']
			return 22
		case r == 117: // ['u','u']
			return 12
		case r == 118: // ['v','v']
			return 23
		case 119 <= r && r <= 122: // ['w','z']
			return 12
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 24
		default:
			return 2
		}
	},
	// S3
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 47: // ['/','/']
			return 25
		case r == 48: // ['0','0']
			return 26
		case 49 <= r && r <= 57: // ['1','9']
			return 27
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 11
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 115: // ['s','s']
			return 28
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 29
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 104: // ['h','h']
			return 30
		case r == 111: // ['o','o']
			return 31
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 32
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 33
		case r == 114: // ['r','r']
			return 34
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 35
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 36
		case r == 105: // ['i','i']
			return 37
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 38
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 39
		default:
			return 25
		}
	},
	// S26
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 27
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case r == 115: // ['s','s']
			return 40
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 41
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 42
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 43
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case r == 100: // ['d','d']
			return 44
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 45
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 46
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 47
		case r == 111: // ['o','o']
			return 48
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case r == 109: // ['m','m']
			return 49
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case r == 109: // ['m','m']
			return 50
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 51
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case r == 105: // ['i','i']
			return 52
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 53
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 54
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		case r == 121: // ['y','y']
			return 55
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		case r == 103: // ['g','g']
			return 56
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 57
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case r == 112: // ['p','p']
			return 58
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		case r == 112: // ['p','p']
			return 59
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 60
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 61
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		case r == 103: // ['g','g']
			return 62
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 63
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 64
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 65
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 66
		}
		return NoState
	},
	// S58
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S59
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 67
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		case r == 115: // ['s','s']
			return 68
		}
		return NoState
	},
	// S61
	func(r rune) int {
		switch {
		case r == 99: // ['c','c']
			return 69
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 70
		}
		return NoState
	},
	// S63
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 71
		}
		return NoState
	},
	// S64
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 72
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 73
		}
		return NoState
	},
	// S66
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S67
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case r == 105: // ['i','i']
			return 74
		}
		return NoState
	},
	// S69
	func(r rune) int {
		switch {
		case r == 105: // ['i','i']
			return 75
		}
		return NoState
	},
	// S70
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 76
		}
		return NoState
	},
	// S72
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 77
		}
		return NoState
	},
	// S73
	func(r rune) int {
		switch {
		case r == 109: // ['m','m']
			return 78
		}
		return NoState
	},
	// S74
	func(r rune) int {
		switch {
		case r == 103: // ['g','g']
			return 79
		}
		return NoState
	},
	// S75
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 80
		}
		return NoState
	},
	// S76
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S77
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S78
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S79
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S80
	func(r rune) int {
		switch {
		case r == 121: // ['y','y']
			return 81
		}
		return NoState
	},
	// S81
	func(r rune) int {
		switch {
		}
		return NoState
	},
}
