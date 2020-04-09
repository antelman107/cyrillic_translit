package cyrillic_translit

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	for _, v := range []struct {
		input  string
		output string
	}{
		{"Усть Ижора", "Ust Izhora"},
	} {
		res := Do("Усть Ижора")
		if res != v.output {
			panic(
				fmt.Sprintf(
					"invalid transliteration: %s -> %s (%s)",
					v.input,
					res,
					v.output,
				))
		}
	}

}
