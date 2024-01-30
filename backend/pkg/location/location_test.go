package location

import (
	"testing"
)

func TestLoc(t *testing.T) {
	cases := []struct {
		longitude float64
		latitude  float64
		expected  string
	}{
		{84.10818142193905, 28.17080768929822, "Kaski"},
		{82.12732911296517, 28.952104929237983, "Jajarkot"},
		{86.8343048966655, 27.61985296103812, "Solukhumbu"},
	}

	// Iterate over test cases
	for _, tc := range cases {
		result := GetDistrict(tc.longitude, tc.latitude)
		if result != tc.expected {
			t.Errorf("GetDistrict(%f, %f) = %v, expected %v", tc.longitude, tc.latitude, result, tc.expected)
		}
	}
}
