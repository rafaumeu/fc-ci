package main
import "fmt"
	func main() {
		fmt.Println(sum(2, 2))
		fmt.Println(sub(3,1))
		fmt.Println(div(20,2))
		fmt.Println(times(5,2))
	}

	func sum(a int, b int) int {
		return a + b
	}
	func sub(a int, b int) int {
		return a - b
	}
	func div(a int, b int) int {
		return a / b
	}
	func times(a int, b int) int {
		return a * b
	}
	