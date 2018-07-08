package main

import "testing"

type testCase struct {
	str StringToParse
	ans bool
}

var testCases = []testCase{
	{"", true},
	{"qwe123QWE", true},
	{"{qwe123QWE}", true},
	{"(qwe123QWE)", true},
	{"[qwe123QWE]", true},
	{"<qwe123QWE>", true},
	{"()", true},
	{"{}", true},
	{"<>", true},
	{"[]", true},
	{"<{([])}>", true},
	{"<<{{(([[]]))}}>>", true},
	{"1<2{3{4<5[6[7]8]9>}}>10", true},
	{"[[((()){{}})<123>]](){s}", true},
	{"(", false},
	{"(qwe123QWE", false},
	{"({qwe123QWE}", false},
	{"(qwe{123QWE)", false},
	{"[qwe123QWE]<", false},
	{"<qwe123QWE>}", false},
	{"(()", false},
	{"{{}", false},
	{"<>>", false},
	{"[]]", false},
	{"<{([)}>", false},
	{"<<{{(([[]])}}>>", false},
	{"1<2{3{4<5[6[7]8]9>}>10", false},
	{"[[((()){{}})<123>](){s}", false},
}

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		err := tc.str.Parse()
		if err != nil && tc.ans == true {
			t.Errorf("Function returns error on correct string")
		}
		if err == nil && tc.ans == false {
			t.Errorf("Function returns no error on incorrect string")
		}
	}
}
