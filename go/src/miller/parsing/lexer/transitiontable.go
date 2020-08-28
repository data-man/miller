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
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 36: // ['$','$']
			return 2
		case r == 38: // ['&','&']
			return 3
		case r == 45: // ['-','-']
			return 4
		case r == 48: // ['0','0']
			return 5
		case 49 <= r && r <= 57: // ['1','9']
			return 6
		case r == 58: // [':',':']
			return 7
		case r == 59: // [';',';']
			return 8
		case r == 61: // ['=','=']
			return 9
		case r == 63: // ['?','?']
			return 10
		case r == 94: // ['^','^']
			return 11
		case r == 102: // ['f','f']
			return 12
		case r == 116: // ['t','t']
			return 13
		case r == 124: // ['|','|']
			return 14
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
		case 48 <= r && r <= 57: // ['0','9']
			return 15
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 38: // ['&','&']
			return 18
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 19
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 6
		case r == 120: // ['x','x']
			return 20
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 6
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
		case r == 94: // ['^','^']
			return 21
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 22
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 23
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 124: // ['|','|']
			return 24
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 15
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 15
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 15
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 19
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 25
		case 65 <= r && r <= 70: // ['A','F']
			return 25
		case 97 <= r && r <= 102: // ['a','f']
			return 25
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 26
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case r == 117: // ['u','u']
			return 27
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
		case 48 <= r && r <= 57: // ['0','9']
			return 25
		case 65 <= r && r <= 70: // ['A','F']
			return 25
		case 97 <= r && r <= 102: // ['a','f']
			return 25
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 115: // ['s','s']
			return 28
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 29
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 30
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		}
		return NoState
	},
}
