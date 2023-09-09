package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularSalario(t *testing.T) {
	testCases := []struct {
		minutosTrabajados int
		categoria         string
		expectedSalario   float64
	}{
		{1800, "C", 30000},
		{1800, "B", 42000},
		{1800, "A", 63000},
		{1800, "D", 0}, // Categoría inválida
	}

	for _, tc := range testCases {
		t.Run(tc.categoria, func(t *testing.T) {
			salario := calcularSalario(tc.minutosTrabajados, tc.categoria)
			assert.Equal(t, tc.expectedSalario, salario, "Salario incorrecto")
		})
	}
}
