package main

import "fmt"
func main() {
	fmt.Println("===== 2.1　基本型 =====")
	fmt.Println("===== 2.1.1 ゼロ値 =====")
	{
		var x int
		var y float32
		var z string

		fmt.Println("x:", x) // x: 0
		fmt.Println("y:", y)  // y: 0
		fmt.Printf("z:|%s|\n", z)	 // z:|| （空文字列）
	}
}
