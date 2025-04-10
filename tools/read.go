package tools

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(route string) []string {
	data, err := os.ReadFile(route)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
	dataFormat := string(data)
	return strings.Split(dataFormat, "\r\n")
}
