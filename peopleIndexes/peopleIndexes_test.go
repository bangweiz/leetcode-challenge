package peopleIndexes

import "testing"

func TestPeopleIndexes(t *testing.T) {
	a := make([][]string, 3)
	a[0] = []string{"nxaqhyoprhlhvhyojanr","ovqdyfqmlpxapbjwtssm","qmsbphxzmnvflrwyvxlc","udfuxjdxkxwqnqvgjjsp","yawoixzhsdkaaauramvg","zycidpyopumzgdpamnty"}
	a[1] = []string{"nxaqhyoprhlhvhyojanr","ovqdyfqmlpxapbjwtssm","udfuxjdxkxwqnqvgjjsp","yawoixzhsdkaaauramvg","zycidpyopumzgdpamnty"}
	a[2] = []string{"ovqdyfqmlpxapbjwtssm","qmsbphxzmnvflrwyvxlc","udfuxjdxkxwqnqvgjjsp","yawoixzhsdkaaauramvg","zycidpyopumzgdpamnty"}
	res := peopleIndexes(a)

	if len(res) != 1 || res[0] != 0 {
		t.Errorf("Wrong answer")
	} else {
		t.Logf("Success")
	}
}