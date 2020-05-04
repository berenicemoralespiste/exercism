package spaceAge

import (
	"math"
	"testing"
)

type cases struct {
	planet Planet
	in     float64
	want   float64
}

func TestSpaceAge(t *testing.T) {

	data := []cases{
		{"Mars", 2329871239, 39.25379849361774},
		{"Neptune", 8210123456, 1.58},
		{"Jupiter", 901876382, 2.41},
		{"Saturn", 3000000000, 3.23},
		{"Earth", 1000000000, 31.69},
		{"Venus", 189839836, 9.78},
	}

	for _, value := range data {
		got := spaceAge(value.in, value.planet)

		const precision = 0.01
		if math.Abs(got - value.want) > precision  {
			t.Fatalf("The func spaceAge (%v, %v) = %v is no what we need %v", value.in, value.planet, got, value.want)
		}

		t.Log("Success")
	}
}
