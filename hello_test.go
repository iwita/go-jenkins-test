package main

import "testing"

func TestHello(t *testing.T) {
	emptyResult := hello("")  //should return "Hello Evolve-User"

	if emptyResult != "Hello Evolve-User!" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Evolve-User!", emptyResult)
	}


	//test for valid arguement

	result := hello("Achilleas") //should return Hello Achilleas

	if result != "Hello Achilleas" {
		t.Errorf("hello(\"Achilleas\") failed, expected %v, got %v", "Hello Achilleas", result)
	}
}
