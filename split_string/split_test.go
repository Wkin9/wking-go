package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {

	got := Split("babcbef", "b")
	want := []string{"", "a", "c", "ef"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v,but got: %v", got, want)

	}
}

func TestSplit_2(t *testing.T) {

	got := Split("abvcddbbx", "b")
	want := []string{"a", "vcdd", "", "x"}

	if !reflect.DeepEqual(got, want) {

		t.Errorf("want: %v,but got: %v", want, got)
	}

}
