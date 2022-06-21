package lib

import (
	"testing"
)

func TestSwap(t *testing.T){
	actual1 := 2 
	actual2 := 1

	x, y := Swap(1,2)
	if actual1 != x{
		t.Error("error")
	}

	if actual2 != y{
		t.Error("error")
	}
}