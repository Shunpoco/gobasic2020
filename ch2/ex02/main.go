package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ch2/ex02/conv"
)

func main() {
	// 入力
	// 引数として渡されている場合はそれを、渡されていない場合は標準入力から取得する
	var inputs []string
	if len(os.Args) > 1 {
		inputs = os.Args[1:]
	} else {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		inputs = strings.Split(stdin.Text(), " ")
	}

	// Convert
	for _, input := range inputs {
		t, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Converter: %v\n", err)
			os.Exit(1)
		}
		tempF := conv.Fahrenheit(t)
		tempC := conv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", tempF, conv.FToC(tempF), tempC, conv.CToF(tempC))
		meter := conv.Meter(t)
		feet := conv.Feet(t)
		fmt.Printf("%s = %s, %s = %s\n", meter.String(), conv.MToFt(meter), feet.String(), conv.FtToM(feet))
		kg := conv.Kilogram(t)
		lb := conv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n", kg.String(), conv.KgToLb(kg), lb.String(), conv.LbToKg(lb))
	}

}
