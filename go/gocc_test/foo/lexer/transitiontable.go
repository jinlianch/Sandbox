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
		case 97 <= r && r <= 103: // ['a','g']
			return 2
		case r == 104: // ['h','h']
			return 3
		case 105 <= r && r <= 122: // ['i','z']
			return 2
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
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case 97 <= r && r <= 100: // ['a','d']
			return 2
		case r == 101: // ['e','e']
			return 4
		case 102 <= r && r <= 122: // ['f','z']
			return 2
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case 97 <= r && r <= 107: // ['a','k']
			return 2
		case r == 108: // ['l','l']
			return 5
		case 109 <= r && r <= 122: // ['m','z']
			return 2
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		case 97 <= r && r <= 107: // ['a','k']
			return 2
		case r == 108: // ['l','l']
			return 6
		case 109 <= r && r <= 122: // ['m','z']
			return 2
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case 97 <= r && r <= 110: // ['a','n']
			return 2
		case r == 111: // ['o','o']
			return 7
		case 112 <= r && r <= 122: // ['p','z']
			return 2
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		}
		return NoState
	},
}
