package math

import (
	"math"
)

// DuracionProps - Struct con data para sacar duracion
type DuracionProps struct {
	Cup float64
	T   float64
	K   float64
}

// Duracion - Valuacion de bonos
func Duracion(props DuracionProps) float64 {
	tasaCupon := props.Cup / 100
	tasaRendimiento := props.K / 100
	firstOperand := (1 + tasaRendimiento) / tasaRendimiento
	secondOperandNum := props.T*(tasaCupon-tasaRendimiento) +
		(1 + tasaRendimiento)
	secondOperandDiv := (tasaCupon * math.Pow(1+tasaRendimiento, props.T)) -
		(tasaCupon - tasaRendimiento)

	return firstOperand - (secondOperandNum / secondOperandDiv)
}
