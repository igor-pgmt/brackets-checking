package main

import (
	"errors"
	"strconv"
)

// StringToParse is a type to define parsing method
type StringToParse string

// Parse is the main checking function
func (s *StringToParse) Parse() error {
	checkStack := new(filo)   // new stack
	for i, char := range *s { // iterate string rune by rune
		pair, pos := lp.isLex(char) // checking for a pair
		if pair != nil {
			if pos == true { // if the char is an opening char,
				checkStack.push(char) // add to stack
			} else { // if the char is a closing char,
				last, err := checkStack.getLast() // get the last rune from the stack
				if err != nil {                   // if there is no runes anymore
					return errors.New("Лишняя закрывающая скобка: \"" + string(char) + "\" на позиции " + strconv.Itoa(i) + ": " + string(*s)[:i+1] + "←")
				}
				if pair.checkOpen(last) { // if the pair is correct,
					checkStack.pop() // remove the rune from the stack
				} else {
					return errors.New("Несовпадение скобок: найдено \"" + string(char) + " на позиции " + strconv.Itoa(i) + ": " + string(*s)[:i+1] + "←")
				}
			}
		}
	}
	// if there are some unclosed characters
	if len(checkStack.stack) > 0 {
		return errors.New("Отсутствуют закрывающие скобки для: \"" + string(checkStack.stack[len(checkStack.stack)-1]) + "\"")
	}

	return nil
}
