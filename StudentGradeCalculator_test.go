package main

import (
	"testing"
)
// test for average calculator
 func TestAverage(t *testing.T){
	grades :=[]int {90,80,70}
	expected := 80.0


	result:= average(grades)
	if result!= expected{
		t.Errorf("expected average %v, but got %v",expected,result)

	}
 }