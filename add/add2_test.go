package add

import "testing"

func TestAdd2(t *testing.T){
	result := add2(10, 10);
	expectedResult := 20;
	
	if result != expectedResult{
		t.Errorf("Got %q, expected result %q", result, expectedResult)
	}
}
	
