package main

import "testing"

func TestSoma(t *testing.T) {
	
	total := soma(5, 5)
	if total != 10 {
		t.Errorf("Soma incorreta, resultado; %d, resultado correto: %d", total, 10)
	}
}