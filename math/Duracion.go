package math

import (
	"math"
)

// Duracion - Valuacion de bonos
func Duracion(cup float64, t float64, k float64) float64 {
	tasaCupon := cup / 100
	tasaRendimiento := k / 100
	firstOperand := (1+tasaRendimiento)/ tasaRendimiento
	secondOperandNum := t * (tasaCupon - tasaRendimiento) + 
							(1+tasaRendimiento)
	secondOperandDiv := (tasaCupon * math.Pow(1+tasaRendimiento,t)) - 
						(tasaCupon - tasaRendimiento)

	return firstOperand - (secondOperandNum / secondOperandDiv);
}