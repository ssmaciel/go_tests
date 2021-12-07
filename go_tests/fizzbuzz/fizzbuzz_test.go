package fizzbuzz

import "testing"

func TestFizzbuzzNumero(t *testing.T) {
	resultado := Fizzbuzz(1)
	esperado := "1"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
func TestFizzbuzzFizz(t *testing.T) {
	resultado := Fizzbuzz(3)
	esperado := "fizz"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
func TestFizzbuzzBuzz(t *testing.T) {
	resultado := Fizzbuzz(5)
	esperado := "buzz"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
func TestFizzbuzz(t *testing.T) {
	resultado := Fizzbuzz(15)
	esperado := "fizzbuzz"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
