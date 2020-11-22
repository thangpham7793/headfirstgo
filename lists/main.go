package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/headfirstgo/general"
	"github.com/headfirstgo/lists/utils"
)

func prettyPrintFloat(num float64) {
	fmt.Printf("%.2f\n", num)
}

func grow(original *[]float64) {
	extra := make([]float64, len(*original))
	*original = append(*original, extra...)
}

func main() {

	if len(os.Args) == 1 {
		log.Fatal("no file path found")
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		if len(os.Args) >= 2 {
			var nums []float64
			for _, n := range os.Args[1:] {
				num, err := strconv.ParseFloat(n, 64)
				general.HandleErr((err))
				nums = append(nums, num)
			}
			prettyPrintFloat(utils.Average(nums...))
			return
		}
		general.HandleErr(err)
	}

	defer func() {
		err := f.Close()
		general.HandleErr(err)
	}()

	//nil slice is treated as an empty slice
	var sum []float64
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		general.HandleErr((err))
		sum = append(sum, num)
	}
	general.HandleErr(scanner.Err())

	prettyPrintFloat(utils.Average(sum...))
}
