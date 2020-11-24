package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func openFile(name string) (*os.File, error) {
	fmt.Println("Opening file")
	return os.Open(name)
}

func closeFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func getFloats(fileName string) ([]float64, error) {
	var numbers []float64
	f, err := openFile(fileName)

	//error handling and clean up should stay together
	if err != nil {
		return nil, err
	}
	defer closeFile(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return numbers, nil
}

func getSum(numbers ...float64) (sum float64) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func checkArg() {
	if len(os.Args) != 2 {
		log.Fatal("argument not specified")
	}
}

func firstExercise() {
	checkArg()

	numbers, err := getFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	sum := getSum(numbers...)
	fmt.Printf("The sum is %.2f\n", sum)
}

//recover from outside (but this doesn't allow the recursive functions to continue, and hoist the error to the main caller!)
func listFilesPanic() {
	e := recover()
	if e, ok := e.(error); ok {
		fmt.Printf("Skipping dir due to error %s\n", e.Error())
	} else {
		panic(e)
	}
}

func getFileInfos(dir string) []os.FileInfo {

	infos, err := ioutil.ReadDir(dir)
	if err != nil {

		//must be called after panicking
		//using iffe allows recursion to continue as long as it's valid, since the outer caller of a panicked func is another panicked func

		//v1 using error
		// defer func(err error) {
		// 	fmt.Printf("Skipping %s due to error: %v\n", dir, err)
		// 	recover()
		// }(err)

		//using whatever value returned by recover, need to type assert to call underlying methods on concrete type of e
		// defer func() {
		// 	e := recover()
		// 	if e, ok := e.(error); ok {
		// 		fmt.Printf("Skipping dir due to error %s\n", e.Error())
		// 	} else {
		// 		panic(e)
		// 	}
		// }()

		panic(err)
	}

	return infos
}

func listFiles(path string) {
	infos := getFileInfos(path)
	for _, f := range infos {
		if f.IsDir() {
			fmt.Printf("Dir: %s\n", f.Name())
			listFiles(filepath.Join(path, f.Name()))
		} else {
			fmt.Printf("File: %s\n", f.Name())
		}
	}
}

func secondExercise() {
	checkArg()
	defer listFilesPanic()
	listFiles(os.Args[1])
}

func main() {
	secondExercise()
}
