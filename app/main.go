package main

import (
	"conversions"
	"fmt"
)

func main() {
	fmt.Printf("\n8 en binaire = %v\n", conversions.NewConversion(8, 2))
	fmt.Printf("\n7 en binaire = %v\n", conversions.NewConversion(7, 2))
	fmt.Printf("\n5 en binaire = %v\n", conversions.NewConversion(5, 2))
	fmt.Printf("\n6 en binaire = %v\n", conversions.NewConversion(6, 2))

	fmt.Printf("\n10 en octal = %v\n", conversions.NewConversion(10, 8))
	fmt.Printf("\n13 en octal = %v\n", conversions.NewConversion(13, 8))

	fmt.Printf("\n10 en hexa = %v\n", conversions.NewConversion(10, 16))
	fmt.Printf("\n13 en hexa = %v\n", conversions.NewConversion(13, 16))
	fmt.Printf("\n254 en hexa = %v\n", conversions.NewConversion(254, 16))

	conv := conversions.NewConversion(254, 16)
	fmt.Printf("\nOxFE en base 10 = %v\n", conv.ToDec())

}
