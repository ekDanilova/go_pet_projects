package main

import "fmt"

func main() {
    const (
		Addition = "+"
		Subtraction = "-"
		Multiplication = "*"
		Division = "/"
		Modulus = "%"
    )

    var input1, input2 float64
    var operation string
    fmt.Scan(&input1, &input2, &operation)
    fmt.Println(input1, input2, operation)
    switch operation {
        case Addition:
        fmt.Println(input1 + input2)
        case Subtraction:
        fmt.Println(input1 - input2)
        case Multiplication:
        fmt.Println(input1 * input2)
        case Division:
            if input2 == 0 {
                fmt.Println("Делить на ноль нельзя!")
            } else {
                fmt.Println(input1 / input2)
            }
        case Modulus:
            if input2 == 0 {
                fmt.Println("Делить на ноль нельзя!")
            } else {
        		fmt.Println(int(input1)%int(input2))
            }
    }
}
