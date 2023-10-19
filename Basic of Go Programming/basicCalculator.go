package main
import (
    "fmt"
    "strconv"
)
func main() {
    var num1, num2 float64
    var operator string
    fmt.Println("Enter two numbers:")
    fmt.Scanln(&num1, &num2)
    fmt.Println("Enter an operator (+, -, *, /):")
    fmt.Scanln(&operator)
    switch operator {
    case "+":
        fmt.Println(num1 + num2)
    case "-":
        fmt.Println(num1 - num2)
    case "*":
        fmt.Println(num1 * num2)
    case "/":
        if num2 == 0 {
            fmt.Println("Error: division by zero")
        } else {
            fmt.Println(num1 / num2)
        }
    default:
        fmt.Println("Invalid operator")
    }
}