package gigasecond

import (
	"fmt"
	"testing"
	"time"
)

const (
	fmtD  = "2006-01-02"
	fmtDT = "2006-01-02T15:04:05"
)

type test struct {
	in   time.Time
	want time.Time
}

func TestAddGigasecond(t *testing.T) {
	test1 := test{
		in:   parse("1977-06-13"),
		want: parse("2009-02-19T01:46:40"),
	}
	test2 := test{
		in:   parse("2011-04-25"),
		want: parse("2043-01-01T01:46:40"),
	}
	test3 := test{
		in:   parse("1959-07-19"),
		want: parse("1991-03-27T01:46:40"),
	}
	cases := []test{test1, test2, test3}

	for _, tc := range cases {
		in := tc.in
		want := tc.want
		got := AddGigaSecond(in)
		if !got.Equal(want) {
			t.Fatalf(`FAIL: AddGigasecond(%s) = %s want %s`, in, got, want)
		}
		t.Log("PASS:")
	}
	t.Log("Tested", len(cases), "cases.")
}
func parse(s string) time.Time {

	parseTime, error := time.Parse(fmtDT, s)
	if error != nil {
		parseTime, error = time.Parse(fmtD, s)
	}

	if error != nil {
		fmt.Println("Error on Parse time", error)
	}

	return parseTime

}
