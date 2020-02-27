package main


import (
	"fmt"
	"math"
)


func main() {
	fmt.Print("Ceil 切り上げ\n")
	fmt.Printf("%.1f\n%.1f\n\n", math.Ceil(1.1), math.Ceil(1.9))

	fmt.Print("Floor 切り下げ\n")
	fmt.Printf("%.1f\n%.1f\n\n", math.Floor(1.1), math.Floor(1.9))

	fmt.Print("Round 四捨五入\n")
	fmt.Printf("%.1f\n%.1f\n\n", math.Round(1.1), math.Round(1.9))

	fmt.Print("小数点 四捨五入\n")
	fmt.Printf("%.2f\n%.2f\n\n", math.Round(1.12345*100)/100, math.Round(1.98765*100)/100)

	var i int = 5
	fmt.Print(float64(i)*1.1)
}
