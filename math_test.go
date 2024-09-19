package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(10, 10)
	esperado := 20

	if resultado != esperado {
		t.Errorf("Resultado incorreto: obtido %v, esperado %v", resultado, esperado)
	}
}
