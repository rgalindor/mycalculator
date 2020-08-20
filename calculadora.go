package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (Calc) Operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return (operador1 + operador2)
	case "-":
		fmt.Println(operador1 - operador2)
		return (operador1 - operador2)
	case "*":
		fmt.Println(operador1 * operador2)
		return (operador1 * operador2)
	case "/":
		fmt.Println(operador1 / operador2)
		return (operador1 / operador2)
	default:
		fmt.Println(operador, "No esta definido")
		return (0)
	}
}

func parsear(entrada string) int {
	operador, err1 := strconv.Atoi(entrada)
	return (operador)
	if err1 != nil {
		fmt.Println(err1)
		return (0)
	} else {
		return (operador)
	}
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	return operacion
}
