package main

import (
	"GoFinance/math"
	"GoFinance/utils"
	"errors"
	"fmt"
	"strconv"
)

func main() {
	f, err := utils.ReadExcel("TablaDuracion.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}

	rows := f.GetRows("Sheet1")
	duracionSlice, errorSlice := fillValues(rows[:][:])

	if errorSlice != nil {
		fmt.Println(errorSlice)
		return
	}

	results := make([]float64, 0)

	for _, value := range duracionSlice {
		results = append(results, math.Duracion(&value))
	}

	fmt.Print("Resultados:\n\n")
	for i, value := range results {
		fmt.Printf("[%d]: %v\n", i, value)
	}
}

func fillValues(rows [][]string) ([]math.DuracionProps, error) {
	tempProps := math.DuracionProps{Cup: 0, T: 0, K: 0}
	duracionSlice := make([]math.DuracionProps, 0)

	for itrRow, row := range rows {
		if itrRow < 2 {
			continue
		}

		for itrCol, colCell := range row {
			if itrCol > 2 {
				break
			}

			value, err := strconv.ParseFloat(colCell, 64)

			if err != nil {
				message := fmt.Sprintf("Error al leer datos, uno de estos no es un número. Dato con error: %s", colCell)
				return nil, errors.New(message)
			}

			switch opt := itrCol; opt {
			case 0:
				tempProps.Cup = value
			case 1:
				tempProps.T = value
			case 2:
				tempProps.K = value
			default:
				break
			}

			fmt.Print(colCell, "\t")
		}

		fmt.Printf("\nDatos de fila %d: ", itrRow)
		fmt.Print(tempProps, "\n\n")
		duracionSlice = append(duracionSlice, tempProps)
	}
	return duracionSlice, nil
}
