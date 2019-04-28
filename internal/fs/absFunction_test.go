package fs

import "testing"

func TestNewAbsFunction(t *testing.T){
	abs:= NewAbsFunction()
	if abs.Name()!="Abs"{
		t.Fatal()
	}
}