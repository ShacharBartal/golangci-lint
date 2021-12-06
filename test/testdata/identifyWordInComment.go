//args: -EidentifyWordInComment

package testdata

// nolint - before function declaration  // ERROR "word `nolint` is contained"
func checkSpecialWordInCommentTst() {
	//nolint - in function // ERROR "word `nolint` is contained"
}

func checkSpecialWordInCommentTst2() int {
	return 1 // nolint:check for nolint with a linter report // ERROR "word `nolint` is contained"
}

func checkSpecialWordInCommentTst3() bool {
	return true // nolint:check // with some explain // ERROR "word `nolint` is contained"
}

func checkSpecialWordInCommentTst4() string {
	return "dominos" // dominos should not be in our code // ERROR "word `dominos` is contained"
}

// nolint free in the air // ERROR "word `nolint` is contained"

func checkWordInCommentTst5() { // here please todo things // ERROR "word `todo` is contained"
}
