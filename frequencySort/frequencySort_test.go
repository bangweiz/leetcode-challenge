package frequencySort

import "testing"

func TestFrequencySort(t *testing.T) {
	res := frequencySort("tree")
	if res == "eetr" || res == "eert" {
		t.Log("Success")
	} else {
		t.Errorf("Failed, the res should be %v or %v, but it is %v", "eetr", "eert", res)
	}
}