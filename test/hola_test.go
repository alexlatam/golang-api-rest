package test

import (
	"testing"
)

// Toda funcoin de test debe empezar en Mayuscula
func TestHola(t *testing.T) {
	src := "Hola"
	if src != "Hola" {
		// Si el test falla, se ejecuta el metodo Error
		t.Error("La variable src no retorna Hola", src)
	}
}
