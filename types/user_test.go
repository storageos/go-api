package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCommaSliceUnmarshal(t *testing.T) {
	expected := commaStr{"a", "b", "c"}
	for _, data := range []string{
		`"a,b,c"`,
		`["a","b","c"]`,
	} {
		s := commaStr{}
		if err := json.Unmarshal([]byte(data), &s); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(s, expected) {
			t.Fatalf("expect %v got %v", expected, s)
		}
	}
}
