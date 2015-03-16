package main

import (
	"math"
	"testing"
)

func TestBindiff(t *testing.T) {
	eps := 0.000000000000001
	
	act := bindiff(4, 2, 4, 2)
	exp := 1.0
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
	
	act = bindiff(10, 7, 10, 3)
	exp = 0.090169906616211
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
	
	act = bindiff(20, 3, 10, 7)
	exp = 0.0033253903340978055
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
	
	act = bindiff(200, 3, 200, 4)
	exp = 0.7199739329900385
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
}

func TestBinoPdf(t *testing.T) {
	eps := 0.000000000000001

	act := binoPdf(4, 2, 0.5)
	exp := float64(6) / float64(16)
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad pdf result: got %f expected %f.", act, exp)
	}
	
	act = binoPdf(5, 2, 0.3)
	exp = 0.3087
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad pdf result: got %f expected %f.", act, exp)
	}
}

