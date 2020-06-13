package slides

import (
	"testing"
)

func TestStubbedImplementation(t *testing.T) {

	_, err := Execute("./test")		
	if err != nil {
		t.Log("Error testing stubbed implementation")
		t.Fail()
	}

}
