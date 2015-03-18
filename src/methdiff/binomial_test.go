package main

import (
	"math"
	"testing"
)

func TestBindiff(t *testing.T) {
	eps := 0.0000001
	
	act := bindiff(4, 2, 4, 2)
	exp := 1.0
	if math.IsNaN(act) {
		t.Errorf("Bindiff returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
	
	act = bindiff(10, 7, 10, 3)
	exp = 0.090169906616211
	if math.IsNaN(act) {
		t.Errorf("Bindiff returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
	
	act = bindiff(20, 3, 10, 7)
	exp = 0.0033253903340978055
	if math.IsNaN(act) {
		t.Errorf("Bindiff returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
	
	act = bindiff(40, 15, 50, 20)
	exp = 0.8072440871041476
	if math.IsNaN(act) {
		t.Errorf("Bindiff returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
	
	act = bindiff(40, 15, 5, 2)
	exp = 0.9137384313794469
	if math.IsNaN(act) {
		t.Errorf("Bindiff returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
	
	act = bindiff(2000, 750, 5, 2)
	exp = 0.8309711086876369
	if math.IsNaN(act) {
		t.Errorf("Bindiff returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
	
	act = bindiff(2000, 750, 500, 200)
	exp = 0.30192128455458545
	if math.IsNaN(act) {
		t.Errorf("Bindiff returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}

	act = bindiff(2000, 750, 2000, 755)
	exp = 0.8703855631191247
	if math.IsNaN(act) {
		t.Errorf("Bindiff returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}

	act = bindiff(2000, 35, 2000, 30)
	exp = 0.4968381164131785
	if math.IsNaN(act) {
		t.Errorf("Bindiff returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad bindiff result: got %f expected %f.", act, exp)
	}
}

func TestBinoPdf(t *testing.T) {
	eps := 0.0000000001

	act := binoPdf(4, 2, 0.5)
	exp := float64(6) / float64(16)
	if math.IsNaN(act) {
		t.Errorf("BinoPdf returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad pdf result: got %f expected %f.", act, exp)
	}
	
	act = binoPdf(5, 2, 0.3)
	exp = 0.3087
	if math.IsNaN(act) {
		t.Errorf("BinoPdf returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad pdf result: got %f expected %f.", act, exp)
	}
	
	act = binoPdf(200, 3, 3.5/200.0)
	exp = 0.2172827705907441
	if math.IsNaN(act) {
		t.Errorf("BinoPdf returned NaN.")
	}
	if math.Abs(act - exp) > eps {
		t.Errorf("Bad pdf result: got %.20f expected %.20f.", act, exp)
	}
}

