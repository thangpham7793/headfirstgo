package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/thangpham7793/headfirstgo/calculator/number"
)

func splitArgs(args []string) (operands []float64, operations []string, err error) {

	operands = make([]float64, 0)
	operations = make([]string, 0)

	for i, a := range args {
		fmt.Println(a)
		switch {
		case i%2 == 0:
			num, err := strconv.ParseFloat(a, 64)
			if err != nil {
				log.Fatalf("%v\n\nPerhaps you forget to escape operand *", err)
			}
			operands = append(operands, num)
		default:
			operations = append(operations, a)
		}
	}
	return
}

func calculate(operands []float64, operations []string) (number.Number, error) {

	res := number.Number(operands[0])
	for i := 1; i < len(operands); i++ {
		switch operations[i-1] {
		case "+":
			res = res.Add(operands[i])
		case "-":
			res = res.Substract(operands[i])
		case "*", "x":
			res = res.Multiply(operands[i])
		case "/":
			res = res.Divide(operands[i])
		default:
			return 0, fmt.Errorf(`invalid operation "%s"`, operations[i/2])
		}
	}
	return res, nil
}

func printRes(res number.Number, args []string) {
	fmt.Printf("%s is %.2f \n", strings.Join(args, " "), res)
}

func main() {

	//using * will make bash interpret it as all args though, need to escape it or not use it at all

	args := os.Args[1:]
	operands, operations, err := splitArgs(args)

	if err != nil {
		log.Fatal(err)
	}

	res, err := calculate(operands, operations)

	if err != nil {
		log.Fatal(err)
	}

	printRes(res, args)
	// ten := number.Number(10)
	// res = ten.Add(3).Add(5).Substract(23)

	// fmt.Println(res.Absolute(), ten)
}
