package main

import (
	"fmt"
	"github.com/hardselius/gtpl/ch02/ex02/lengthconv"
	"github.com/hardselius/gtpl/ch02/ex02/weightconv"
	"os"
	"strconv"

	"github.com/hardselius/gtpl/ch02/ex02/tempconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert: %v\n",  err)
				os.Exit(1)
			}
			convert(t)
		}
		os.Exit(0)
	}

	for {
		var t float64
		if _, err := fmt.Scanf("%f", &t); err != nil {
			fmt.Fprintf(os.Stderr, "convert: %v\n", err)
		}
		convert(t)
	}
}

func convert(t float64) {
	temp(t)
	length(t)
	weight(t)
}

func temp(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func length(t float64) {
	f := lengthconv.Foot(t)
	m := lengthconv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, lengthconv.FtToM(f), m, lengthconv.MToFt(m))
}

func weight(t float64) {
	lb := weightconv.Pound(t)
	kg := weightconv.Kilogram(t)
	fmt.Printf("%s = %s, %s = %s\n",
		lb, weightconv.LbToKg(lb), kg, weightconv.KgToLb(kg))
}
