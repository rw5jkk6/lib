package lib

import (
	"testing"
)

func TestSwap(t *testing.T){
	tests := []struct{
		arg1 int
		arg2 int
		exp1 int
		exp2 int
	}{
		{1,2,2,1},
		{0,1,1,0},
		{100,-1,-1,100},
	}
	
	for _, tt := range tests{
		x, y := Swap(tt.arg1, tt.arg2)
		if tt.exp1 != x{
			t.Error("Error")
		}
		if tt.exp2 != y{
			t.Error("Error")
		}
	}

	// actual1 := 2 
	// actual2 := 1

	// x, y := Swap(1,2)
	// if actual1 != x{
	// 	t.Error("error")
	// }

	// if actual2 != y{
	// 	t.Error("error")
	// }
}

func TestSwapPointa(t *testing.T){
	actual1 := 1
	actual2 := 2

	SwapPointa(&actual1, &actual2)

	expect1 := 2
	expect2 := 1

	if expect1 != actual1{
		t.Error("error")
	}
	if expect2 != actual2{
		t.Error("error")
	}
}