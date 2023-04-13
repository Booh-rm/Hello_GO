package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64
    var operator string

    fmt.Print("Introduzca el primer número: ")
    fmt.Scanln(&num1)

    fmt.Print("Introduzca el segundo número: ")
    fmt.Scanln(&num2)

    fmt.Print("Introduzca el operador (+, -, *, /): ")
    fmt.Scanln(&operator)

    switch operator {
    case "+":
        fmt.Printf("Resultado: %.2f\n", num1+num2)
    case "-":
        fmt.Printf("Resultado: %.2f\n", num1-num2)
    case "*":
        fmt.Printf("Resultado: %.2f\n", num1*num2)
    case "/":
        if num2 == 0 {
            fmt.Println("Error: No se puede dividir por cero")
        } else {
            fmt.Printf("Resultado: %.2f\n", num1/num2)
        }
    default:
        fmt.Println("Operador inválido")
    }
}
