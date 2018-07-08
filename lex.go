package main

// lexPair is a type for opening ang closing characters
type lexPair struct {
	open, close rune
}

// checkOpen checks if the input rune is an opening bracket or not
func (lp *lexPair) checkOpen(s rune) bool {
	if lp.open == s {
		return true
	}
	return false
}

// lexPairs is a type for the array of the lexical pairs
type lexPairs []lexPair

// Definition of all lexical pairs
var lp = lexPairs{
	{'{', '}'},
	{'[', ']'},
	{'(', ')'},
	{'<', '>'},
}

// isLex checks if the rune has a lexical pair
// and returns both the pair and information about the bracket:
// "true" for the opening symbol and "false" for the closing symbol.
func (lp *lexPairs) isLex(s rune) (*lexPair, bool) {
	for _, pair := range *lp {
		if pair.open == s {
			return &pair, true
		}
		if pair.close == s {
			return &pair, false
		}
	}
	return nil, false
}
