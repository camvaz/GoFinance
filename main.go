package main

import (
	"GoFinance/math"
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f, err := excelize.OpenFile("TablaDuracion.xlsx")
	fmt.Print("Iniciando lectura de archivo...\n\n")
	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println(err)
		return
	}

	rows := f.GetRows("Sheet1")
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
				fmt.Printf("Error al leer datos, uno de estos no es un número. Dato con error: %s", colCell)
				return
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
		tempProps = math.DuracionProps{Cup: 0, T: 0, K: 0}
	}

	results := make([]float64, 0)

	for _, value := range duracionSlice {
		results = append(results, math.Duracion(value))
	}

	fmt.Print("Resultados:\n\n")
	for i, value := range results {
		fmt.Printf("[%d]: %v\n", i, value)
	}
}
