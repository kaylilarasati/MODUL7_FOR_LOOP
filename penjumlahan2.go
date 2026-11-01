package main
import "fmt"

func main() {
	var x, y, n int
	var s float64
	fmt.Scan(&x, &y)
	n = (y - x) + 1
	s = float64(n * (x + y)) / 2
	fmt.Println(s)
}