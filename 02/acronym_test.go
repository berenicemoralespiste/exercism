package acronym

import "testing"

type test struct {
	in string
	want string

}

func TestAcronym(t *testing.T)  {

	cases := [] test{
		{in: "Chief executive officer", want: "CEO"},
		{in: "Do it yourself)", want: "DIY"},
		{in: "Frequently asked questions", want: "FAQ"},
		{in: "For your information", want: "FYI"},
	}

	for _ , value := range cases{
		got := acronym(value.in)

		if got != value.want{
			t.Fatalf("The result of acronym(%s) = %s and we want %s", value.in, got, value.want)
		}
		t.Log("PASS:")
	}
	t.Log("Tested", len(cases), "cases.")

}