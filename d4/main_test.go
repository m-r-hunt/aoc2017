package d4

import "testing"

func TestValidPassPhrase(t *testing.T) {
	if validPassphrase("aa bb cc dd ee") != true {
		t.Fail()
	}
	if validPassphrase("aa bb cc dd aa") != false {
		t.Fail()
	}
	if validPassphrase("aa bb cc dd aaa") != true {
		t.Fail()
	}
}
