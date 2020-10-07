package utils

import (
	"fmt"
	"time"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// ReadExcel - Lectura de archivos excel
func ReadExcel(path string) (*excelize.File, error ) {
	f, err := excelize.OpenFile("TablaDuracion.xlsx")
	fmt.Print("Iniciando lectura de archivo...\n\n")
	time.Sleep(2 * time.Second)

	if err != nil {
		return nil, errors.New("Error al abrir excel")
	}

	return f, nil
}