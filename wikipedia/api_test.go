package wikipedia

import (
	"fmt"
	"testing"
)

func TestWikipedia_api(t *testing.T) {

	if s := WikipediaAPI(`test`); s == nil {
		t.Fail()
	}

	if s := Search("325f547b68987098234c46:::::@#$%!@#@^%&&^*", "en"); s[0] != errStr {
		t.Fail()
	}

	if s := Search("test", "en"); s[0] == errStr {
		t.Fail()
	} else {
		fmt.Println(s)
	}
}
