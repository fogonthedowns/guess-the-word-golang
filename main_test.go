package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	out1 := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	if out1 != 5 {
		t.Errorf("got %d, want %d", out1, 5)
	}
}
