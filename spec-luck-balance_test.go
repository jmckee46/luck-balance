package main

import "testing"

func TestLuckBalance(t *testing.T) {
	k := int32(3)
	contests := [][]int32{
		[]int32{5, 1},
		[]int32{2, 1},
		[]int32{1, 1},
		[]int32{8, 1},
		[]int32{10, 0},
		[]int32{5, 0},
	}

	luck := luckBalance(k, contests)

	if luck != 29 {
		t.Errorf("got %d instead of 29", luck)
	}

}
